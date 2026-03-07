package cmd

import (
	"flux/internal/config"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:  "flux <endpoint>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Parsefile("Fluxfile")
		if err != nil ||  {
			//NOTE: something went wrong reading the Fluxfile
		}
		endpoint, ok := cfg.Endpoints[args[0]]
		if !ok {
			//NOTE: endpoint not found
		}



	},
}
