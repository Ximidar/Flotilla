/*
* @Author: ximidar
* @Date:   2018-06-16 16:29:17
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 19:12:42
 */

package main

import (
	"log"
	"os"

	"github.com/ximidar/Flotilla/Flotilla_CLI/CLTools"
)

func main() {
	file, err := os.OpenFile("CLI.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
	CLTools.Execute()
}
