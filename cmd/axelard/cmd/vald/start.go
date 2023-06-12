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

	broadcast "github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcaster"
	//axelarnet "github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	//"github.com/golang/protobuf/proto"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	//	"go.starlark.net/lib/proto"
	_ "google.golang.org/protobuf/proto"
	//tx2 "github.com/tendermint/tendermint/proto/tendermint/tx"
	//"github.com/axelarnetwork/tm-events/events"
	//	"github.com/axelarnetwork/tm-events/events"
	//"github.com/axelarnetwork/tm-events/events"
	tmEvents "github.com/axelarnetwork/tm-events/events"
	"github.com/axelarnetwork/tm-events/pubsub"
	"github.com/axelarnetwork/utils/jobs"

	//"github.com/axelarnetwork/utils/jobs"
	"github.com/cosmos/cosmos-sdk/client"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/config"

	"github.com/axelarnetwork/axelar-core/app"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/utils"

	//"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcaster"
	//"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/btc"
	//btcRPC "github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/btc/rpc"
	//"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/evm"
	//evmRPC "github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/evm/rpc"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/tss"

	//utils2 "github.com/axelarnetwork/axelar-core/utils"

	//btcTypes "github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	//evmTypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
)
// const addressForChain = "cosmos150lcfqj44zx8aljqn4za4pp2384k5gw3hpypm2" 
//go run cmd/axelard/main.go vald-start --validator-addr cosmos1exfcnjtc30msg2py3utlf0mmlq8ex32aadxlf3 --validator-key 85bc470c18f113a15384660980fc8e4000f9d5aacc129b02ef4851c4126d82bb
// const keyForTest = "00b183d4a1e6ba3fa5a036afabeb4644f1a24ad2b11cf3e6da2de96454c9fb8a" 
//go run cmd/axelard/main.go vald-start --validator-addr cosmos150lcfqj44zx8aljqn4za4pp2384k5gw3hpypm2 --validator-key 00b183d4a1e6ba3fa5a036afabeb4644f1a24ad2b11cf3e6da2de96454c9fb8a
// const addressForChain = "cosmos1exfcnjtc30msg2py3utlf0mmlq8ex32aadxlf3"
// const keyForTest = "85bc470c18f113a15384660980fc8e4000f9d5aacc129b02ef4851c4126d82bb"
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

