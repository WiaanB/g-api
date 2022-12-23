package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	port       int
	mode       string
	configFile string
)

func init() {
	configureCmd.Flags().StringVarP(&configFile, "configFile", "c", "gotcha-config.json", "A name you would like to give the config file, multiple may be configured")
	configureCmd.Flags().StringVarP(&mode, "mode", "m", "dev", "The mode you would like your gotcha instance to run in")
	configureCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port that the gotcha instance will run on")
	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:  "configure",
	Long: "Setting up the base config file required for the Gotcha application",
	Run: func(cmd *cobra.Command, args []string) {
		setupConfigFolder()
	},
}

func setupConfigFolder() {
	exist, notExist := os.Stat("configs")
	if notExist != nil {
		if os.IsNotExist(notExist) {
			fmt.Println("Does not exist")
		}
	}
	fmt.Println(exist)
}
