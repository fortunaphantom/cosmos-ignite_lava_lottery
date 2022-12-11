package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"lavalottery/x/lavalottery/types"
)

var _ = strconv.Itoa(0)

func CmdSendTicket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-ticket [fee] [bet]",
		Short: "Broadcast message send-ticket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFee := args[0]
			argBet := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendTicket(
				clientCtx.GetFromAddress().String(),
				argFee,
				argBet,
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
