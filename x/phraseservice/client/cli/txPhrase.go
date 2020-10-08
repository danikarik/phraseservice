package cli

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/danikarik/phraseservice/x/phraseservice/types"
	"github.com/spf13/cobra"
)

func GetCmdCreatePhrase(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-phrase [text] [block]",
		Short: "Creates a new phrase",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				text       = args[0]
				block, err = strconv.ParseInt(args[1], 10, 64)
			)

			if err != nil {
				return fmt.Errorf("cannot parse block value: %w", err)
			}

			var (
				ctx     = context.NewCLIContext().WithCodec(cdc)
				buf     = bufio.NewReader(cmd.InOrStdin())
				builder = auth.NewTxBuilderFromCLI(buf).WithTxEncoder(utils.GetTxEncoder(cdc))
			)

			msg := types.NewMsgCreatePhrase(ctx.GetFromAddress(), text, block)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, builder, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetPhrase(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-phrase [text] [block]",
		Short: "Set a new phrase",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				text       = args[0]
				block, err = strconv.ParseInt(args[1], 10, 64)
			)

			if err != nil {
				return fmt.Errorf("cannot parse block value: %w", err)
			}

			var (
				ctx     = context.NewCLIContext().WithCodec(cdc)
				buf     = bufio.NewReader(cmd.InOrStdin())
				builder = auth.NewTxBuilderFromCLI(buf).WithTxEncoder(utils.GetTxEncoder(cdc))
			)

			msg := types.NewMsgSetPhrase(ctx.GetFromAddress(), text, block)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, builder, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeletePhrase(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-phrase [text]",
		Short: "Delete a new phrase by text",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				ctx     = context.NewCLIContext().WithCodec(cdc)
				buf     = bufio.NewReader(cmd.InOrStdin())
				builder = auth.NewTxBuilderFromCLI(buf).WithTxEncoder(utils.GetTxEncoder(cdc))
			)

			msg := types.NewMsgDeletePhrase(ctx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, builder, []sdk.Msg{msg})
		},
	}
}
