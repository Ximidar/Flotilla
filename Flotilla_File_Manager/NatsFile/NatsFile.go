/*
* @Author: Ximidar
* @Date:   2018-10-10 06:10:39
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-05-07 20:26:20
 */

package NatsFile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	version "github.com/Ximidar/Flotilla/CommonTools/versioning"
	DS "github.com/Ximidar/Flotilla/DataStructures"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	CRS "github.com/Ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	PS "github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	FM "github.com/Ximidar/Flotilla/Flotilla_File_Manager/FileManager"
	"github.com/Ximidar/Flotilla/Flotilla_File_Manager/FileStreamer"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

// NatsFile will broadcast file system services to the NATS server.
type NatsFile struct {
	NC *nats.Conn

	FileManager  *FM.FileManager
	FileStreamer *FileStreamer.FileStreamer
	RNode        *PS.RegisteredNode

	version *version.Nats
	config  *config
}

// NewNatsFile Use this function to create a new NatsFile Object.
func NewNatsFile() (nf *NatsFile, err error) {
	nf = new(NatsFile)

	// Make Nats Connection
	nf.NC, err = NatsConnect.DefaultConn(NatsConnect.DockerNATS, FS.Name)
	if err != nil {
		fmt.Printf("Can't connect: %v\n", err)
		return nil, err
	}

	// Attach Version
	nf.version, err = version.NewNats(nf.NC, FS.Name)
	if err != nil {
		fmt.Println("Could not Attach version")
		return nil, err
	}

	// Create Requests
	err = nf.createReqs()
	if err != nil {
		fmt.Println("Could not create reqs")
		return nil, err
	}

	// Create the playstatus
	nf.RNode, err = PS.NewRegisteredNode(FS.Name+"HeartBeat", nf.NC)
	if err != nil {
		fmt.Println("Could not register Node", err)
		return nil, err
	}

	// Get the Config
	// nf.config, err = newConfig(nf.NC)
	// if err != nil {
	// 	fmt.Println("Could not create config object", err)
	// 	return nil, err
	// }
	// err = nf.config.GetConfig()
	// if err != nil {
	// 	fmt.Println("Could not get config vars", err)
	// 	return nil, err
	// }

	// Create File manager
	home := os.Getenv("HOME")
	defaultPath := path.Clean(home + "/gcode")
	nf.FileManager, err = FM.NewFileManager(defaultPath)
	nf.FileStreamer, err = FileStreamer.NewFileStreamer(nf)

	// Setup RNode Observers
	err = nf.setupRNodeObservers()
	if err != nil {
		fmt.Println("Could not setup RNode Observers")
		return nil, err
	}

	// subscribe to NATS Subjects
	err = nf.SubscribeSubjects()
	if err != nil {
		fmt.Println("Couldn't subscribe")
		return nil, err
	}

	return nf, nil
}

func (nf *NatsFile) setupRNodeObservers() error {
	var err error
	_, err = nf.RNode.StatusObserver.AddFunction(PS.CANCEL, nf.FileStreamer.Cancel)
	return err
}

func (nf *NatsFile) createReqs() (err error) {
	// Assign each to err. At the end if there are errors we will
	// return the most recent error
	_, err = nf.NC.Subscribe(FS.RunFileAction, nf.runFileAction)
	if err != nil {
		fmt.Println("Couldn't register select file")
		return err
	}

	//TODO Add in File modifiers.
	// Add a function in FFM that will accept a http(s) url and get a file.

	return nil
}

// SendStatus will fire a status event to the play status
func (nf *NatsFile) SendStatus(status string) {
	err := nf.RNode.UpdateNode(status)
	if err != nil {
		fmt.Println("Could not update node because:", err)
	}
}

// NotifyDone will request a Done Status be applied to the server
func (nf *NatsFile) NotifyDone() {
	pstatus, err := PS.NewPlayAction(PS.DONE)
	if err != nil {
		fmt.Println("Cannot send DONE status because", err.Error())
		return
	}
	err = pstatus.Send(nf.NC)
	if err != nil {
		fmt.Println("Could not send DONE status because", err.Error())
	}

}

