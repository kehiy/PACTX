package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func dead(cmd *cobra.Command, err error) {
	cmd.PrintErr(err)
	os.Exit(1)
}

func result(cmd *cobra.Command, result string) {
	cmd.Println(result)
	os.Exit(0)
}
