package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPhrase{}

type MsgSetPhrase struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Text  string         `json:"text" yaml:"text"`
	Block int64          `json:"block" yaml:"block"`
}

func NewMsgSetPhrase(owner sdk.AccAddress, text string, block int64) MsgSetPhrase {
	return MsgSetPhrase{
		Owner: owner,
		Text:  text,
		Block: block,
	}
}

func (msg MsgSetPhrase) Route() string {
	return RouterKey
}

func (msg MsgSetPhrase) Type() string {
	return "SetPhrase"
}

func (msg MsgSetPhrase) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSetPhrase) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetPhrase) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}

	if msg.Text == "" || msg.Block == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "text and/or block cannot be empty")
	}

	return nil
}
