package backend

import (
	"fmt"
	"os"

	"github.com/Ximidar/Flotilla/FlotillaWebAPI/backend/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "FlotillaWeb",
	Short: "FlotillaWeb will start the API server",
	Long:  `Use this tool to launch a Flotilla website`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		} else {
			fmt.Println("Flotilla Web")
			fmt.Println("Written By: Matt Pedler")
		}
	},
}

var WorkingPath string

// Port for the API to use
var Port int

var serveAPI = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Long:  `Serve the API`,
	Run: func(cmd *cobra.Command, args []string) {

		api.Serve(Port, WorkingPath)

	},
}

// Execute will be run as the root command
func Execute(workingPath string) {

	WorkingPath = workingPath
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Serve API command
	serveAPI.Flags().IntVarP(&Port, "port", "p", 3000, "Port to serve API")
	rootCmd.AddCommand(serveAPI)
}
