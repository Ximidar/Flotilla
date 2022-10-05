package FlotillaWebAPI

import (
	"log"
	"os"

	"github.com/Ximidar/Flotilla/FlotillaWebAPI/backend/api"
	"github.com/spf13/cobra"
)

func Init(rootCmd *cobra.Command) {
	RunAPI.Flags().IntVarP(&Port, "port", "p", 3000, "Port to serve API")
	FlotWebApi.AddCommand(RunAPI)
	rootCmd.AddCommand(FlotWebApi)
}

var FlotWebApi = &cobra.Command{
	Use:   "api",
	Short: "commands for Flotilla Web API",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var WorkingPath string

// Port for the API to use
var Port int
var RunAPI = &cobra.Command{
	Use:   "start",
	Short: "commands for Flotilla Status",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		api.Serve(Port, dir)
	},
}
