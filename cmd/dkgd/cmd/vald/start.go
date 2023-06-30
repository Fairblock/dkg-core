package vald

import (
	//"context"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"

	//	"github.com/cosmos/cosmos-sdk/types/tx"
	//"strings"
	"sync"
	"syscall"
	"time"

	broadcast "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"

	//"github.com/golang/protobuf/proto"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	//	"go.starlark.net/lib/proto"
	_ "google.golang.org/protobuf/proto"
	
	"github.com/axelarnetwork/utils/jobs"

	
	"github.com/cosmos/cosmos-sdk/client"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"

	//rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/config"

	"github.com/fairblock/dkg-core/app"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/utils"

	//"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"
	//"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/btc"
	//btcRPC "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/btc/rpc"
	//"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/evm"
	//evmRPC "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/evm/rpc"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss"

	//utils2 "github.com/fairblock/dkg-core/utils"

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

			// in case of panic we still want to try and cleanup resources,
			// but we have to make sure it's not called more than once if the program is stopped by an interrupt signal
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

			// valList := serverCtx.Viper.GetString("validator-key")
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
			listen(cliCtx, txf, valdConf, valAddr, valKey, recoveryJSON, stateSource, logger)
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
	cmd.PersistentFlags().String(flags.FlagChainID, app.Name, "The network chain ID")
}

func listen(ctx sdkClient.Context, txf tx.Factory, dkgCfg config.ValdConfig, valAddr string, valKey string, recoveryJSON []byte, stateSource ReadWriter, logger log.Logger) {
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
	tssMgr := createTSSMgr(bc, ctx, dkgCfg, logger, valAddr, cdc)
	if recoveryJSON != nil && len(recoveryJSON) > 0 {
		if err = tssMgr.Recover(recoveryJSON); err != nil {
			panic(fmt.Errorf("unable to perform tss recovery: %v", err))
		}
	}

	query := "tm.event = 'Tx'"
	subscriber, err := client.Subscribe(context.Background(), "", query)
	if err != nil {
		panic(err)
	}
	out, err := client.Subscribe(context.Background(), "", "tm.event = 'NewBlockHeader'")
	if err != nil {
		panic(err)
	}

	js := []jobs.Job{

		Consume(subscriber, tssMgr),
		ConsumeH(out, tssMgr),
	}

	// errGroup runs async processes and cancels their context if ANY of them returns an error.
	// Here, we don't want to stop on errors, but simply log it and continue, so errGroup doesn't cut it
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

					if events[0].Events[0].Type == "keygen" {

						key := events[0].Events[0].Attributes[0].Key

						if key == "start" {

							if err := tssMgr.ProcessKeygenStart(events, e.Data.(tmtypes.EventDataTx).Height); err != nil {
								errChan <- err
							}

						}
						if key == "message" {
							
							if err := tssMgr.ProcessKeygenMsg(e2,e.Data.(tmtypes.EventDataTx).Height); err != nil {
								errChan <- err
							}
						}
						if key == "dispute" {
							if err := tssMgr.ProcessKeygenMsgDispute(events[0].Events); err != nil {
								errChan <- err
							}
						}
					}

				}()
			// default:
			// 	// Sleep for a short duration to yield CPU time
			// 	time.Sleep(10 * time.Millisecond)
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

					newBlockHeader := e.Data.(tmtypes.EventDataNewBlockHeader)

					height := newBlockHeader.Header.Height
					//fmt.Println(height)
					tssMgr.CheckTimeout(int(height))
				}()
			default:
				// Sleep for a short duration to yield CPU time
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}
func recovery(errChan chan<- error) {
	if r := recover(); r != nil {
		errChan <- fmt.Errorf("job panicked:%s", r)
	}
}

func createTSSMgr(broadcaster *broadcast.CosmosClient, cliCtx client.Context, dkgCfg config.ValdConfig, logger log.Logger, valAddr string, cdc *codec.LegacyAmino) *tss.Mgr {
	create := func() (*tss.Mgr, error) {
		
		conn, err := tss.Connect(dkgCfg.TssConfig.Host, dkgCfg.TssConfig.Port, dkgCfg.TssConfig.DialTimeout, logger)
		if err != nil {
			return nil, err
		}
		logger.Debug("successful connection to tofnd gRPC server")

		// creates clients to communicate with the external tofnd process service
		gg20client := tofnd.NewGG20Client(conn)
		//multiSigClient := tofnd.NewMultisigClient(conn)
		
		tssMgr := tss.NewMgr(gg20client, cliCtx, 2*time.Hour, valAddr, broadcaster, logger, cdc)

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
