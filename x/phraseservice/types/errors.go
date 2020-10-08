package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrPhraseDoesNotExist = sdkerrors.Register(ModuleName, 1, "phrase does not exist")
	ErrPhraseIsTaken      = sdkerrors.Register(ModuleName, 2, "phrase is taken")
)
