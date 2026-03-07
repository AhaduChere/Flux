package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "flux",
	// what shows up when u do flux --help
	Short: "flux expects an argument",
}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}

}
