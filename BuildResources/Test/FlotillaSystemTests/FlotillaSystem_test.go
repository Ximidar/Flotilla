/*
* @Author: Ximidar
* @Date:   2019-01-13 15:38:04
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 23:30:25
 */

// FlotillaSystemTest is a test package to test multiple nodes together.
// theres bound to be a bunch of mixed tests in here. (It's going to be messy.)
package FlotillaSystemTest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"syscall"
	"testing"
	"time"

	"github.com/Ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/Ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/Printer"
	FakeSerialDevice "github.com/Ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
	DS "github.com/Ximidar/Flotilla/DataStructures"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/Ximidar/Flotilla/Flotilla_File_Manager/FileManager"
	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/gnatsd/test"
	"github.com/nats-io/nats.go"
)

var TestLocation = "/tmp/FlotillaSystem/Test/Flotilla"

// StartTestFlotilla will start a Flotilla instance with the most recently compiled binaries
// Make sure you actually compile Flotilla before running, This also does not include
// A NATS instance yet. So you will have to supply that for now
func StartTestFlotilla() (chan bool, error) {
	os.RemoveAll(TestLocation)

	// Figure out if we need to start a nats server
	var nserver *server.Server
	_, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		nserver = test.RunDefaultServer()
		// wait for server startup to happen
		<-time.After(200 * time.Millisecond)
	} else {
		fmt.Println("Server already exists on the default port. No need to create another one")
	}

	NodeLauncher, err := CreateProcess("NodeLauncher", "NodeLauncher",
		"CreateRoot", "-p", TestLocation, "-a=amd64", "-l=true")
	if err != nil {
		return nil, err
	}

	err = NodeLauncher.Run()
	NodeLauncher.Wait()
	if err != nil {
		return nil, err
	}
	err = setupTestFiles()
	if err != nil {
		return nil, err
	}

	NodeLauncher, err = CreateProcess("NodeLauncher", "NodeLauncher",
		"Start", "-p", TestLocation, "-n=false")
	if err != nil {
		return nil, err
	}
	exitChan := make(chan bool, 10)

	killEverything := func() {
		fmt.Println("Killing Everything")

		killerr := syscall.Kill(NodeLauncher.Process.Pid, syscall.SIGINT)
		if killerr != nil {
			fmt.Println(killerr)
		}
		fmt.Println("Everything has been killed")

	}

	StartProc := func() {
		NodeLauncher.Start()
		select {
		case <-exitChan:
			killEverything()
			NodeLauncher.Wait()
			os.RemoveAll(TestLocation)
			nserver.Shutdown()
			exitChan <- true
		}

	}
	go StartProc()

	// Wait a little bit for the server to start up before continuing the test
	<-time.After(2 * time.Second)

	return exitChan, nil

}

func CreateProcess(FolderName string, ExeName string, args ...string) (*exec.Cmd, error) {
	ProgBinFolder, err := CommonTestTools.FindFlotillaBinFolder(FolderName)
	if err != nil {
		return nil, err
	}

	Executable := path.Clean(ProgBinFolder + "/" + ExeName)
	cmd := CommonTestTools.MakeProcess(Executable, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil

}

func TestFlotillaPrinting(t *testing.T) {
	exitChan, err := StartTestFlotilla()
	CommonTestTools.CheckErr(t, "Could not start Flotilla test", err)
	FI, err := FlotillaInterface.NewFlotillaInterface()
	CommonTestTools.CheckErr(t, "Could not create new flotilla interface", err)

	// In case we need to ctrl-c out this will always run
	sendExitSig := func() {
		fmt.Println("Sending exit sig")
		exitChan <- true
		select {
		case <-exitChan:
			fmt.Println("Exiting Complete")
		case <-time.After(5 * time.Second):
			fmt.Println("Exiting After 5 Seconds")
		}

	}
	defer sendExitSig()

	// Allow time for server startup
	<-time.After(100 * time.Millisecond)

	// Create a Fake Serial Device
	os.RemoveAll("/tmp/fakeprinter")
	serial := FakeSerialDevice.NewFakeSerial()
	printer, _ := Printer.NewPrinter()
	go serial.ReadMaster()
	go printer.Run(serial, exitChan)

	// Connect to the Fake Serial Device
	err = FI.CommSetConnectionOptions(FakeSerialDevice.SerialName, 115200)
	CommonTestTools.CheckErr(t, "Could not set connection options", err)
	err = FI.CommConnect()
	CommonTestTools.CheckErr(t, "Could not connect", err)
	// send Start
	serial.SendBytes([]byte("start\n"))
	<-time.After(100 * time.Millisecond)

	// Check the status
	PrintStatus(FI.NC)

	// Select the Benchy
	Files, err := FI.GetFileStructure()
	CommonTestTools.CheckErr(t, "Could not Get file structure", err)
	fullbprint, err := FileManager.PullFile(Files.GetContents(), "full_bed_print.gcode")
	CommonTestTools.CheckErr(t, "could not pull file", err)
	err = selectFile(fullbprint, FI.NC)
	CommonTestTools.CheckErr(t, "could not select file", err)

	// Play
	fmt.Println("Attempting to Play file", fullbprint.GetName())
	err = playFile(FI.NC)
	CommonTestTools.CheckErr(t, "could not set play", err)
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

func setupTestFiles() error {
	gopath := os.Getenv("GOPATH")
	testPrintDir := path.Clean(gopath + "/src/github.com/Ximidar/Flotilla/BuildResources/Test/FlotillaSystemTests/staticfiles")

	if _, err := os.Stat(testPrintDir); os.IsNotExist(err) {
		fmt.Println("Directory does not exist", testPrintDir)
		return err
	}

	// Place all files in locations
	files, err := ioutil.ReadDir(testPrintDir)
	if err != nil {
		fmt.Println("Could not read directory", testPrintDir)
		return err
	}

	if len(files) == 0 {
		return errors.New("no test files")
	}

	for _, file := range files {
		err := Copy(path.Clean(testPrintDir+"/"+file.Name()), path.Clean(TestLocation+"/GCODE/"+file.Name()))
		if err != nil {
			fmt.Println("Couldn't copy file", file.Name())
			return err
		}
	}

	return nil
}
func Copy(src, dst string) error {
	fmt.Printf("Copying file#######\nfrom %v\nto %v\n", src, dst)
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func selectFile(file *FS.File, nc *nats.Conn) error {

	if file == nil || nc == nil {
		fmt.Println("Either the nats connection or the file is invalid", file, nc)
		return errors.New("either the nats connection or the file is invalid")
	}

	selectAction, err := FS.NewFileAction(FS.FileAction_SelectFile, file.GetPath())
	if err != nil {
		fmt.Println("Could not create file action", err)
		return err
	}

	reply, err := FS.SendAction(nc, 5*time.Second, selectAction)
	if err != nil {
		fmt.Println("Could not send file action", err)
		return err
	}

	rString := new(DS.ReplyString)
	err = json.Unmarshal(reply.Data, &rString)
	if err != nil {
		fmt.Println("Couldn't read the response", string(reply.Data))
		return err
	}
	if rString.Success {
		return nil
	}
	return errors.New(string(rString.Message))
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
