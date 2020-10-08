package keeper

import (
	// this line is used by starport scaffolding # 1
	"fmt"
	"strings"

	"github.com/danikarik/phraseservice/x/phraseservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for phraseservice clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		fmt.Printf("path: %s\n", strings.Join(path, ", "))

		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListPhrase:
			return listPhraseByOwner(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown phraseservice query endpoint")
		}
	}
}
