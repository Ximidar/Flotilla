/*
* @Author: Ximidar
* @Date:   2018-08-25 10:12:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 19:26:40
 */

package FlotillaInterface

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	DS "github.com/Ximidar/Flotilla/DataStructures"
	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"

	CRS "github.com/Ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
)

// EMPTY []byte for giving an empty payload
var EMPTY []byte

// FlotillaInterface is an interface to the Nats server
type FlotillaInterface struct {
	NC       *nats.Conn
	EmitLine chan string

	Timeout time.Duration
}

// NewFlotillaInterface will construct the FlotillaInterface
func NewFlotillaInterface() (*FlotillaInterface, error) {
	fi := new(FlotillaInterface)
	var err error
	fi.NC, err = NatsConnect.DefaultConn(NatsConnect.LocalNATS, "flotillaInterface")
	fi.Timeout = nats.DefaultTimeout

	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
		return nil, err
	}

	return fi, nil
}

// MakeRequest will construct a Nats Request and send it
func (fi *FlotillaInterface) MakeRequest(subject string, payload []byte) ([]byte, error) {

	msg, err := fi.NC.Request(subject, payload, fi.Timeout)

	if err != nil {
		if err == nats.ErrTimeout {
			return nil, err
		}
		return nil, err
	}

	fi.Timeout = nats.DefaultTimeout

	return msg.Data, nil

}

// CommSetConnectionOptions will set the connection options to the Comm Object
func (fi *FlotillaInterface) CommSetConnectionOptions(port string, baud int) error {

	initComm := new(CS.InitComm)
	initComm.Port = port
	initComm.Baud = int32(baud)
	payload, err := proto.Marshal(initComm)
	if err != nil {
		fmt.Println("Could not marshal InitComm", initComm)
	}
	fi.Timeout = 10 * time.Second
	call, err := fi.MakeRequest(CS.InitializeComm, payload)

	if err != nil {
		return err
	}

	response := new(DS.ReplyString)
	err = json.Unmarshal(call, response)

	if err != nil {
		log.Println("Could not unmarshal data")
		return err
	}

	log.Printf("\nInitialize Comm\nSuccess: %v\nResponse: %v\n", response.Success, response.Message)

	return nil

}

// CommGetStatus will get the status of the Comm Object
func (fi *FlotillaInterface) CommGetStatus() (*CS.CommStatus, error) {
	call, err := fi.MakeRequest(CS.GetStatus, EMPTY)

	if err != nil {
		fmt.Println("Could not get status")
		return nil, err
	}

	return fi.DeconstructStatus(call)

}

// DeconstructStatus will figure out if the call succeeded or not
func (fi *FlotillaInterface) DeconstructStatus(call []byte) (*CS.CommStatus, error) {
	status := new(CS.CommStatus)
	err := proto.Unmarshal(call, status)

	if err != nil {
		fmt.Println("Could not deconstruct proto status", string(call))
		return nil, err
	}

	return status, nil

}

// CommConnect will query the Nats object to connect the Comm Object
func (fi *FlotillaInterface) CommConnect() error {
	fi.Timeout = 10 * time.Second //Ten Seconds to Connect
	call, err := fi.MakeRequest(CS.ConnectComm, EMPTY)

	if err != nil {
		fmt.Println("Could not connect")
		return err
	}

	reply := new(DS.ReplyString)
	err = json.Unmarshal(call, reply)
	if err != nil {
		return err
	}

	//log.Printf("\nInitialize Comm\nSuccess: %v\nResponse: %v\n", reply.Success, reply.Message)

	if !reply.Success {
		return fmt.Errorf("Could not connect: %v", reply.Message)
	}

	return nil
}

// CommDisconnect will query the Nats Server to Disconnect the Comm Object
func (fi *FlotillaInterface) CommDisconnect() error {
	fi.Timeout = 10 * time.Second //Ten Seconds to disconnect
	call, err := fi.MakeRequest(CS.DisconnectComm, EMPTY)

	if err != nil {
		fmt.Println("Could not disconnect")
		return err
	}

	reply := new(DS.ReplyString)
	err = json.Unmarshal(call, reply)
	if err != nil {
		return err
	}

	//log.Printf("\nInitialize Comm\nSuccess: %v\nResponse: %v\n", reply.Success, reply.Message)

	if !reply.Success {
		return fmt.Errorf("Could not disconnect: %v", reply.Message)
	}

	return nil
}

