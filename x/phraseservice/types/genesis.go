package types

import "fmt"

// GenesisState - all phraseservice state that must be provided at genesis
type GenesisState struct {
	Phrases []Phrase `json:"phrases"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{Phrases: nil}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{Phrases: make([]Phrase, 0)}
}

// ValidateGenesis validates the phraseservice genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Phrases {
		if record.Owner == nil {
			return fmt.Errorf("invalid Phrase: Text: %s. Error: Missing Owner", record.Text)
		}
		if record.Text == "" {
			return fmt.Errorf("invalid Phrase: Owner: %s. Error: Missing Text", record.Owner)
		}
		if record.Block == 0 {
			return fmt.Errorf("invalid Phrase: Text: %s. Error: Missing Block", record.Text)
		}
	}

	return nil
}
