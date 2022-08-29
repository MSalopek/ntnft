package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"nt-nft/x/ntnft/types"
)

var _ = strconv.Itoa(0)

func CmdEditToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-token [token-id] [uri] [uri-hash] [data]",
		Short: "Broadcast message editToken",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenId := args[0]
			argUri := args[1]
			argUriHash := args[2]
			argData := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditToken(
				clientCtx.GetFromAddress().String(),
				argTokenId,
				argUri,
				argUriHash,
				argData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
