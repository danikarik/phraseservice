package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
	"github.com/spf13/cobra"
)

func GetCmdListPhrase(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-phrase [owner]",
		Short: "list all phrase",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListPhrase, args[1]), nil)
			if err != nil {
				fmt.Printf("could not list Phrase\n%s\n", err.Error())
				return nil
			}

			var out []types.Phrase
			cdc.MustUnmarshalJSON(res, &out)
			return ctx.PrintOutput(out)
		},
	}
}
