package phraseservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/danikarik/phraseservice/x/phraseservice/keeper"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	for _, record := range data.Phrases {
		k.SetPhrase(ctx, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	var records []types.Phrase

	iterator := k.GetPhrasesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		text := string(iterator.Key())
		phrase, _ := k.GetPhrase(ctx, text)
		records = append(records, phrase)
	}

	return types.GenesisState{Phrases: records}
}
