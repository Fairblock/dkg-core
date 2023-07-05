package vald

import (
	"context"

	"fmt"
	"io"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	//broadcast "github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/broadcaster"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/config"
	"github.com/fairblock/dkg-core/cmd/dkgd/cmd/vald/tss"

	

	//"github.com/fairblock/dkg-core/x/tss/tofnd"
	tssTypes "github.com/fairblock/dkg-core/x/tss/types"
)

const (
	keyID      = "testkey"
	tokenDenom = "uaxl"
	minBalance = 5000000
	timeout    = time.Hour

	flagSkipTofnd       = "skip-tofnd"
	flagSkipBroadcaster = "skip-broadcaster"
	flagSkipOperator    = "skip-operator"
	flagOperatorAddr    = "operator-addr"
	flagTofndHost       = "tofnd-host"
	flagTofndPort       = "tofnd-port"
)

// GetHealthCheckCommand returns the command to execute a node health check
func GetHealthCheckCommand() *cobra.Command {
	var skipTofnd bool
	var skipBroadcaster bool
	var skipOperator bool

	cmd := &cobra.Command{
		Use: "health-check",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			serverCtx := server.GetServerContextFromCmd(cmd)
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithSimulateAndExecute(true)
			valdConf := config.DefaultValdConfig()
			if err := serverCtx.Viper.Unmarshal(&valdConf); err != nil {
				panic(err)
			}
			logger := serverCtx.Logger.With("module", "vald")
			ok := execCheck(txf, valdConf, context.Background(), clientCtx, serverCtx, logger, "tofnd", skipTofnd, checkTofnd)

			// enforce a non-zero exit code in case health checks fail without printing cobra output
			if !ok {
				os.Exit(1)
			}

			return nil
		},
	}

	defaultConf := tssTypes.DefaultConfig()
	cmd.PersistentFlags().String(flagTofndHost, defaultConf.Host, "host name for tss daemon")
	cmd.PersistentFlags().String(flagTofndPort, defaultConf.Port, "port for tss daemon")
	cmd.PersistentFlags().String(flagOperatorAddr, "", "operator address")
	cmd.PersistentFlags().BoolVar(&skipTofnd, flagSkipTofnd, false, "skip tofnd check")
	cmd.PersistentFlags().BoolVar(&skipBroadcaster, flagSkipBroadcaster, false, "skip broadcaster check")
	cmd.PersistentFlags().BoolVar(&skipOperator, flagSkipOperator, false, "skip operator check")

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

type checkCmd func(txf tx.Factory, dkgCfg config.ValdConfig, ctx context.Context, clientCtx client.Context, serverCtx *server.Context, logger log.Logger) error

func execCheck(txf tx.Factory, dkgCfg config.ValdConfig, ctx context.Context, clientCtx client.Context, serverCtx *server.Context, logger log.Logger, name string, skip bool, check checkCmd) bool {
	if skip {
		fmt.Printf("%s check: skipped\n", name)
		return true
	}

	err := check(txf, dkgCfg, ctx, clientCtx, serverCtx, logger)
	if err != nil {
		fmt.Printf("%s check: failed (%s)\n", name, err.Error())
		return false
	}

	fmt.Printf("%s check: passed\n", name)
	return true
}

func checkTofnd(txf tx.Factory, dkgCfg config.ValdConfig, ctx context.Context, clientCtx client.Context, serverCtx *server.Context, logger log.Logger) error {
	valdCfg := config.DefaultValdConfig()
	if err := serverCtx.Viper.Unmarshal(&valdCfg); err != nil {
		panic(err)
	}

	nopLogger := server.ZeroLogWrapper{Logger: zerolog.New(io.Discard)}

	_, err := tss.Connect(valdCfg.TssConfig.Host, valdCfg.TssConfig.Port, valdCfg.TssConfig.DialTimeout, nopLogger)
	if err != nil {
		return fmt.Errorf("failed to reach tofnd: %s", err.Error())
	}
	nopLogger.Debug("successful connection to tofnd gRPC server")

	//keygenStart := subscribe(tssTypes.EventTypeKeygen, tssTypes.ModuleName, tssTypes.AttributeValueStart)
	// creates client to communicate with the external tofnd process gg20 service
	// gg20client := tofnd.NewGG20Client(conn)

	// // bc := createBroadcaster(txf, dkgCfg, logger)
	// bc, err := broadcast.NewCosmosClient(
	// 	fmt.Sprintf(
	// 		"%s:%s",
	// 		"127.0.0.1",
	// 		"9090",
	// 	),
	// 	"86bc175431f38c7c1b021a0c1aedbd1696773501844607cb672f4d434467e669",
	// 	"fairyring",
	// )
	//tssMgr := tss.NewMgr(gg20client, clientCtx, 2*time.Hour, "valAddr", bc, logger, codec.NewLegacyAmino())

	// grpcCtx, cancel := context.WithTimeout(ctx, timeout)
	// defer cancel()

	// request := &tofnd.KeygenInit{
	// 	// we do not need to look for a key ID that exists to obtain a successful healthcheck,
	// 	// all we need to do is obtain err == nil && response != FAIL
	// 	// TODO: this kind of check should have its own dedicated GRPC
	// 	NewKeyUid: keyID,
	// 	PartyUids: []string{keyID},
	// 	MyPartyIndex: 2,
	// 	Threshold: 2,
	// }

	//err = tssMgr.ProcessKeygenStart([]tss.EventMsg{}, 0)

	if err != nil {
		return fmt.Errorf("failed to invoke tofnd grpc: %s", err.Error())
	}

	//panic(response)

	return nil
}
