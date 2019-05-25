/*
* @Author: Ximidar
* @Date:   2018-12-17 10:31:03
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 19:16:40
 */

package main

import (
	"fmt"

	"github.com/ximidar/Flotilla/FlotillaStatus/NatsStatus"
)

func main() {
	fmt.Println("Attempting to start")
	ns, err := NatsStatus.NewNatsStatus()
	fmt.Println("Made ns")
	if err != nil {
		panic(err)
	}
	fmt.Println("Started!")
	ns.Serve()
}
