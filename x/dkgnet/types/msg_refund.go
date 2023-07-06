package types

import (
	//"crypto/rand"
	//	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dkgnet "github.com/fairblock/dkg-core/x/dkgnet/exported"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//proto "github.com/gogo/protobuf/proto"
)

// NewMsgRefundMsgRequest creates a message of type MsgRefundMsgRequest
func NewMsgRefundMsgRequest(creator string, sender sdk.AccAddress, innerMessage sdk.Msg) *MsgRefundMsgRequest {

	messageAny, err := cdctypes.NewAnyWithValue(innerMessage)
	if err != nil {
		panic(err)
	}

	return &MsgRefundMsgRequest{
		Creator:      creator,
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

// GetInnerMessage unwrap the inner message
// func (m MsgRefundMsgRequest) GetInnerMessage() dkgnet.Refundable {
// 	innerMsg := new(dkgnet.)
// 	innerMsg.Unmarshal([]byte(m.InnerMessage))
// 	// innerMsg, ok := m.InnerMessage.(dkgnet.Refundable)
// 	if !ok {
// 		return nil
// 	}
// 	return innerMsg
// }

func (m MsgRefundMsgRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m MsgRefundMsgRequest) Type() string {
	return "MsgRefundMsgRequest"
}

// ValidateBasic executes a stateless message validation
func (m MsgRefundMsgRequest) ValidateBasic() error {
	// if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
	// 	return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	// }

	// if m.InnerMessage == nil {
	// 	return fmt.Errorf("missing inner message")
	// }

	// innerMessage := m.GetInnerMessage()
	// if innerMessage == nil {
	// 	return fmt.Errorf("invalid inner message")
	// }

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m MsgRefundMsgRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the set of signers for this message
func (m MsgRefundMsgRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (m MsgRefundMsgRequest) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	if m.InnerMessage != nil {
		var refundableMsg dkgnet.Refundable
		return unpacker.UnpackAny(m.InnerMessage, &refundableMsg)
	}
	return nil
}

// GetInnerMessage unwrap the inner message
func (m MsgRefundMsgRequest) GetInnerMessage() dkgnet.Refundable {
	innerMsg, ok := m.InnerMessage.GetCachedValue().(dkgnet.Refundable)
	if !ok {
		return nil
	}
	return innerMsg
}

func (m MsgFileDispute) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Dispute.AddressOfAccuser}
}

func (msg *MsgFileDispute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFileDispute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
func (m MsgFileDispute) Type() string {
	return "MsgFileDispute"
}

func (m MsgKeygenResult) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Creator)}
}

func (msg *MsgKeygenResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgKeygenResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
func (m MsgKeygenResult) Type() string {
	return "MsgKeygenResult"
}