package vald

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"

	"sync"
	"syscall"
	"time"

	broadcast "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"

	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	_ "google.golang.org/protobuf/proto"

	"github.com/axelarnetwork/utils/jobs"

	"github.com/cosmos/cosmos-sdk/client"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/config"

	"github.com/fairblock/dkg-core/app"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/utils"

	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss"

	"github.com/fairblock/dkg-core/x/tss/tofnd"
	tssTypes "github.com/fairblock/dkg-core/x/tss/types"
)

// RW grants -rw------- file permissions
const RW = 0600

// RWX grants -rwx------ file permissions
const RWX = 0700

var once sync.Once
var cleanupCommands []func()

// GetValdCommand returns the command to start vald
func GetValdCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "vald-start",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			if !cmd.Flags().Changed(flags.FlagFrom) {
				if err := cmd.Flags().Set(flags.FlagFrom, serverCtx.Viper.GetString("broadcast.broadcaster-account")); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			logger := serverCtx.Logger.With("module", "vald")

			
			defer once.Do(cleanUp)

			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				sig := <-sigs
				logger.Info(fmt.Sprintf("captured signal \"%s\"", sig))
				once.Do(cleanUp)
			}()

			cliCtx, err := sdkClient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// dynamically adjust gas limit by simulating the tx first
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithSimulateAndExecute(true)

			valdConf := config.DefaultValdConfig()
			if err := serverCtx.Viper.Unmarshal(&valdConf); err != nil {
				panic(err)
			}

			valAddr := serverCtx.Viper.GetString("validator-addr")
			valKey := serverCtx.Viper.GetString("validator-key")
			capacity := serverCtx.Viper.GetString("channel-capacity")

			// valList := serverCtx.Viper.GetString("validator-key")
			if capacity == "" {
				return fmt.Errorf("capacity not set")
			}

			if valAddr == "" {
				return fmt.Errorf("validator address not set")
			}
			if valKey == "" {
				return fmt.Errorf("validator key not set")
			}
			valdHome := filepath.Join(cliCtx.HomeDir, "vald")
			if _, err := os.Stat(valdHome); os.IsNotExist(err) {
				logger.Info(fmt.Sprintf("folder %s does not exist, creating...", valdHome))
				err := os.Mkdir(valdHome, RWX)
				if err != nil {
					return err
				}
			}

			var recoveryJSON []byte
			recoveryFile := serverCtx.Viper.GetString("tofnd-recovery")
			if recoveryFile != "" {
				recoveryJSON, err = ioutil.ReadFile(recoveryFile)
				if err != nil {
					return err
				}
				if len(recoveryJSON) == 0 {
					return fmt.Errorf("JSON file is empty")
				}
			}

			fPath := filepath.Join(valdHome, "state.json")
			stateSource := NewRWFile(fPath)

			logger.Info("start listening to events")
			listen(cliCtx, txf, valdConf, valAddr, valKey, recoveryJSON, stateSource, logger, capacity)
			logger.Info("shutting down")
			return nil
		},
	}
	setPersistentFlags(cmd)
	flags.AddTxFlagsToCmd(cmd)
	values := map[string]string{
		flags.FlagKeyringBackend: "test",
		flags.FlagGasAdjustment:  "2",
		flags.FlagBroadcastMode:  flags.BroadcastSync,
	}
	utils.OverwriteFlagDefaults(cmd, values, true)

	return cmd
}

func cleanUp() {
	for _, cmd := range cleanupCommands {
		cmd()
	}
}

func setPersistentFlags(cmd *cobra.Command) {
	defaultConf := tssTypes.DefaultConfig()
	cmd.PersistentFlags().String("tofnd-host", defaultConf.Host, "host name for tss daemon")
	cmd.PersistentFlags().String("tofnd-port", defaultConf.Port, "port for tss daemon")
	cmd.PersistentFlags().String("tofnd-recovery", "", "json file with recovery request")
	cmd.PersistentFlags().String("validator-addr", "", "the address of the validator operator")
	cmd.PersistentFlags().String("validator-key", "", "the key of the validator operator")
	cmd.PersistentFlags().String("channel-capacity", "", "the capacity of channel")
	cmd.PersistentFlags().String(flags.FlagChainID, app.Name, "The network chain ID")
}

