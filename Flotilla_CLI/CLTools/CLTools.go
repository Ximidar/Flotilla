/*
* @Author: Ximidar
* @Date:   2018-06-16 16:53:05
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-11-30 14:08:27
 */

package CLTools

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ximidar/Flotilla/Flotilla_CLI/UserInterface"
)

var rootCmd = &cobra.Command{
	Use:   "Flotilla_CLI",
	Short: "Flotilla_CLI is the cli tool for the Flotilla system",
	Long: `Use this tool to control Flotilla from the command line
  		 This tool will help you print files, check the status of
  		 a print, or help you control and monitor the printer command line`,
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

// Execute will be run as the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(printerface)
}

var printerface = &cobra.Command{
	Use:   "ui",
	Short: "Show the cli UI for Flotilla",
	Long:  `This will open the cli UI for Flotilla. This has tools for monitoring the command line and starting prints (or it will in the future)`,
	Run: func(cmd *cobra.Command, args []string) {
		cligui, err := UserInterface.NewCliGui()
		if err != nil {
			panic(err)
		}
		cligui.ScreenInit()
	},
}
