package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var layout = `
Owner: %s
Value: %s
Block: %d
`

// Phrase represents each copyrighted text.
type Phrase struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Text  string         `json:"text" yaml:"text"`
	Block int64          `json:"block" yaml:"block"`
}

func (p Phrase) String() string {
	return fmt.Sprintf(layout, p.Owner, p.Text, p.Block)
}
