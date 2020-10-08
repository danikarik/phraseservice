package phraseservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/danikarik/phraseservice/x/phraseservice/keeper"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

func handleMsgCreatePhrase(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePhrase) (*sdk.Result, error) {
	var phrase = types.Phrase{
		Owner: msg.Owner,
		Text:  msg.Text,
		Block: msg.Block,
	}

	if k.PhraseExists(ctx, msg.Text) {
		return nil, types.ErrPhraseIsTaken
	}

	k.CreatePhrase(ctx, phrase)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
