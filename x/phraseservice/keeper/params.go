package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

func (k Keeper) GetOwner(ctx sdk.Context, text string) sdk.AccAddress {
	phrase, _ := k.GetPhrase(ctx, text)
	return phrase.Owner
}

func (k Keeper) Exists(ctx sdk.Context, text string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PhrasePrefix + text))
}

func (k Keeper) SetName(ctx sdk.Context, text string, value string) {
	phrase, _ := k.GetPhrase(ctx, text)
	phrase.Text = value
	k.SetPhrase(ctx, phrase)
}

func (k Keeper) HasOwner(ctx sdk.Context, text string) bool {
	phrase, _ := k.GetPhrase(ctx, text)
	return !phrase.Owner.Empty()
}

func (k Keeper) SetOwner(ctx sdk.Context, text string, owner sdk.AccAddress) {
	phrase, _ := k.GetPhrase(ctx, text)
	phrase.Owner = owner
	k.SetPhrase(ctx, phrase)
}

func (k Keeper) GetBlock(ctx sdk.Context, text string) int64 {
	phrase, _ := k.GetPhrase(ctx, text)
	return phrase.Block
}

func (k Keeper) SetBlock(ctx sdk.Context, text string, block int64) {
	phrase, _ := k.GetPhrase(ctx, text)
	phrase.Block = block
	k.SetPhrase(ctx, phrase)
}

func (k Keeper) IsPhrasePresent(ctx sdk.Context, text string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(text))
}

func (k Keeper) GetPhrasesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.PhrasePrefix))
}