func listen(ctx sdkClient.Context, txf tx.Factory, axelarCfg config.ValdConfig, valAddr string, valKey string,recoveryJSON []byte, stateSource ReadWriter, logger log.Logger) {
	encCfg := app.MakeEncodingConfig()
	cdc := encCfg.Amino
	// sender, err := ctx.Keyring.Key(ctx.From)
	// if err != nil {
	// 	panic(sdkerrors.Wrap(err, "failed to read broadcaster account info from keyring"))
	// }
	//fmt.Println(valAddr)
	//fmt.Println(valKey)
	ctx = ctx.
		WithFromAddress(sdk.AccAddress{}).
		WithFromName("seti")

	//bc := createBroadcaster(txf, axelarCfg, logger)

	//stateStore := NewStateStore(stateSource)
	//startBlock, err := stateStore.GetState()
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	startBlock = 0
	// }

	// tmClient, err := ctx.GetNode()
	// if err != nil {
	// 	panic(err)
	// }
	// // in order to subscribe to events, the client needs to be running
	// if !tmClient.IsRunning() {
	// 	if err := tmClient.Start(); err != nil {
	// 		panic(fmt.Errorf("unable to start client: %v", err))
	// 	}
	// }
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
	
	eventBus := createEventBus(client, 396620, logger)
	
	bc,err := broadcast.NewCosmosClient(
		fmt.Sprintf(
			"%s:%s",
			"127.0.0.1",
			"9090",
		),
		valKey,
		"dkg",
	)
	if err != nil{
		logger.Error(err.Error())
	}
	tssMgr := createTSSMgr(bc, ctx, axelarCfg, logger, valAddr, cdc)
	if recoveryJSON != nil && len(recoveryJSON) > 0 {
		if err = tssMgr.Recover(recoveryJSON); err != nil {
			panic(fmt.Errorf("unable to perform tss recovery: %v", err))
		}
	}

	// btcMgr := createBTCMgr(axelarCfg, ctx, bc, logger, cdc)
	// evmMgr := createEVMMgr(axelarCfg, ctx, bc, logger, cdc)

	// we have two processes listening to block headers
	blockHeaderForTSS := tmEvents.MustSubscribeBlockHeader(eventBus)
	// blockHeaderForStateUpdate := tmEvents.MustSubscribeBlockHeader(eventBus)

	// subscribe := func(eventType, module, action string) tmEvents.FilteredSubscriber {
	// 	return tmEvents.MustSubscribeWithAttributes(eventBus,
	// 		eventType, module, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: action})

	// }
	query := "tm.event = 'Tx'"
	subscriber , err := client.Subscribe(context.Background(), "", query)
	if err != nil {
		panic(err)
	}
	out, err := client.Subscribe(context.Background(), "", "tm.event = 'NewBlockHeader'")
	if err != nil {
		panic(err)
	}
	//keygenStart := subscribe(tssTypes.EventTypeKeygen, tssTypes.ModuleName, tssTypes.AttributeValueStart)
	//queryHeartBeat := createNewBlockEventQuery(tssTypes.EventTypeHeartBeat, tssTypes.ModuleName, tssTypes.AttributeValueSend)
	// heartbeat, err := tmEvents.Subscribe(eventBus, queryHeartBeat)
	// if err != nil {
	// 	panic(fmt.Errorf("unable to subscribe with ack event query: %v", err))
	// }

	//keygenStart := subscribe(tssTypes.EventTypeKeygen, tssTypes.ModuleName, tssTypes.AttributeValueStart)

	//querySign := createNewBlockEventQuery(tssTypes.EventTypeSign, tssTypes.ModuleName, tssTypes.AttributeValueStart)
	// signStart, err := tmEvents.Subscribe(eventBus, querySign)
	// if err != nil {
	// 	panic(fmt.Errorf("unable to subscribe with sign event query: %v", err))
	// }

//	keygenMsg := subscribe(tssTypes.EventTypeKeygen, "dkg", tssTypes.AttributeValueMsg)
	//	signMsg := subscribe(tssTypes.EventTypeSign, tssTypes.ModuleName, tssTypes.AttributeValueMsg)

	// btcConf := subscribe(btcTypes.EventTypeOutpointConfirmation, btcTypes.ModuleName, btcTypes.AttributeValueStart)

	// evmNewChain := subscribe(evmTypes.EventTypeNewChain, evmTypes.ModuleName, evmTypes.AttributeValueUpdate)
	// evmChainConf := subscribe(evmTypes.EventTypeChainConfirmation, evmTypes.ModuleName, evmTypes.AttributeValueStart)
	// evmDepConf := subscribe(evmTypes.EventTypeDepositConfirmation, evmTypes.ModuleName, evmTypes.AttributeValueStart)
	// evmTokConf := subscribe(evmTypes.EventTypeTokenConfirmation, evmTypes.ModuleName, evmTypes.AttributeValueStart)
	// evmTraConf := subscribe(evmTypes.EventTypeTransferKeyConfirmation, evmTypes.ModuleName, evmTypes.AttributeValueStart)

	eventCtx, cancelEventCtx := context.WithCancel(context.Background())
	// stop the jobs if process gets interrupted/terminated
	cleanupCommands = append(cleanupCommands, func() {
		logger.Info("stopping listening for blocks...")
		blockHeaderForTSS.Close()
		logger.Info("block listener stopped")
		logger.Info("stop listening for events...")
		cancelEventCtx()
		<-eventBus.Done()
		logger.Info("event listener stopped")
	})

	fetchEvents := func(errChan chan<- error) { errChan <- <-eventBus.FetchEvents(eventCtx) }
	js := []jobs.Job{
		fetchEvents,
		// tmEvents.Consume(blockHeaderForStateUpdate, tmEvents.OnlyBlockHeight(stateStore.SetState)),
		// tmEvents.Consume(blockHeaderForTSS, tmEvents.OnlyBlockHeight(func(height int64) error {
		// 	tssMgr.ProcessNewBlockHeader(height)
		// 	return nil
		// })),
		//  tmEvents.Consume(heartbeat, tssMgr.ProcessHeartBeatEvent),
		// tmEvents.Consume(keygenStart, tssMgr.ProcessKeygenStart),
		Consume(subscriber, tssMgr),
		ConsumeH(out, tssMgr),
		// tmEvents.Consume(signStart, tssMgr.ProcessSignStart),
		// tmEvents.Consume(signMsg, tssMgr.ProcessSignMsg),
		// tmEvents.Consume(btcConf, btcMgr.ProcessConfirmation),
		// tmEvents.Consume(evmNewChain, evmMgr.ProcessNewChain),
		// tmEvents.Consume(evmChainConf, evmMgr.ProcessChainConfirmation),
		// tmEvents.Consume(evmDepConf, evmMgr.ProcessDepositConfirmation),
		// tmEvents.Consume(evmTokConf, evmMgr.ProcessTokenConfirmation),
		// tmEvents.Consume(evmTraConf, evmMgr.ProcessTransferOwnershipConfirmation),
	}

	// errGroup runs async processes and cancels their context if ANY of them returns an error.
	// Here, we don't want to stop on errors, but simply log it and continue, so errGroup doesn't cut it
	logErr := func(err error) { logger.Error(err.Error()) }

	mgr := jobs.NewMgr(logErr)

	mgr.AddJobs(js...)

	mgr.Wait()

}
//go run cmd/axelard/main.go vald-start --validator-addr cosmos1u0sv2a225ualg7t35pxux8453t4d3huv8966p9 --validator-key 28c1d3cb9c1ce8fb04f26f55fc859d9ed262feb1b1fc5dc06ad502c947a932e7
//go run cmd/axelard/main.go vald-start --validator-addr cosmos1gh7gxcfcmvxq6vysh0dtey26erpaawjnfpc289 --validator-key  58ef035d49e62e6a214cc8c1501db1a3218e777fe3b623cbec9c1465ed710b78
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
					fmt.Println("event : ", d)
					e2 := e.Data.(tmtypes.EventDataTx).Result.Events
					
					var events []tss.EventMsg
					if err := json.Unmarshal([]byte(d), &events); err != nil {
						errChan <- err
					}
					//fmt.Println(e.Data)
					if events[0].Events[0].Type == "keygen"{

						key := events[0].Events[0].Attributes[0].Key
						
						if key == "start"{
							
							if err := tssMgr.ProcessKeygenStart(events, e.Data.(tmtypes.EventDataTx).Height); err != nil {
								errChan <- err
							}
							
						}
						if key == "message"{
							
							if err := tssMgr.ProcessKeygenMsg(e2); err != nil {
								errChan <- err
							}
						}
						if key == "dispute"{
							if err := tssMgr.ProcessKeygenMsgDispute(events[0].Events); err != nil {
								errChan <- err
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
					
					newBlockHeader := e.Data.(tmtypes.EventDataNewBlockHeader)

					height := newBlockHeader.Header.Height
					//fmt.Println(height)
					tssMgr.CheckTimeout(int(height));
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

func createNewBlockEventQuery(eventType, module, action string) tmEvents.Query {
	return tmEvents.Query{
		TMQuery: tmEvents.NewBlockHeaderEventQuery(eventType).MatchModule(module).MatchAction(action).Build(),
		Predicate: func(e tmEvents.Event) bool {
			return e.Type == eventType && e.Attributes[sdk.AttributeKeyModule] == module && e.Attributes[sdk.AttributeKeyAction] == action
		},
	}
}

func createEventBus(client rpcclient.Client, startBlock int64, logger log.Logger) *tmEvents.Bus {
	notifier := tmEvents.NewBlockNotifier(tmEvents.NewBlockClient(client), logger).StartingAt(startBlock)
	return tmEvents.NewEventBus(tmEvents.NewBlockSource(client, notifier), pubsub.NewBus, logger)
}

// func createBroadcaster(txf tx.Factory, axelarCfg config.ValdConfig, logger log.Logger) broadcast.CosmosClient {
// 	pipeline := broadcaster.NewPipelineWithRetry(10000, axelarCfg.MaxRetries, utils2.LinearBackOff(axelarCfg.MinTimeout), logger)
// 	return broadcaster.NewBroadcaster(txf, pipeline, logger)
// }

func createTSSMgr(broadcaster *broadcast.CosmosClient, cliCtx client.Context, axelarCfg config.ValdConfig, logger log.Logger, valAddr string, cdc *codec.LegacyAmino) *tss.Mgr {
	create := func() (*tss.Mgr, error) {
		conn, err := tss.Connect(axelarCfg.TssConfig.Host, axelarCfg.TssConfig.Port, axelarCfg.TssConfig.DialTimeout, logger)
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

// func createBTCMgr(axelarCfg config.ValdConfig, cliCtx client.Context, b broadcast.CosmosClient, logger log.Logger, cdc *codec.LegacyAmino) *btc.Mgr {
// 	rpc, err := btcRPC.NewRPCClient(axelarCfg.BtcConfig, logger)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		panic(err)
// 	}
// 	// clean up btcRPC connection on process shutdown
// 	cleanupCommands = append(cleanupCommands, rpc.Shutdown)

// 	logger.Info("Successfully connected to Bitcoin bridge ")

// 	btcMgr := btc.NewMgr(rpc, cliCtx, b, logger, cdc)
// 	return btcMgr
// }

// func createEVMMgr(axelarCfg config.ValdConfig, cliCtx client.Context, b broadcast.CosmosClient, logger log.Logger, cdc *codec.LegacyAmino) *evm.Mgr {
// 	rpcs := make(map[string]evmRPC.Client)

// 	for _, evmChainConf := range axelarCfg.EVMConfig {
// 		if !evmChainConf.WithBridge {
// 			continue
// 		}

// 		if _, found := rpcs[strings.ToLower(evmChainConf.Name)]; found {
// 			msg := fmt.Errorf("duplicate bridge configuration found for EVM chain %s", evmChainConf.Name)
// 			logger.Error(msg.Error())
// 			panic(msg)
// 		}

// 		rpc, err := evmRPC.NewClient(evmChainConf.RPCAddr)
// 		if err != nil {
// 			logger.Error(err.Error())
// 			panic(err)
// 		}
// 		// clean up evmRPC connection on process shutdown
// 		cleanupCommands = append(cleanupCommands, rpc.Close)

// 		rpcs[strings.ToLower(evmChainConf.Name)] = rpc
// 		logger.Info(fmt.Sprintf("Successfully connected to EVM bridge for chain %s", evmChainConf.Name))
// 	}

// 	evmMgr := evm.NewMgr(rpcs, cliCtx, b, logger, cdc)
// 	return evmMgr
// }

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
