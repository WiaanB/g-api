package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:  "serve",
	Long: "This will serve the Gotcha application server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Serving the Gotcha application...")
		fmt.Println(args)
	},
}
