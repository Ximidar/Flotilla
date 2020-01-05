/*
* @Author: Ximidar
* @Date:   2018-12-12 14:49:39
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-25 13:56:45
 */

package main

import (
	"fmt"
	"strings"
	"time"

	FakeSerialDevice "github.com/Ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
)

func main() {
	serial := FakeSerialDevice.NewFakeSerial()
	go serial.ReadMaster()
	//go serial.SendMaster()

	run(serial)

}

func run(serial *FakeSerialDevice.FakeSerial) {
	var buffer []byte
	ok := "ok\n"
	for {
		select {
		case buf := <-serial.ReceiveStream:
			buffer = append(buffer, buf)
			if buf == 10 { // if we detect a newline
				fmt.Print(string(buffer))
				line := string(buffer)
				buffer = []byte{}

				if strings.Contains(line, "start") {
					startb := []byte("start\n")
					<-time.After(10 * time.Millisecond) // Pretend to process command
					serial.SendBytes(startb)
				} else {
					okb := []byte(ok)
					<-time.After(10 * time.Millisecond) // Pretend to process command
					serial.SendBytes(okb)
				}

			}
		}
	}
}
