package api

import (
	"fmt"
	"io/ioutil"
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

	go fw.ReceivedComm()
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
func (fw *FlotillaWeb) ReceivedComm() {
	for {
		select {
		case mess := <-fw.wsRead:
			fmt.Println("Got Message from WebSocket!", string(mess))
			cm := new(CS.CommMessage)
			cm.Message = string(mess)

			cmProto, err := proto.Marshal(cm)
			if err != nil {
				fmt.Println("ReceivedComm Could not package: ", mess)
				continue
			}

			_, err = fw.MakeNatsRequest(CS.WriteComm, cmProto)
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
	options, err := fw.MakeNatsRequest(CS.ListOptions, EMPTY)
	if err != nil {
		fmt.Println("Error while listing Comm Ports: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(options)

}

// CommInit will take in a comm init structure then Connect the Comm. POST function
func (fw *FlotillaWeb) CommInit(w http.ResponseWriter, r *http.Request) {

	// get the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error CommInit: Cannot ReadAll of the Body")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// attempt to unmarshal the CommInit
	commInit := new(CS.InitComm)
	err = proto.Unmarshal(body, commInit)
	if err != nil {
		fmt.Println("CommInit Could not unmarshal Proto!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//send the CommInit out over nats
	_, err = fw.MakeNatsRequest(CS.InitializeComm, body)
	if err != nil {
		fmt.Println("CommInit Could not make Nats Request!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}

// CommConnect will make an attempt to connect the comm
func (fw *FlotillaWeb) CommConnect(w http.ResponseWriter, r *http.Request) {
	resp, err := fw.MakeNatsRequest(CS.ConnectComm, EMPTY)
	if err != nil {
		fmt.Println("CommConnect Could not make Nats Request!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// CommDisconnect will make an attempt to disconnect the comm
func (fw *FlotillaWeb) CommDisconnect(w http.ResponseWriter, r *http.Request) {
	resp, err := fw.MakeNatsRequest(CS.DisconnectComm, EMPTY)
	if err != nil {
		fmt.Println("CommDisconnect Could not make Nats Request!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
