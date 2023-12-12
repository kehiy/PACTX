package commands

import (
	"context"
	"fmt"

	"github.com/kehiy/PACTX/lib"
	cobra "github.com/spf13/cobra"
)

func BuildFeeCommands(parentCmd *cobra.Command) {
	feeCmd := &cobra.Command{
		Use:   "fee",
		Short: "fee utils.",
	}
	BuildCalcCommand(feeCmd)

	parentCmd.AddCommand(feeCmd)
}

func BuildCalcCommand(parentCmd *cobra.Command) {
	calcCmd := &cobra.Command{
		Use:   "calc",
		Short: "calculate fee for given PAC coin amount.",
	}
	parentCmd.AddCommand(parentCmd)

	amt := calcCmd.Flags().Int64P("amount", "a", 0, "amount in PAC.")
	rpcServer := calcCmd.Flags().StringP("rpc", "r", "", "rcp client")
	payloadType := calcCmd.Flags().StringP("payloadtype", "pt", "transfer",
		"payload type `transfer` | `bond` | `withdraw` | `unbond`")

	calcCmd.Run = func(cmd *cobra.Command, _ []string) {
		tm, err := lib.NewTxManager(lib.MainNet, *rpcServer, "", "")
		if err != nil {
			dead(cmd, err)
		}

		fee, err := tm.CalcFee(context.Background(), *amt, *payloadType)
		if err != nil {
			dead(cmd, err)
		}

		err = tm.Close()
		if err != nil {
			dead(cmd, err)
		}

		result(cmd, fmt.Sprintf("fee for %v PACs is %v", fee, amt))
	}
}
