/*
* @Author: Ximidar
* @Date:   2018-10-01 18:58:24
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 18:26:07
 */
package FFM

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	NI "github.com/Ximidar/Flotilla/Flotilla_File_Manager/NatsFile"
	"github.com/spf13/cobra"
)

func Init(rootCmd *cobra.Command) {
	ffmCLI.AddCommand(startFFM)
	rootCmd.AddCommand(ffmCLI)
}

var ffmCLI = &cobra.Command{
	Use:     "file-manager",
	Aliases: []string{"ffm"},
	Short:   "commands for flotilla file manager",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var startFFM = &cobra.Command{
	Use:   "start",
	Short: "start the flotilla file manager",
	Run: func(cmd *cobra.Command, args []string) {
		RunFFM()
	},
}

// TermChannel will monitor for an exit signal
var TermChannel chan os.Signal

func RunFFM() {
	fmt.Println("Creating File Manager")
	NatsIO, err := NI.NewNatsFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(NatsIO.FileManager.RootFolderPath)
	Run()
}

// Run will keep the program alive
func Run() {
	// Function for waiting for exit on the main loop
	// Wait for termination
	TermChannel = make(chan os.Signal)
	signal.Notify(TermChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Flotilla File Manager Started")
	<-TermChannel
	fmt.Println("Recieved Interrupt Sig, Now Exiting.")
	os.Exit(0)
}
