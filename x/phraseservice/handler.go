package phraseservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/danikarik/phraseservice/x/phraseservice/keeper"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreatePhrase:
			return handleMsgCreatePhrase(ctx, k, msg)
		case types.MsgSetPhrase:
			return handleMsgSetPhrase(ctx, k, msg)
		case types.MsgDeletePhrase:
			return handleMsgDeletePhrase(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
