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

	"github.com/spf13/cobra"
	"github.com/Ximidar/Flotilla/NodeLauncher/FlotillaInstance"
	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
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

var startFlotillaInstance = &cobra.Command{
	Use:   "Start",
	Short: "Start the Flotilla Instance at a specified package location",
	Long:  `Start the Flotilla Instance at a specific package location.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Flotilla Instance at:", pathFlag)
		flotilla, err := FlotillaInstance.NewFlotillaInstance(pathFlag, natsFlag, tlsFlag)
		if err != nil {
			fmt.Println("Could not start Flotilla instance due to:", err)
			return
		}
		err = flotilla.Serve()
		if err != nil {
			fmt.Println(err)
		}

	},
}

var createFlotillaRootFolder = &cobra.Command{
	Use:   "CreateRoot",
	Short: "Generate the files for a Flotilla Package",
	Long:  `This will take a directory and create a Flotilla Package inside of it`,
	Run: func(cmd *cobra.Command, args []string) {

		if pathFlag == "" {
			cmd.Help()
			os.Exit(1)
		}
		if localPopulateFlag == false {
			fmt.Println("Populating from internet does not exist yet. Sorry")
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

		fmt.Println("Populating Root Folder")
		pf, err := RootFolder.NewPopulateFolder(rf, archFlag)
		if err != nil {
			fmt.Println("Could not populate due to: ", err)
			os.Exit(1)
		}
		err = pf.Populate(skipMakeAllFlag)
		if err != nil {
			fmt.Println("Could not populate due to: ", err)
			os.Exit(1)
		}
	},
}

var PrintVersion = &cobra.Command{
	Use:   "version",
	Short: "print the logo and version",
	Long:  `This command will show the current version and logo of the program`,
	Run: func(cmd *cobra.Command, args []string) {
		fi := new(FlotillaInstance.FlotillaInstance)
		fi.Logo()
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
	createFlotillaRootFolder.Flags().BoolVarP(&localPopulateFlag, "populateLocal", "l", false, "Set to true to populate from internet. Otherwise it will populate from local GOPATH")
	createFlotillaRootFolder.Flags().BoolVarP(&skipMakeAllFlag, "skip-make-all", "m", false, "Set to true to skip making all binaries")
	createFlotillaRootFolder.Flags().StringVarP(&pathFlag, "path", "p", "", "Directory to create Root Folder in")
	createFlotillaRootFolder.Flags().StringVarP(&archFlag, "arch", "a", "amd64", "Architechture to populate with. options are: amd64, arm, arm64")
	rootCmd.AddCommand(createFlotillaRootFolder)

	// Command to start Flotilla Instance
	startFlotillaInstance.Flags().BoolVarP(&natsFlag, "nats", "n", true, "Set to false to not look for and start a Nats server")
	startFlotillaInstance.Flags().BoolVar(&tlsFlag, "tls", false, "Start Flotilla Instance with TLS configuration enabled")
	startFlotillaInstance.Flags().StringVarP(&pathFlag, "path", "p", "", "Directory to Flotilla Root Folder")
	rootCmd.AddCommand(startFlotillaInstance)

	// Command to show version and logo
	rootCmd.AddCommand(PrintVersion)
}
