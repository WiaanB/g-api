package cmd

import (
	"encoding/json"
	"fmt"
	"gotcha/util"
	"os"

	"github.com/spf13/cobra"
)

var (
	port       int
	mode       string
	configFile string
)

type configurationFile struct {
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

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
		util.FatalErrorWrapper(setupConfigFolder(), "Failed to set up configuration folder")
	},
}

func setupConfigFolder() (err error) {
	_, err = os.Stat("configs")
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir("configs", 0755)
		if err != nil && !os.IsExist(err) {
			return
		}
	} else if err != nil {
		return
	}
	err = createConfigFile()
	return
}

func createConfigFile() error {
	cfg := configurationFile{
		Port: port,
		Mode: mode,
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	util.ErrorWrapper(err, fmt.Sprintf("Failed to marshal the config file '%s'\n", configFile))
	err = os.WriteFile(fmt.Sprintf("configs/%s", configFile), data, 0755)
	util.FatalErrorWrapper(err, fmt.Sprintf("Failed to create the config file '%s'\n", configFile))
	return nil
}
