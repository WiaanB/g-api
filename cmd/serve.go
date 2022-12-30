package cmd

import (
	"encoding/json"
	"fmt"
	"gotcha/util"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     ConfigurationFile
)

func init() {
	serveCmd.Flags().StringVarP(&cfgFile, "cfgFile", "c", "gotcha-config.json", "The config file you want to use for your gotcha instance")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:  "serve",
	Long: "This will serve the Gotcha application server.",
	Run: func(cmd *cobra.Command, args []string) {
		util.FatalErrorWrapper(getConfigFile(), "Failed to read configure file")
		server()
	},
}

func server() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomeHandler)

	fmt.Printf("serving up your hot plate of gotcha on port :%d\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "welcome to gotcha :))",
	})
}

func getConfigFile() error {
	formattedName := cfgFile
	if !strings.HasSuffix(formattedName, ".json") {
		formattedName = formattedName + ".json"
	}

	finfo, _ := os.Stat(fmt.Sprintf("configs/%s", formattedName))
	if finfo == nil {
		return fmt.Errorf("unfound find file, '%s', please ensure it exists", formattedName)
	}

	file, err := os.Open(fmt.Sprintf("configs/%s", formattedName))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	return nil
}
