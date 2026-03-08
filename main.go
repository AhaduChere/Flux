package main

import (
	"encoding/json"
	"flux/internal/config"
	"flux/internal/output"
	"flux/internal/request"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// NOTE: Entry Point
var Program = &cobra.Command{
	Use:   "flux",
	Short: "flux expects an argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//NOTE: Reads Fluxfile
		cfg, err := config.Parsefile("Fluxfile")
		if err != nil {
			fmt.Println("Error reading Fluxfile")
			return
		}
		// NOTE: Looks up which endpoint user wants to use
		endpoint, ok := cfg.Endpoints[args[0]]
		if !ok {
			fmt.Println("Endpoint not found")
			return
		}
		// NOTE: Build the request
		var bodyStr *string
		if len(endpoint.Body) > 0 {
			bodyBytes, err := json.Marshal(endpoint.Body)
			if err != nil {
				fmt.Println("Failed to encode body")
				return
			}
			s := string(bodyBytes)
			bodyStr = &s
		}

		req := request.Request{
			URL:    cfg.URL + endpoint.Path,
			Method: endpoint.Method,
			Body:   bodyStr,
		}
		// NOTE: Send the request
		rep, err := request.SendRequest(req)
		if err != nil {
			fmt.Println("Request failed:", err)
			return
		}
		// NOTE: print the response
		output.Print(rep)
	},
}

// NOTE: Main function
func main() {
	err := Program.Execute()
	if err != nil {
		os.Exit(1)
	}
}
