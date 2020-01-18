/*
* @Author: Ximidar
* @Date:   2018-10-01 18:58:24
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 18:26:07
 */
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	NI "github.com/Ximidar/Flotilla/Flotilla_File_Manager/NatsFile"
)

// TermChannel will monitor for an exit signal
var TermChannel chan os.Signal

func main() {
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
	os.Exit(1)
}
