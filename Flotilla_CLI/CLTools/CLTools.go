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

	"github.com/Ximidar/Flotilla/Flotilla_CLI/ui/ContentBox"

	"github.com/Ximidar/Flotilla/Flotilla_CLI/Helm"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/UserInterface"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/ui"
	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(tcellInterface)
	rootCmd.AddCommand(helm)
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

var tcellInterface = &cobra.Command{
	Use:   "tui",
	Short: "Show the cli UI for Flotilla",
	Long:  `This will open the cli UI for Flotilla. This has tools for monitoring the command line and starting prints (or it will in the future)`,
	Run: func(cmd *cobra.Command, args []string) {
		tgui, err := ui.NewMainScreen()
		if err != nil {
			panic(err)
		}
		tgui.AddQuitKey("q")
		ContentBox.NewContentBox(tgui.Screen, "MainBox!", 10, 10, 20, 20)

		tgui.Run()
	},
}

var helm = &cobra.Command{
	Use:   "helm",
	Short: "start a flotilla instance",
	Run: func(cmd *cobra.Command, args []string) {
		Helm.StartHelm("")
	},
}