// PackageResponse will respond to the subject with a positive message
func (nf *NatsFile) PackageResponse(success bool, message string, subject string) {
	rm := new(DS.ReplyString)
	rm.Success = true
	rm.Message = message

	mReply, err := json.Marshal(rm)
	if err != nil {
		fmt.Println("Could not Marshal a Response", err, rm)
		return
	}

	err = nf.NC.Publish(subject, mReply)
	if err != nil {
		fmt.Println("Could not Send a Negative Reply", err, rm)
	}
}

func (nf *NatsFile) runFileAction(msg *nats.Msg) {
	action, err := FS.NewFileActionFromMSG(msg)
	if err != nil {
		nf.PackageResponse(false, err.Error(), msg.Reply)
	}

	switch action.GetAction() {
	case FS.FileAction_AddFile:
		nf.PackageResponse(false, "Not Implemented", msg.Reply)
	case FS.FileAction_DeleteFile:
		nf.PackageResponse(false, "Not Implemented", msg.Reply)
	case FS.FileAction_MoveFile:
		nf.PackageResponse(false, "Not Implemented", msg.Reply)

	case FS.FileAction_SelectFile:
		err := nf.selectFile(action)
		if err != nil {
			nf.PackageResponse(false, err.Error(), msg.Reply)
		}
		nf.PackageResponse(true, "file selected", msg.Reply)
	case FS.FileAction_GetFileStructure:
		structure := nf.FileManager.Structure
		sb, err := proto.Marshal(structure)
		if err != nil {
			fmt.Println("Could not send file structure")
			nf.PackageResponse(false, err.Error(), msg.Reply)
		}
		nf.NC.Publish(msg.Reply, sb)
	}
}

func (nf *NatsFile) selectFile(action *FS.FileAction) error {

	file, err := nf.FileManager.GetFileByPath(action.GetPath())
	if err != nil {
		// If we get an error, return an error
		return err
	}

	err = nf.FileStreamer.SelectFile(file)
	if err != nil {
		return err
	}

	return nil

}

func (nf *NatsFile) getStructureJSON(action *FS.FileAction) ([]byte, error) {
	jsonStructure, err := nf.FileManager.GetJSONStructure()
	if err != nil {
		return []byte(nil), err
	}

	return jsonStructure, nil
}

// SubscribeSubjects will subscribe NatsFile to different NATS Subjects
func (nf *NatsFile) SubscribeSubjects() error {

	_, err := nf.NC.Subscribe(FS.RequestLines, nf.RequestLines)
	return err
}

// RequestLines will return the requested amount of lines over NATS
func (nf *NatsFile) RequestLines(msg *nats.Msg) {

	// figure out the amount of requested lines
	rl := new(CRS.RequestLines)
	err := proto.Unmarshal(msg.Data, rl)
	if err != nil {
		fmt.Println("Error: Could not parse requested lines", err)
		return
	}

	// Request the next N lines
	amount := int(rl.GetAmount())
	returnVessel := new(CRS.ReturnLines)

	// attempt to get the requested amount of lines. If there aren't any, then
	// we will exit early and return the amount of lines we recieved
	returnVessel.Lines, err = nf.FileStreamer.GetLines(amount)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error: Could not get lines", err)
		}
		returnVessel.EOF = true
	}

	brv, err := proto.Marshal(returnVessel)
	if err != nil {
		fmt.Println("Error: Could not return requested lines", err)
		return
	}

	nf.NC.Publish(msg.Reply, brv)

}

// ProgressUpdate is part of the adapter to send to the File Streamer
func (nf *NatsFile) ProgressUpdate(file *FS.File, currentLine uint64, readBytes uint64) {
	fp := new(FS.FileProg)
	fp.FileName = path.Base(file.Name)
	fp.Size = file.Size
	fp.CurrentLine = currentLine
	fp.BytesRead = readBytes
	fp.Progress = float32(readBytes) / float32(file.Size) * 100

	mfp, err := proto.Marshal(fp)
	if err != nil {
		fmt.Println("ERROR!", err)
		return
	}

	nf.NC.Publish(FS.FileProgress, mfp)
}
