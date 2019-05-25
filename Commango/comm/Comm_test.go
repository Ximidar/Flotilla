/*
* @Author: Ximidar
* @Date:   2019-04-02 15:35:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 16:13:48
 */

package commango

import (
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/BasicSerial"
	CS "github.com/ximidar/Flotilla/DataStructures/CommStructures"
)

type TestAdaptor struct {
}

func (*TestAdaptor) ReadLineEmitter(line string) error         { return nil }
func (*TestAdaptor) WriteLineEmitter(line string) error        { return nil }
func (*TestAdaptor) PublishStatus(status *CS.CommStatus) error { return nil }

// TestCommConstructor will test if a Comm Object can be made
func TestCommConstructor(t *testing.T) {
	comm := NewComm(&TestAdaptor{})

	if comm == nil {
		t.Fatal("Comm could not be made")
	}
}

// TestCommBasicFunction will test if the comm can run basic functionality like
// connect, disconnect, reconnect, read, write
func TestCommBasicFunction(t *testing.T) {
	// setup fake serial device
	fsd := BasicSerial.Make()
	defer fsd.Close()

	// setup Comm
	comm := NewComm(&TestAdaptor{})
	defer comm.CloseComm()

	// Connect Comm to fsd
	err := comm.InitComm(&CS.InitComm{
		Port: fsd.Address,
		Baud: 115200,
	})
	CommonTestTools.CheckErr(t, "Attempt to Init Comm", err)

	// Open Port
	err = comm.OpenComm()
	CommonTestTools.CheckErr(t, "Attempt to open comm", err)

	// Write Hello
	expected := "Hello\n"
	lenWritten, err := comm.WriteComm(expected)
	CommonTestTools.CheckErr(t, "Attempting to write comm", err)
	CommonTestTools.CheckEquals(t, len(expected), lenWritten)

	// Read Hello
	line := <-comm.ReadStream
	CommonTestTools.CheckEquals(t, expected, line)

}
