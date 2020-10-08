package phraseservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/danikarik/phraseservice/x/phraseservice/keeper"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

// Handle a message to delete name
func handleMsgDeletePhrase(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeletePhrase) (*sdk.Result, error) {
	if !k.PhraseExists(ctx, msg.Text) {
		return nil, types.ErrPhraseDoesNotExist
	}

	if !msg.Owner.Equals(k.GetPhraseOwner(ctx, msg.Text)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeletePhrase(ctx, msg.Text)
	return &sdk.Result{}, nil
}
