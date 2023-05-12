package types

import (
	//"crypto/rand"
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//proto "github.com/gogo/protobuf/proto"

	axelarnet "github.com/axelarnetwork/axelar-core/x/axelarnet/exported"
)
const addressForChain = "cosmos150lcfqj44zx8aljqn4za4pp2384k5gw3hpypm2"
const keyForTest = "00b183d4a1e6ba3fa5a036afabeb4644f1a24ad2b11cf3e6da2de96454c9fb8a"
// NewRefundMsgRequest creates a message of type RefundMsgRequest
func NewRefundMsgRequest(sender sdk.AccAddress, innerMessage sdk.Msg) *RefundMsgRequest {

	messageAny, err := cdctypes.NewAnyWithValue(innerMessage)
	if err != nil {
		panic(err)
	}

	return &RefundMsgRequest{
		Creator:      addressForChain,
		Sender:       sender,
		InnerMessage: messageAny,
	}
}

// func NewAnyWithValue(v proto.Message) (*cdctypes.Any, error) {
// 	if v == nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrPackAny, "Expecting non nil value to create a new Any")
// 	}
// // 	buf := make([]byte, 128)
// // // then we can call rand.Read.
// // _, err := rand.Read(buf)
// // _ = err
// // 	bz:= buf
// 	bz, err := proto.Marshal(v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Println("_------------------------------------------------------")
// 	// fmt.Println(proto.MessageName(v))
// 	return &cdctypes.Any{
// 		TypeUrl:     "/dkg.dkg.MsgRefundMsgRequest",
// 		Value:       bz,
// 		//cachedValue: v,
// 	}, nil
// }


// Route returns the route for this message
func (m RefundMsgRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m RefundMsgRequest) Type() string {
	return "RefundMsgRequest"
}

// ValidateBasic executes a stateless message validation
func (m RefundMsgRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if m.InnerMessage == nil {
		return fmt.Errorf("missing inner message")
	}

	innerMessage := m.GetInnerMessage()
	if innerMessage == nil {
		return fmt.Errorf("invalid inner message")
	}

	return innerMessage.ValidateBasic()
}

// GetSignBytes returns the message bytes that need to be signed
func (m RefundMsgRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the set of signers for this message
func (m RefundMsgRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (m RefundMsgRequest) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	if m.InnerMessage != nil {
		var refundableMsg axelarnet.Refundable
		return unpacker.UnpackAny(m.InnerMessage, &refundableMsg)
	}
	return nil
}

// GetInnerMessage unwrap the inner message
func (m RefundMsgRequest) GetInnerMessage() axelarnet.Refundable {
	innerMsg, ok := m.InnerMessage.GetCachedValue().(axelarnet.Refundable)
	if !ok {
		return nil
	}
	return innerMsg
}
