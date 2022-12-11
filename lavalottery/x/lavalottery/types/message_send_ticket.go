package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendTicket = "send_ticket"

var _ sdk.Msg = &MsgSendTicket{}

func NewMsgSendTicket(creator string, fee string, bet string) *MsgSendTicket {
	return &MsgSendTicket{
		Creator: creator,
		Fee:     fee,
		Bet:     bet,
	}
}

func (msg *MsgSendTicket) Route() string {
	return RouterKey
}

func (msg *MsgSendTicket) Type() string {
	return TypeMsgSendTicket
}

func (msg *MsgSendTicket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendTicket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
