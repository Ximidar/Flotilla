/*
* @Author: Ximidar
* @Date:   2018-12-12 14:49:39
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-25 13:56:45
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	FakeSerialDevice "github.com/Ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
)

func main() {
	serial := FakeSerialDevice.NewFakeSerial()
	go serial.ReadMaster()
	//go serial.SendMaster()

	go watchSignals(serial)
	run(serial)

}

func watchSignals(serial *FakeSerialDevice.FakeSerial) {
	// shutdown signal
	quit := make(chan os.Signal, 1)
	all_sig := make(chan os.Signal, 100)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	signal.Notify(all_sig)

	for {
		select {
		case sig := <-quit:
			fmt.Println("Got Quit Signal:", sig)
			serial.Close()
			os.Exit(0)
		case sig := <-all_sig:
			fmt.Println("Got Signal:", sig)
			os.Exit(0)
		}
	}
}

func run(serial *FakeSerialDevice.FakeSerial) {
	var buffer []byte
	ok := "ok\n"

	for {
		select {
		case buf := <-serial.ReceiveStream:
			buffer = append(buffer, buf)
			if buf == 10 { // if we detect a newline
				// fmt.Print(string(buffer))
				//line := string(buffer)
				buffer = []byte{}
				okb := []byte(ok)
				<-time.After(10 * time.Millisecond) // Pretend to process command
				serial.SendBytes(okb)
			}
		case <-serial.SlaveOpenChan:
			buffer = []byte{}
			startb := []byte("start\n")
			serial.SendBytes(startb)
		case <-time.After(10 * time.Second):
			waitb := []byte("wait\n")
			serial.SendBytes(waitb)

		}

	}
}
