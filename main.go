package main

import (
	"fmt"
	"os"

	commango "github.com/Ximidar/Flotilla/Commango"
	version "github.com/Ximidar/Flotilla/CommonTools/versioning"
	"github.com/Ximidar/Flotilla/FlotillaStatus"
	"github.com/Ximidar/Flotilla/FlotillaWebAPI"
	CLTools "github.com/Ximidar/Flotilla/Flotilla_CLI"
	FFM "github.com/Ximidar/Flotilla/Flotilla_File_Manager"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flot",
	Short: "Flotilla_CLI is the cli tool for the Flotilla system",
	Long:  `Use this tool for starting up any module of Flotilla`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		} else {
			fmt.Println("Flotilla CLI")
			fmt.Println("Written By: Matt Pedler")
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Version Details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:       %s\n", version.Version)
		fmt.Printf("Compiled By:   %s\n", version.CompiledBy)
		fmt.Printf("Compiled Date: %s\n", version.CompiledDate)
		fmt.Printf("Commit Hash:   %s\n", version.CommitHash)
	},
}

func init() {
	commango.Init(rootCmd)
	CLTools.Init(rootCmd)
	FFM.Init(rootCmd)
	FlotillaStatus.Init(rootCmd)
	FlotillaWebAPI.Init(rootCmd)

	rootCmd.AddCommand(versionCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
