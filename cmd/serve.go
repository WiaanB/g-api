package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

func init() {
	serveCmd.Flags().StringVarP(&cfgFile, "cfgFile", "c", "gotcha-config.json", "The config file you want to use for your gotcha instance")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:  "serve",
	Long: "This will serve the Gotcha application server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Serving the Gotcha application...")
		server()
	},
}

func server() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "welcome to gotcha :))",
	})
}
