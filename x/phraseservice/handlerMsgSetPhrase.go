package phraseservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/danikarik/phraseservice/x/phraseservice/keeper"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

func handleMsgSetPhrase(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetPhrase) (*sdk.Result, error) {
	var phrase = types.Phrase{
		Owner: msg.Owner,
		Text:  msg.Text,
		Block: msg.Block,
	}

	if !msg.Owner.Equals(k.GetPhraseOwner(ctx, msg.Text)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPhrase(ctx, phrase)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
