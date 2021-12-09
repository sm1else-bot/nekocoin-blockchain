package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var nkcCmd = &cobra.Command{
		Use:   "nkc",
		Short: "NekoCoin Command Line Interface",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	nkcCmd.AddCommand(versionCmd)
	nkcCmd.AddCommand(balancesCmd())
	nkcCmd.AddCommand(txCmd())

	err := nkcCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