// CommGetAvailablePorts will query the Nats Server for all available ports
func (fi *FlotillaInterface) CommGetAvailablePorts() (*CS.Ports, error) {

	uPorts, err := fi.MakeRequest(CS.ListPorts, EMPTY)
	if err != nil {
		log.Println("Could not Get available ports")
		return nil, err
	}

	Ports := new(CS.Ports)
	err = proto.Unmarshal(uPorts, Ports)

	if err != nil {
		log.Println("Could not unmarshal ports")
		return nil, err
	}

	return Ports, nil

}

// CommWrite will write an unstructured command to the CommRelay over nats
func (fi *FlotillaInterface) CommWrite(command string) error {

	line, err := CRS.NewLine(command, 0, false)
	if err != nil {
		return err
	}

	err = CRS.SendLine(fi.NC, line)
	return err
}

// GetFileStructure will use Request the NATS server for the structure of the filesystem
func (fi *FlotillaInterface) GetFileStructure() (*FS.File, error) {

	fileRequest, err := FS.NewFileAction(FS.FileAction_GetFileStructure, "")
	if err != nil {
		return nil, err
	}

	msg, err := FS.SendAction(fi.NC, 5*time.Second, fileRequest)
	if err != nil {
		return nil, err
	}

	fs := new(FS.File)
	err = proto.Unmarshal(msg.Data, fs)
	if err != nil {
		return nil, err
	}

	return fs, nil

}

// ExtractDataJSON will take a flotilla message and detect success, then return the raw json data.
func (fi *FlotillaInterface) ExtractDataJSON(rawData []byte) ([]byte, error) {
	msgdata := DS.ReplyJSON{}

	// unmarshal msg data
	err := json.Unmarshal(rawData, &msgdata)
	if err != nil {
		return nil, err
	}

	if msgdata.Success {
		return msgdata.Message, nil
	}

	return nil, errors.New("JSON Call failed")
}

// SelectAndPlayFile will select a file over Nats and play it.
func (fi *FlotillaInterface) SelectAndPlayFile(file *FS.File) error {
	err := fi.selectFile(file)
	if err != nil {
		b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
		n := runtime.Stack(b, false)
		s := string(b[:n])
		log.Println(s)
		return err
	}

	err = fi.playFile()
	if err != nil {
		b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
		n := runtime.Stack(b, false)
		s := string(b[:n])
		log.Println(s)
		return err
	}
	return nil
}

func (fi *FlotillaInterface) selectFile(file *FS.File) error {
	selectAction, err := FS.NewFileAction(FS.FileAction_SelectFile, file.Path)
	if err != nil {
		return err
	}

	reply, err := FS.SendAction(fi.NC, 5*time.Second, selectAction)
	if err != nil {
		return err
	}

	rSTR, err := fi.returnReplyString(reply)
	if err != nil {
		return err
	}
	if rSTR.Success {
		return nil
	}
	return errors.New(string(rSTR.Message))
}

func (fi *FlotillaInterface) playFile() error {
	err := PlayStructures.ProposeAction(PlayStructures.PLAY, fi.NC)
	if err != nil {
		return err
	}
	return nil
}

func (fi *FlotillaInterface) returnReplyJSON(msg *nats.Msg) (*DS.ReplyJSON, error) {
	msgdata := DS.ReplyJSON{}

	// unmarshal msg data
	err := json.Unmarshal(msg.Data, &msgdata)
	if err != nil {
		return nil, err
	}

	return &msgdata, nil
}

func (fi *FlotillaInterface) returnReplyString(msg *nats.Msg) (*DS.ReplyString, error) {
	msgdata := DS.ReplyString{}

	// unmarshal msg data
	err := json.Unmarshal(msg.Data, &msgdata)
	if err != nil {
		return nil, err
	}

	return &msgdata, nil
}
