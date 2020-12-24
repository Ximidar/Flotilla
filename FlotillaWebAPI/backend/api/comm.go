package api

import (
	"fmt"
	"net/http"

	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

func (fw *FlotillaWeb) setupCommRelay() {
	_, err := fw.Nats.Subscribe(CS.ReadLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.ReadLine, err)
	}
	_, err = fw.Nats.Subscribe(CS.WriteLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.WriteLine, err)
	}
}

// CommRelay will receive COMM messages from NATS
func (fw *FlotillaWeb) CommRelay(msg *nats.Msg) {
	cm := new(CS.CommMessage)
	err := proto.Unmarshal(msg.Data, cm)
	if err != nil {
		fmt.Println("Could not deconstruct proto message for commrelay")
	}
	fw.wsWrite <- []byte(cm.Message)
}

// ReceivedComm will monitor for new messages from the websocket
func (fw *FlotillaWeb) ReceivedComm(commMsg string) {
	for {
		select {
		case mess := <-fw.wsRead:
			fmt.Println("Got Message from WebSocket!", string(mess))
			// cm := new(CS.CommMessage)
			// cm.Message = string(mess)

			//TODO send message back to Comm
		}
	}
}

// CommStatus will return the current status of Comm. GET Function
func (fw *FlotillaWeb) CommStatus(w http.ResponseWriter, r *http.Request) {

	// Get the Comm Status
	commStatus, err := fw.MakeNatsRequest(CS.GetStatus, EMPTY)

	if err != nil {
		fmt.Println("Error with CommStatus", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}

	// Package it up and send back
	w.WriteHeader(http.StatusOK)
	w.Write(commStatus)

}

// CommOptions will return the available Ports and Speeds
func (fw *FlotillaWeb) CommOptions(w http.ResponseWriter, r *http.Request) {

	// get ports
	ports, err := fw.MakeNatsRequest(CS.ListPorts, EMPTY)
	if err != nil {
		fmt.Println("Error while listing Comm Ports: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}

	speeds := []string{"250000", "230400", "115200", "57600", "38400", "19200", "9600"}

}

// CommInit will take in a comm init structure then Connect the Comm. POST function
func (fw *FlotillaWeb) CommInit(w http.ResponseWriter, r *http.Request) {

	// get message

	// figure out action

	// do action

}
