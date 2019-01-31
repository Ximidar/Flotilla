/*
* @Author: Ximidar
* @Date:   2019-01-13 15:38:04
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-29 16:40:58
 */

// FlotillaSystemTest is a test package to test multiple nodes together.
// theres bound to be a bunch of mixed tests in here. (It's going to be messy.)
package FlotillaSystemTest

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
	"testing"
	"time"

	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
	DS "github.com/ximidar/Flotilla/DataStructures"
	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
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
				//fmt.Print(string(buffer))
				buffer = []byte{}

				okb := []byte(ok)
				<-time.After(50 * time.Microsecond) // Pretend to process command
				serial.SendBytes(okb)

			}
		case <-time.After(100 * time.Millisecond):
			serial.SendBytes([]byte("wait\n"))
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

	// Check the status
	PrintStatus(FI.NC)

	// Select the Benchy
	Files, err := FI.GetFileStructure()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem Get file structure", err)
	Benchy := Files["root"].Contents["full_bed_print.gcode"]
	err = selectFile(Benchy, FI.NC)
	CommonTestTools.CheckErr(t, "TestFlotillaSystem select file", err)

	// Play
	err = playFile(FI.NC)
	CommonTestTools.CheckErr(t, "TestFlotillaSystem set play", err)
	<-time.After(100 * time.Millisecond)
	PrintStatus(FI.NC)
	serial.SendBytes([]byte("ok\n"))

	// Wait until the benchy has finished playing
	waitForDoneSignal(FI.NC)
	PrintStatus(FI.NC)

	// Exit Flotilla system
	<-time.After(2 * time.Second)

}

func TestFlotillaPrintingLongPrint(t *testing.T) {
	t.Skip("Skip due to taking too long. Comment out if you want a long print to occur")
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

	// Check the status
	PrintStatus(FI.NC)

	// Select the Benchy
	Files, err := FI.GetFileStructure()
	CommonTestTools.CheckErr(t, "TestFlotillaSystem Get file structure", err)
	Benchy := Files["root"].Contents["3D_Benchy.gcode"]
	err = selectFile(Benchy, FI.NC)
	CommonTestTools.CheckErr(t, "TestFlotillaSystem select file", err)

	// Play
	err = playFile(FI.NC)
	CommonTestTools.CheckErr(t, "TestFlotillaSystem set play", err)
	<-time.After(100 * time.Millisecond)
	PrintStatus(FI.NC)
	serial.SendBytes([]byte("ok\n"))

	// Wait until the benchy has finished playing
	waitForDoneSignal(FI.NC)
	PrintStatus(FI.NC)

	// Exit Flotilla system
	<-time.After(2 * time.Second)

}

func PrintStatus(nc *nats.Conn) {
	reply, err := nc.Request(PlayStructures.GetStatus, []byte(nil), nats.DefaultTimeout)
	if err != nil {
		fmt.Println(err)
		return
	}
	status, err := PlayStructures.NewStatusFlagsFromJSON(reply.Data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("IsPlaying: %v\nIsPaused: %v\nIsReady: %v\nIsError: %v\nCurrentAction: %v\n",
		status.IsPlaying,
		status.IsPaused,
		status.IsReady,
		status.IsError,
		status.CurrentAction)
}

func selectFile(file *Files.File, nc *nats.Conn) error {
	selectAction, err := FS.NewFileAction(FS.SelectFile, file.Path)
	if err != nil {
		return err
	}

	reply, err := selectAction.SendAction(nc, 5*time.Second)
	if err != nil {
		return err
	}

	rJSON, err := returnReplyJSON(reply)
	if err != nil {
		return err
	}
	if rJSON.Success {
		return nil
	}
	return errors.New(string(rJSON.Message))
}

func playFile(nc *nats.Conn) error {
	streamAction, err := PlayStructures.NewPlayAction(PlayStructures.PLAY)
	if err != nil {
		return err
	}

	err = streamAction.Send(nc)
	if err != nil {
		return err
	}
	return nil

}

func returnReplyJSON(msg *nats.Msg) (*DS.ReplyJSON, error) {
	msgdata := DS.ReplyJSON{}

	// unmarshal msg data
	err := json.Unmarshal(msg.Data, &msgdata)
	if err != nil {
		return nil, err
	}

	return &msgdata, nil
}

func returnReplyString(msg *nats.Msg) (*DS.ReplyString, error) {
	msgdata := DS.ReplyString{}

	// unmarshal msg data
	err := json.Unmarshal(msg.Data, &msgdata)
	if err != nil {
		return nil, err
	}

	return &msgdata, nil
}

func waitForDoneSignal(NC *nats.Conn) {
	flags, err := PlayStructures.NewStatusObserver()
	if err != nil {
		fmt.Println("Error in testing!", err)
	}

	NC.Subscribe(PlayStructures.PublishStatus, flags.UpdateStatusFromNats)

	blockChan := make(chan bool, 10)

	doneFunc := func() {
		blockChan <- true
	}

	flags.AddFunction(PlayStructures.DONE, doneFunc)

	fmt.Println("Waiting for done")
	<-blockChan
	fmt.Println("Done has been set")
}
