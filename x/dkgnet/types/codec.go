package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/exported"
)

// RegisterLegacyAminoCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&LinkRequest{}, "dkgnet/Link", nil)
	cdc.RegisterConcrete(&ConfirmDepositRequest{}, "dkgnet/ConfirmDeposit", nil)
	cdc.RegisterConcrete(&ExecutePendingTransfersRequest{}, "dkgnet/ExecutePendingTransfers", nil)
	cdc.RegisterConcrete(&RegisterIBCPathRequest{}, "dkgnet/RegisterIBCPath", nil)
	cdc.RegisterConcrete(&AddCosmosBasedChainRequest{}, "dkgnet/AddCosmosBasedChain", nil)
	cdc.RegisterConcrete(&RegisterAssetRequest{}, "dkgnet/RegisterAsset", nil)
	cdc.RegisterConcrete(&MsgRefundMsgRequest{}, "dkgnet/MsgRefundMsgRequest", nil)
	cdc.RegisterConcrete(&RouteIBCTransfersRequest{}, "dkgnet/RouteIBCTransfers", nil)
	cdc.RegisterConcrete(&RegisterFeeCollectorRequest{}, "dkgnet/RegisterFeeCollector", nil)
	cdc.RegisterConcrete(&MsgFileDispute{}, "dkg/FileDispute", nil)
}

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&LinkRequest{},
		&ConfirmDepositRequest{},
		&ExecutePendingTransfersRequest{},
		&RegisterIBCPathRequest{},
		&AddCosmosBasedChainRequest{},
		&RegisterAssetRequest{},
		&MsgRefundMsgRequest{},
		&RouteIBCTransfersRequest{},
		&RegisterFeeCollectorRequest{},
	)
	registry.RegisterInterface("dkgnet.v1beta1.Refundable",
		(*dkgnet.Refundable)(nil))
}

var amino = codec.NewLegacyAmino()

// ModuleCdc defines the module codec
var ModuleCdc = codec.NewAminoCodec(amino)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
