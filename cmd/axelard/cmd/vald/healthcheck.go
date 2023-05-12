package vald

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
	broadcast "github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcaster"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/config"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/tss"
	"github.com/axelarnetwork/tm-events/events"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/x/snapshot/keeper"
	"github.com/axelarnetwork/axelar-core/x/snapshot/types"
	snapshotTypes "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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
			ok := execCheck(txf,valdConf, context.Background(), clientCtx, serverCtx,logger, "tofnd", skipTofnd, checkTofnd) 
			
			

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

type checkCmd func(txf tx.Factory, axelarCfg config.ValdConfig,ctx context.Context, clientCtx client.Context, serverCtx *server.Context,logger log.Logger) error

func execCheck(txf tx.Factory, axelarCfg config.ValdConfig,ctx context.Context, clientCtx client.Context, serverCtx *server.Context,logger log.Logger, name string, skip bool, check checkCmd) bool {
	if skip {
		fmt.Printf("%s check: skipped\n", name)
		return true
	}

	err := check(txf,axelarCfg ,ctx,clientCtx, serverCtx,logger)
	if err != nil {
		fmt.Printf("%s check: failed (%s)\n", name, err.Error())
		return false
	}

	fmt.Printf("%s check: passed\n", name)
	return true
}

func checkTofnd(txf tx.Factory, axelarCfg config.ValdConfig,ctx context.Context, clientCtx client.Context, serverCtx *server.Context,  logger log.Logger) error {
	valdCfg := config.DefaultValdConfig()
	if err := serverCtx.Viper.Unmarshal(&valdCfg); err != nil {
		panic(err)
	}

	nopLogger := server.ZeroLogWrapper{Logger: zerolog.New(io.Discard)}

	conn, err := tss.Connect(valdCfg.TssConfig.Host, valdCfg.TssConfig.Port, valdCfg.TssConfig.DialTimeout, nopLogger)
	if err != nil {
		return fmt.Errorf("failed to reach tofnd: %s", err.Error())
	}
	nopLogger.Debug("successful connection to tofnd gRPC server")
	
	//keygenStart := subscribe(tssTypes.EventTypeKeygen, tssTypes.ModuleName, tssTypes.AttributeValueStart)
	// creates client to communicate with the external tofnd process gg20 service
	gg20client := tofnd.NewGG20Client(conn)

	// bc := createBroadcaster(txf, axelarCfg, logger)
	bc,err := broadcast.NewCosmosClient(
		fmt.Sprintf(
			"%s:%s",
			"127.0.0.1",
			"9090",
		),
		"86bc175431f38c7c1b021a0c1aedbd1696773501844607cb672f4d434467e669",
		"fairyring",
	)
	tssMgr := tss.NewMgr(gg20client, clientCtx, 2*time.Hour, "valAddr",bc, logger,codec.NewLegacyAmino())

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

	err = tssMgr.ProcessKeygenStart(events.Event{})

	if err != nil {
		return fmt.Errorf("failed to invoke tofnd grpc: %s", err.Error())
	}

	//panic(response)

	return nil
}

func checkBroadcaster(ctx context.Context, clientCtx client.Context, serverCtx *server.Context) error {
	str := serverCtx.Viper.GetString(flagOperatorAddr)
	if str == "" {
		return fmt.Errorf("no operator address specified")
	}
	operator, err := sdk.ValAddressFromBech32(str)
	if err != nil {
		return err
	}

	bz, _, err := clientCtx.Query(fmt.Sprintf("custom/%s/%s/%s", snapshotTypes.QuerierRoute, keeper.QProxy, operator.String()))
	if err != nil {
		return err
	}

	reply := struct {
		Address string `json:"address"`
		Status  string `json:"status"`
	}{}
	json.Unmarshal(bz, &reply)

	broadcaster, err := sdk.AccAddressFromBech32(reply.Address)
	if err != nil {
		return err
	}

	if reply.Status != "active" {
		return fmt.Errorf("broadcaster for operator %s not active", operator.String())
	}

	queryClient := bankTypes.NewQueryClient(clientCtx)
	params := bankTypes.NewQueryBalanceRequest(broadcaster, tokenDenom)

	grpcCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, err := queryClient.Balance(grpcCtx, params)
	if err != nil {
		return err
	}

	if res.Balance.Amount.LTE(sdk.NewInt(minBalance)) {
		return fmt.Errorf("broadcaster does not have enough funds (minimum balance is %d%s)", minBalance, tokenDenom)
	}

	return nil
}

func checkOperator(_ context.Context, clientCtx client.Context, serverCtx *server.Context) error {
	addr := serverCtx.Viper.GetString(flagOperatorAddr)
	if addr == "" {
		return fmt.Errorf("no operator address specified")
	}

	bz, _, err := clientCtx.Query(fmt.Sprintf("custom/%s/%s", snapshotTypes.QuerierRoute, keeper.QValidators))
	if err != nil {
		return err
	}

	var resValidators types.QueryValidatorsResponse
	types.ModuleCdc.MustUnmarshalLengthPrefixed(bz, &resValidators)

	for _, v := range resValidators.Validators {
		if v.OperatorAddress == addr {
			if v.TssIllegibilityInfo.Jailed ||
				v.TssIllegibilityInfo.MissedTooManyBlocks ||
				v.TssIllegibilityInfo.NoProxyRegistered ||
				v.TssIllegibilityInfo.Tombstoned ||
				v.TssIllegibilityInfo.TssSuspended ||
				v.TssIllegibilityInfo.StaleTssHeartbeat {
				return fmt.Errorf("health check to operator %s failed due to the following issues: %v",
					addr, string(snapshotTypes.ModuleCdc.MustMarshalJSON(&v.TssIllegibilityInfo)))
			}
			return nil
		}
	}

	return fmt.Errorf("operator address %s not found amongst current set of validators", addr)
}
