/*
* @Author: Ximidar
* @Date:   2018-05-27 17:44:35
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-17 14:02:18
 */
package main

import (
	"fmt"
	"os"

	"github.com/ximidar/Flotilla/Commango/NatsConn"
)

func main() {
	gnats := NatsConn.NewNatsConn()
	gnats.Serve()

	fmt.Println("Finished")
	os.Exit(0)

}
