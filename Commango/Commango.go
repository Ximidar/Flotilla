/*
* @Author: Ximidar
* @Date:   2018-05-27 17:44:35
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-17 14:02:18
 */
package commango

import (
	"fmt"
	"os"

	"github.com/Ximidar/Flotilla/Commango/NatsConn"
	"github.com/spf13/cobra"
)

var CommangoStart = &cobra.Command{
	Use:   "start",
	Short: "start commango",
	Long:  `start the serial communication layer of flotilla`,
	Run: func(cmd *cobra.Command, args []string) {
		gnats := NatsConn.NewNatsConn()
		gnats.Serve()

		fmt.Println("Finished")
		os.Exit(0)
	},
}

var Comm = &cobra.Command{
	Use:   "comm",
	Short: "commango Functions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func Init(rootCmd *cobra.Command) {
	Comm.AddCommand(CommangoStart)
	rootCmd.AddCommand(Comm)
}
