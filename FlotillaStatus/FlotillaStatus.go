/*
* @Author: Ximidar
* @Date:   2018-12-17 10:31:03
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 19:16:40
 */

package FlotillaStatus

import (
	"fmt"
	"os"

	"github.com/Ximidar/Flotilla/FlotillaStatus/NatsStatus"
	"github.com/spf13/cobra"
)

func Init(rootCmd *cobra.Command) {
	FlotStatus.AddCommand(RunStatus)
	rootCmd.AddCommand(FlotStatus)
}

var FlotStatus = &cobra.Command{
	Use:   "status",
	Short: "commands for Flotilla Status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var RunStatus = &cobra.Command{
	Use:   "start",
	Short: "commands for Flotilla Status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Status Service")
		ns, err := NatsStatus.NewNatsStatus()
		fmt.Println("Made ns")
		if err != nil {
			panic(err)
		}
		fmt.Println("Started!")
		ns.Serve()
	},
}