func listen(ctx sdkClient.Context, txf tx.Factory, dkgCfg config.ValdConfig, valAddr string, valKey string, recoveryJSON []byte, stateSource ReadWriter, logger log.Logger, capacity string) {
	encCfg := app.MakeEncodingConfig()
	cdc := encCfg.Amino

	client, err := tmclient.New(
		fmt.Sprintf(
			"%s:%s",
			"http://0.0.0.0",
			"26657",
		),
		"/websocket",
	)
	if err != nil {
		logger.Error(err.Error())
	}
	err = client.Start()
	if err != nil {
		logger.Error(err.Error())
	}

	bc, err := broadcast.NewCosmosClient(
		fmt.Sprintf(
			"%s:%s",
			"127.0.0.1",
			"9090",
		),
		valKey,
		"dkg",
	)
	if err != nil {
		logger.Error(err.Error())
	}
	tssMgr := createTSSMgr(client, bc, ctx, dkgCfg, logger, valAddr, cdc)
	if recoveryJSON != nil && len(recoveryJSON) > 0 {
		if err = tssMgr.Recover(recoveryJSON); err != nil {
			panic(fmt.Errorf("unable to perform tss recovery: %v", err))
		}
	}
	msg := dkgnet.MsgRegisterValidator{Creator:valAddr,Address:valAddr,Participation:true}
	

_, err = bc.BroadcastTx(&msg, false)
if err != nil {
	panic(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing register msg"))
}
	channelCap, err := strconv.ParseUint(capacity, 10, 64)
	if err != nil {
		panic(err)
	}
	query := "tm.event = 'Tx'"
	subscriber, err := client.Subscribe(context.Background(), "", query, int(channelCap))
	if err != nil {
		panic(err)
	}
	out, err := client.Subscribe(context.Background(), "", "tm.event='NewBlock'", int(channelCap))
	if err != nil {
		panic(err)
	}

	js := []jobs.Job{

		Consume(subscriber, tssMgr),
		ConsumeH(out, tssMgr),
	}

	logErr := func(err error) { logger.Error(err.Error()) }

	mgr := jobs.NewMgr(logErr)

	mgr.AddJobs(js...)

	mgr.Wait()

}

// Consume processes all events from the given subscriber with the given function.
// Do not consume the same subscriber multiple times.
func Consume(subscriber <-chan ctypes.ResultEvent, tssMgr *tss.Mgr) jobs.Job {

	return func(errChan chan<- error) {
		for {
			select {
			case e := <-subscriber:
				go func() {
					defer recovery(errChan)
					d := e.Data.(tmtypes.EventDataTx).Result.Log

					e2 := e.Data.(tmtypes.EventDataTx).Result.Events

					var events []tss.EventMsg
					if err := json.Unmarshal([]byte(d), &events); err != nil {
						errChan <- err
					}
					
					if len(events) > 0 {
						if len(events[0].Events) > 0 {

							if events[0].Events[0].Type == "keygen" {

								key := events[0].Events[0].Attributes[0].Key

								// if key == "start" {

								// 	if err := tssMgr.ProcessKeygenStart(events); err != nil {
								// 		errChan <- err
								// 	}

								// }
								if key == "message" {

									if err := tssMgr.ProcessKeygenMsg(e2, e.Data.(tmtypes.EventDataTx).Height); err != nil {
										errChan <- err
									}
								}
								if key == "dispute" {
									fmt.Println("received: ", events[0].Events)
									if err := tssMgr.ProcessKeygenMsgDispute(events[0].Events); err != nil {
										errChan <- err
									}
								}
							}
						}
					}

				}()

			}
		}
	}
}
func ConsumeH(subscriber <-chan ctypes.ResultEvent, tssMgr *tss.Mgr) jobs.Job {

	return func(errChan chan<- error) {
		for {
			select {
			case e := <-subscriber:
				go func() {

					newBlock := e.Data.(tmtypes.EventDataNewBlock).ResultEndBlock.Events
					//fmt.Println(newBlock)
					if len(newBlock) > 0 {
						if newBlock[0].Type == "dkg-timeout" {
							fmt.Println("timeout-----------------------------------------------")
							tssMgr.CheckTimeout(newBlock[0])
						}
						if newBlock[0].Type == "dkg-mpk" {
							fmt.Println("mpk from chain: ", newBlock[0].Attributes[0].Value)
							tss.SetMpk(newBlock[0].Attributes[0].Value, string(newBlock[0].Attributes[1].Value))

						}
						if newBlock[0].Type == "keygen" {
							if err := tssMgr.ProcessKeygenStart(newBlock[0]); err != nil {
								errChan <- err
							}
						}
					}

				}()

			}
		}
	}
}
func recovery(errChan chan<- error) {
	if r := recover(); r != nil {
		errChan <- fmt.Errorf("job panicked:%s", r)
	}
}

func createTSSMgr(client *tmclient.HTTP, broadcaster *broadcast.CosmosClient, cliCtx client.Context, dkgCfg config.ValdConfig, logger log.Logger, valAddr string, cdc *codec.LegacyAmino) *tss.Mgr {
	create := func() (*tss.Mgr, error) {

		conn, err := tss.Connect(dkgCfg.TssConfig.Host, dkgCfg.TssConfig.Port, dkgCfg.TssConfig.DialTimeout, logger)
		if err != nil {
			return nil, err
		}
		logger.Debug("successful connection to tofnd gRPC server")

		gg20client := tofnd.NewGG20Client(conn)

		tssMgr := tss.NewMgr(client, gg20client, cliCtx, 2*time.Hour, valAddr, broadcaster, logger, cdc)

		return tssMgr, nil
	}
	mgr, err := create()
	if err != nil {
		panic(sdkerrors.Wrap(err, "failed to create tss manager"))
	}

	return mgr
}

// RWFile implements the ReadWriter interface for an underlying file
type RWFile struct {
	path string
}

// NewRWFile returns a new RWFile instance for the given file path
func NewRWFile(path string) RWFile {
	return RWFile{path: path}
}

// ReadAll returns the full content of the file
func (f RWFile) ReadAll() ([]byte, error) { return os.ReadFile(f.path) }

// WriteAll writes the given bytes to a file. Creates a new fille if it does not exist, overwrites the previous content otherwise.
func (f RWFile) WriteAll(bz []byte) error { return os.WriteFile(f.path, bz, RW) }
