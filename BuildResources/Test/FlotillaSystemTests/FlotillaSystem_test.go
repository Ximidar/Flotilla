/*
* @Author: Ximidar
* @Date:   2019-01-13 15:38:04
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 21:46:20
 */

// FlotillaSystemTest is a test package to test multiple nodes together.
// theres bound to be a bunch of mixed tests in here. (It's going to be messy.)
package FlotillaSystemTest

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
)

// StartTestFlotilla will start a Flotilla instance with the most recently compiled binaries
// Make sure you actually compile Flotilla before running, This also does not include
// A NATS instance yet. So you will have to supply that for now
func StartTestFlotilla() (chan bool, error) {

	fileManagerProc, err := CreateProcess("Flotilla_File_Manager", "FlotillaFileManager")
	if err != nil {
		return nil, err
	}

	flotillaStatusProc, err := CreateProcess("FlotillaStatus", "FlotillaStatus")
	if err != nil {
		return nil, err
	}

	commango, err := CreateProcess("Commango", "Commango")

	if err != nil {
		return nil, err
	}

	exitChan := make(chan bool, 10)

	killEverything := func() {
		fmt.Println("Killing Everything")
		killerr := syscall.Kill(commango.Process.Pid, syscall.SIGKILL)
		if killerr != nil {
			fmt.Println(killerr)
		}
		killerr = syscall.Kill(flotillaStatusProc.Process.Pid, syscall.SIGKILL)
		if killerr != nil {
			fmt.Println(killerr)
		}
		killerr = syscall.Kill(fileManagerProc.Process.Pid, syscall.SIGKILL)
		if killerr != nil {
			fmt.Println(killerr)
		}
		fmt.Println("Everything has been killed")
		exitChan <- true
	}

	StartProcs := func() {
		commango.Start()
		flotillaStatusProc.Start()
		fileManagerProc.Start()
		select {
		case <-exitChan:
			killEverything()
			return
		}
	}
	go StartProcs()

	return exitChan, nil

}

func CreateProcess(FolderName string, ExeName string) (*exec.Cmd, error) {
	ProgBinFolder, err := CommonTestTools.FindFlotillaBinFolder(FolderName)
	if err != nil {
		return nil, err
	}

	Executable := path.Clean(ProgBinFolder + "/" + ExeName)
	cmd := CommonTestTools.MakeProcess(Executable, "")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil

}

func run(serial *FakeSerialDevice.FakeSerial, exitChan chan bool) {
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
		case <-exitChan:
			return
		}
	}
}

func TestFlotillaPrinting(t *testing.T) {
	FI, err := FlotillaInterface.NewFlotillaInterface()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem", err)
	exitChan, err := StartTestFlotilla()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem", err)

	// In case we need to ctrl-c out this will always run
	sendExitSig := func() {
		fmt.Println("Sending exit sig")
		exitChan <- true
		<-time.After(2 * time.Second)
	}
	defer sendExitSig()

	// Allow time for server startup
	<-time.After(100 * time.Millisecond)

	// Create a Fake Serial Device
	serial := FakeSerialDevice.NewFakeSerial()
	go serial.ReadMaster()
	go run(serial, exitChan)

	// Connect to the Fake Serial Device
	err = FI.CommSetConnectionOptions(FakeSerialDevice.SerialName, 115200)
	CommonTestTools.CheckErr(t, "TestFlotillaSystem", err)
	err = FI.CommConnect()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem", err)
	// send Start
	serial.SendBytes([]byte("start\n"))
	<-time.After(100 * time.Millisecond)

	// Select the Benchy
	Files, err := FI.GetFileStructure()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem", err)
	Benchy := Files["root"].Contents["3D_Benchy.gcode"]
	FI.SelectAndPlayFile(Benchy)
	serial.SendBytes([]byte("ok\n"))

	// Wait until the benchy has finished playing

	// Exit Flotilla system
	<-time.After(5 * time.Second)

}
