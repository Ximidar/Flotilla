/*
* @Author: Ximidar
* @Date:   2019-01-13 15:38:04
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-22 16:18:39
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

	"github.com/nats-io/gnatsd/test"
	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
	"github.com/ximidar/Flotilla/CommonTools/NatsConnect"
	DS "github.com/ximidar/Flotilla/DataStructures"
	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

var TestLocation = "/tmp/FlotillaSystem/Test/Flotilla"

// StartTestFlotilla will start a Flotilla instance with the most recently compiled binaries
// Make sure you actually compile Flotilla before running, This also does not include
// A NATS instance yet. So you will have to supply that for now
func StartTestFlotilla() (chan bool, error) {
	os.RemoveAll(TestLocation)
	server := test.RunDefaultServer()
	// wait for server startup to happen
	<-time.After(200 * time.Millisecond)
	_, err := NatsConnect.DefaultConn(nats.DefaultURL, "testcon")
	if err != nil {
		panic(err)
	}

	NodeLauncher, err := CreateProcess("NodeLauncher", "NodeLauncher",
		"CreateRoot", "-p", TestLocation, "-a", "amd64", "-l", "true")
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
		NodeLauncher.Process.Kill()
		<-time.After(2 * time.Second)
		killerr := syscall.Kill(NodeLauncher.Process.Pid, syscall.SIGKILL)
		if killerr != nil {
			fmt.Println(killerr)
		}
		fmt.Println("Everything has been killed")
		exitChan <- true
	}

	StartProc := func() {
		NodeLauncher.Start()
		select {
		case <-exitChan:
			killEverything()
			os.RemoveAll(TestLocation)
			server.Shutdown()
			return
		}
	}
	go StartProc()

	// Wait a little bit for the server to start up before continuing the test
	<-time.After(5 * time.Second)

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
	exitChan, err := StartTestFlotilla()
	CommonTestTools.CheckErr(t, "Could not start Flotilla test", err)
	FI, err := FlotillaInterface.NewFlotillaInterface()
	CommonTestTools.CheckErr(t, "Could not create new flotilla interface", err)

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
	os.RemoveAll("/tmp/fakeprinter")
	serial := FakeSerialDevice.NewFakeSerial()
	go serial.ReadMaster()
	go run(serial, exitChan)

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
	fullbprint := Files["root"].Contents["full_bed_print.gcode"]
	err = selectFile(fullbprint, FI.NC)
	CommonTestTools.CheckErr(t, "could not select file", err)

	// Play
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
	testPrintDir := path.Clean(gopath + "/src/github.com/ximidar/Flotilla/BuildResources/Test/FlotillaSystemTests/staticfiles")

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

func selectFile(file *Files.File, nc *nats.Conn) error {

	if file == nil || nc == nil {
		fmt.Println("Either the nats connection or the file is invalid", file, nc)
		return errors.New("either the nats connection or the file is invalid")
	}

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
