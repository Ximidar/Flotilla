/*
* @Author: ximidar
* @Date:   2018-06-16 16:29:17
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 19:12:42
 */

package CLTools

import (
	"github.com/Ximidar/Flotilla/Flotilla_CLI/Helm"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/ui/ContentBox"

	"github.com/Ximidar/Flotilla/Flotilla_CLI/UserInterface"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/ui"
	"github.com/spf13/cobra"
)

func Init(rootCmd *cobra.Command) {

	UITools.AddCommand(Printerface)
	UITools.AddCommand(TcellInterface)
	UITools.AddCommand(HelmCLI)
	rootCmd.AddCommand(UITools)

}

var UITools = &cobra.Command{
	Use:   "ui",
	Short: "commands for user interface flotilla",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var Printerface = &cobra.Command{
	Use:   "gui",
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

var TcellInterface = &cobra.Command{
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

var HelmCLI = &cobra.Command{
	Use:   "helm",
	Short: "start a flotilla instance",
	Run: func(cmd *cobra.Command, args []string) {
		Helm.StartHelm("")
	},
}
