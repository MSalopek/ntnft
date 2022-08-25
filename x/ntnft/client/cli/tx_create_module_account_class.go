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

func CmdCreateModuleAccountClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-module-account-class [name] [price] [module-name]",
		Short: "Broadcast message create-module-account-class",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argPrice := args[1]
			argModuleName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModuleAccountClass(
				clientCtx.GetFromAddress().String(),
				argName,
				argPrice,
				argModuleName,
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
