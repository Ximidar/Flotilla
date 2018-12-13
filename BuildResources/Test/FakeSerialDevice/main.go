/*
* @Author: Ximidar
* @Date:   2018-12-12 14:49:39
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-12 16:40:22
 */

package main

import (
	"fmt"
	"time"

	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
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
				buffer = []byte{}

				okb := []byte(ok)
				<-time.After(10 * time.Millisecond) // Pretend to process command
				serial.SendBytes(okb)
			}
		}
	}
}
