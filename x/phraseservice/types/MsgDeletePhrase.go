package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeletePhrase{}

type MsgDeletePhrase struct {
	Owner sdk.AccAddress `json:"creator" yaml:"creator"`
	Text  string         `json:"text" yaml:"text"`
}

func NewMsgDeletePhrase(owner sdk.AccAddress, text string) MsgDeletePhrase {
	return MsgDeletePhrase{
		Owner: owner,
		Text:  text,
	}
}

func (msg MsgDeletePhrase) Route() string {
	return RouterKey
}

func (msg MsgDeletePhrase) Type() string {
	return "DeletePhrase"
}

func (msg MsgDeletePhrase) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgDeletePhrase) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeletePhrase) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}

	return nil
}
