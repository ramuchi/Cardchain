package cli

import (
	"github.com/DecentralCardGame/cardobject/cardobject"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetCardRarity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-card-rarity [card-id] [set-id] [rarity]",
		Short: "Broadcast message SetCardRarity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argSetId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argRarity := args[2]

			rarity := cardobject.Rarity(argRarity)
			err = rarity.ValidateType(cardobject.Card{})
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetCardRarity(
				clientCtx.GetFromAddress().String(),
				argCardId,
				argSetId,
				argRarity,
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
