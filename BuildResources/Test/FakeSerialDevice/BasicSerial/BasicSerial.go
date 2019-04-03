/*
* @Author: Ximidar
* @Date:   2019-04-02 15:49:42
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 15:56:41
 */

package BasicSerial

import (
	"fmt"

	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
)

// MakeBasicSerial will set up an echo serial device
func Make() *FakeSerialDevice.FakeSerial {
	serial := FakeSerialDevice.NewFakeSerial()
	go serial.ReadMaster()
	go runEcho(serial)

	return serial

}

// runEcho will receive all bytes on the receive stream and echo back each line it is given.
func runEcho(serial *FakeSerialDevice.FakeSerial) {
	var buffer []byte
	for {
		select {
		case buf := <-serial.ReceiveStream:
			buffer = append(buffer, buf)
			if buf == 10 { // if we detect a newline
				fmt.Print(string(buffer))
				serial.SendBytes(buffer)
				buffer = []byte{}

			}
		}
	}
}
