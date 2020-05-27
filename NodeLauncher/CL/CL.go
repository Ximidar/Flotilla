/*
* @Author: Ximidar
* @Date:   2019-02-05 15:23:23
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 17:39:51
 */

package CL

import (
	"fmt"
	"os"

	"github.com/Ximidar/Flotilla/NodeLauncher/GetNats"
	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
	"github.com/Ximidar/Flotilla/NodeLauncher/snappy"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "NodeLauncher [COMMAND] [FLAG(s)] ",
	Short: "NodeLauncher will launch a Flotilla instance",
	Long:  `Use this tool to launch a Flotilla instance`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		} else {
			fmt.Println("Node Launcher")
			fmt.Println("Written By: Matt Pedler")
		}
	},
}

var pathFlag string
var localPopulateFlag bool
var skipMakeAllFlag bool

var archFlag string
var natsFlag bool
var tlsFlag bool

// BuildFlotilla will build packages of Flotilla for snapping
var BuildFlotilla = &cobra.Command{
	Use:   "BuildFlotilla",
	Short: "Generate the files for a Flotilla Package",
	Long:  `This command will build packages for all arches or a specific arch`,
	Run: func(cmd *cobra.Command, args []string) {

		if pathFlag == "" {
			cmd.Help()
			os.Exit(1)
		}

		if archFlag == "" {
			fmt.Println("You must Specify an arch")
			os.Exit(1)
		}

		// Attempt to build Root Folder
		rf, err := RootFolder.GenerateRootFolder(pathFlag)
		if err != nil {
			fmt.Println("Could not Generate Root Folder because:", err)
			os.Exit(1)
		}

		fmt.Println("Flotilla Root Folder created at", rf.RootPath)

		// Make Paths
		fmt.Println("Populating Root Folder")
		pf, err := RootFolder.NewPopulateFolder(rf, archFlag)
		if err != nil {
			fmt.Println("Could not populate due to: ", err)
			os.Exit(1)
		}

		// Build Flotilla
		err = pf.Populate(skipMakeAllFlag)
		if err != nil {
			fmt.Println("Could not populate due to: ", err)
			os.Exit(1)
		}

		// Add Nats Server
		err = GetNats.DownloadNats(rf)
		if err != nil {
			fmt.Println("Could not download nats due to: ", err)
			os.Exit(1)
		}
		err = GetNats.UnzipAndClean(rf)
		if err != nil {
			fmt.Println("Could not unzip nats due to: ", err)
			os.Exit(1)
		}

		// package arches into seperate zip files
		err = rf.PackageArches()
		if err != nil {
			fmt.Println("Could not Package Flotilla due to: ", err)
			os.Exit(1)
		}

		// create snaps
		snappy := snappy.NewSnappy(rf)
		err = snappy.MakeSnaps()
		if err != nil {
			fmt.Println("Could not make snaps: ", err)
			panic(err)
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
	// Command to create Root Folder
	BuildFlotilla.Flags().BoolVarP(&localPopulateFlag, "populateLocal", "l", false, "Set to true to populate from internet. Otherwise it will populate from local GOPATH")
	BuildFlotilla.Flags().BoolVarP(&skipMakeAllFlag, "skip-make-all", "m", false, "Set to true to skip making all binaries")
	BuildFlotilla.Flags().StringVarP(&pathFlag, "path", "p", "", "Directory to create Root Folder in")
	BuildFlotilla.Flags().StringVarP(&archFlag, "arch", "a", "amd64", "Architechture to populate with. options are: amd64, arm, arm64")
	rootCmd.AddCommand(BuildFlotilla)
}
