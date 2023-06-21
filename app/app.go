package app

import (
	"encoding/json"
	"io"
	stdlog "log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"

	//"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/modules/core"

	//porttypes "github.com/cosmos/ibc-go/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/modules/core/keeper"

	dkgParams "github.com/fairblock/dkg-core/app/params"
	//"github.com/fairblock/dkg-core/x/ante"
	//"github.com/fairblock/dkg-core/x/dkgnet"
	
	dkgTypes "github.com/fairblock/dkg-core/x/dkgnet/types"
	//	"github.com/fairblock/dkg-core/x/bitcoin"
	//btcKeeper "github.com/fairblock/dkg-core/x/bitcoin/keeper"
	//btcTypes "github.com/fairblock/dkg-core/x/bitcoin/types"
	//	"github.com/fairblock/dkg-core/x/evm"
	//	evmKeeper "github.com/fairblock/dkg-core/x/evm/keeper"
	//evmTypes "github.com/fairblock/dkg-core/x/evm/types"
	//"github.com/fairblock/dkg-core/x/nexus"
	//nexusKeeper "github.com/fairblock/dkg-core/x/nexus/keeper"
	//nexusTypes "github.com/fairblock/dkg-core/x/nexus/types"
	//"github.com/fairblock/dkg-core/x/reward"
	//rewardKeeper "github.com/fairblock/dkg-core/x/reward/keeper"
	//rewardTypes "github.com/fairblock/dkg-core/x/reward/types"
	//"github.com/fairblock/dkg-core/x/snapshot"
	//snapKeeper "github.com/fairblock/dkg-core/x/snapshot/keeper"
	//snapTypes "github.com/fairblock/dkg-core/x/snapshot/types"
	//"github.com/fairblock/dkg-core/x/tss"
	//tssKeeper "github.com/fairblock/dkg-core/x/tss/keeper"
	tssTypes "github.com/fairblock/dkg-core/x/tss/types"
	//	"github.com/fairblock/dkg-core/x/vote"
	//voteKeeper "github.com/fairblock/dkg-core/x/vote/keeper"
	//voteTypes "github.com/fairblock/dkg-core/x/vote/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"

	//	authAnte "github.com/cosmos/cosmos-sdk/x/auth/ante"
	//authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	//authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"

	//bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"

	//	"github.com/cosmos/cosmos-sdk/x/feegrant"
	//	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"

	//	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"

	//	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	//	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	//	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Name is the name of the application
const Name = "dkg"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		vesting.AppModuleBasic{},
		ibc.AppModuleBasic{},
		transfer.AppModuleBasic{},

	//	tss.AppModuleBasic{},
	// vote.AppModuleBasic{},
	// bitcoin.AppModuleBasic{},
	// evm.AppModuleBasic{},
	// snapshot.AppModuleBasic{},
	// nexus.AppModuleBasic{},
	// dkgnet.AppModuleBasic{},
	// reward.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		dkgTypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
		//rewardTypes.ModuleName:         {authtypes.Minter},
	}
)

var (
	_ servertypes.Application = (*DKGApp)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Println("Failed to get home dir %2", err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}


type DKGApp struct {
	*baseapp.BaseApp

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// necessery keepers for export
	stakingKeeper    stakingkeeper.Keeper
	crisisKeeper     crisiskeeper.Keeper
	distrKeeper      distrkeeper.Keeper
	slashingKeeper   slashingkeeper.Keeper
	ibcKeeper        *ibckeeper.Keeper
	evidenceKeeper   evidencekeeper.Keeper
	transferKeeper   ibctransferkeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	mm           *module.Manager
	paramsKeeper paramskeeper.Keeper
}


func NewDKGApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig dkgParams.EncodingConfig,
	appOpts servertypes.AppOptions, baseAppOptions ...func(*bam.BaseApp)) *DKGApp {

	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		evidencetypes.StoreKey,
		ibchost.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,

		// voteTypes.StoreKey,
		// btcTypes.StoreKey,
		// evmTypes.StoreKey,
		// snapTypes.StoreKey,
		tssTypes.StoreKey,
		// nexusTypes.StoreKey,
		// dkgTypes.StoreKey,
		// rewardTypes.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	var app = &DKGApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	paramsK := initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	app.paramsKeeper = paramsK
	// set the BaseApp's parameter store
	bApp.SetParamStore(app.getSubspace(baseapp.Paramspace))

	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		tssTypes.ModuleName,
		//btcTypes.ModuleName,
		// evmTypes.ModuleName,
	//	nexusTypes.ModuleName,
	//	rewardTypes.ModuleName,
	)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibchost.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,

		// snapTypes.ModuleName,
		tssTypes.ModuleName,
		// btcTypes.ModuleName,
		// evmTypes.ModuleName,
		// nexusTypes.ModuleName,
		// voteTypes.ModuleName,
		// dkgTypes.ModuleName,
		// rewardTypes.ModuleName,
	)

	// app.mm.RegisterInvariants(&crisisK)
	// app.mm.RegisterInvariants(&crisisK)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), legacyAmino)
	app.mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The baseAnteHandler handles signature verification and transaction pre-processing
	// baseAnteHandler, err := authAnte.NewAnteHandler(
	// 	authAnte.HandlerOptions{
	// 		AccountKeeper:   accountK,
	// 		BankKeeper:      bankK,
	// 		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	// 		FeegrantKeeper:  feegrantK,
	// 		SigGasConsumer:  authAnte.DefaultSigVerificationGasConsumer,
	// 	},
	// )

	// if err != nil {
	// 	panic(err)
	// }



	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

func initParamsKeeper(appCodec codec.Codec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable())

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)

	//paramsKeeper.Subspace(snapTypes.ModuleName)
	paramsKeeper.Subspace(tssTypes.ModuleName)
	//paramsKeeper.Subspace(btcTypes.ModuleName)
	// paramsKeeper.Subspace(nexusTypes.ModuleName)
	paramsKeeper.Subspace(dkgTypes.ModuleName)
	//paramsKeeper.Subspace(rewardTypes.ModuleName)

	return paramsKeeper
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// InitChainer handles the chain initialization from a genesis file
func (app *DKGApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// BeginBlocker calls the BeginBlock() function of every module at the beginning of a new block
func (app *DKGApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker calls the EndBlock() function of every module at the end of a block
func (app *DKGApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// LoadHeight loads the application version at a given height. It will panic if called
// more than once on a running baseapp.
func (app *DKGApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// LegacyAmino returns DKGApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DKGApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns DKGApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DKGApp) AppCodec() codec.Codec {
	return app.appCodec
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *DKGApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *DKGApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	//authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *DKGApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *DKGApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

func (app *DKGApp) getSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.paramsKeeper.GetSubspace(moduleName)
	return subspace
}
