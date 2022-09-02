package FlotillaWebAPI

import (
	"log"
	"os"

	"github.com/Ximidar/Flotilla/FlotillaWebAPI/backend"
	"github.com/spf13/cobra"
)

func Init(rootCmd *cobra.Command) {
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

var RunAPI = &cobra.Command{
	Use:   "start",
	Short: "commands for Flotilla Status",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		backend.Execute(dir)
	},
}
