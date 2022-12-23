package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "gotcha",
	Long: "Gotcha is an application to track who is the biggest idiot.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd", cmd)
		fmt.Println("args", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
