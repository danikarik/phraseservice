package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
)

func phraseKey(text string) []byte {
	return []byte(types.PhrasePrefix + text)
}

// CreatePhrase creates a phrase
func (k Keeper) CreatePhrase(ctx sdk.Context, phrase types.Phrase) {
	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(phrase)
	store.Set(phraseKey(phrase.Text), value)
}

// GetPhrase returns the phrase information
func (k Keeper) GetPhrase(ctx sdk.Context, text string) (types.Phrase, error) {
	store := ctx.KVStore(k.storeKey)

	var phrase types.Phrase
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(phraseKey(text)), &phrase)
	if err != nil {
		return phrase, err
	}

	return phrase, nil
}

// SetPhrase sets a phrase
func (k Keeper) SetPhrase(ctx sdk.Context, phrase types.Phrase) {
	store := ctx.KVStore(k.storeKey)
	store.Set(phraseKey(phrase.Text), k.cdc.MustMarshalBinaryLengthPrefixed(phrase))
}

// DeletePhrase deletes a phrase
func (k Keeper) DeletePhrase(ctx sdk.Context, text string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(phraseKey(text))
}

//
// Functions used by querier
//

func listPhraseByOwner(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	var phrases []types.Phrase

	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PhrasePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var phrase types.Phrase
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &phrase)

		requester, err := sdk.AccAddressFromBech32(path[1])
		if err != nil {
			return nil, err
		}

		if phrase.Owner.Equals(requester) {
			phrases = append(phrases, phrase)
		}
	}

	res := codec.MustMarshalJSONIndent(k.cdc, phrases)
	return res, nil
}

// GetPhraseOwner gets owner of the item.
func (k Keeper) GetPhraseOwner(ctx sdk.Context, text string) sdk.AccAddress {
	phrase, err := k.GetPhrase(ctx, text)
	if err != nil {
		return nil
	}

	return phrase.Owner
}

// PhraseExists checks if the key exists in the store.
func (k Keeper) PhraseExists(ctx sdk.Context, text string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(phraseKey(text))
}
