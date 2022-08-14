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

func CmdCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-class [name] [uri] [uri-hash] [data] [price]",
		Short: "Broadcast message create-class",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argUri := args[1]
			argUriHash := args[2]
			argData := args[3]
			argPrice := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClass(
				clientCtx.GetFromAddress().String(),
				argName,
				argUri,
				argUriHash,
				argData,
				argPrice,
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
