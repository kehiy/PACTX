package main

import (
	"github.com/kehiy/PACTX/cli/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "pactx",
		Version: "0.1.0-beta",
	}

	commands.BuildFeeCommands(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
