package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/golang/protobuf/proto"
)

// Status will take care of play / pause / resume / cancel commands

func (fw *FlotillaWeb) setupStatus() {
	var err error
	fw.Node, err = PlayStructures.NewRegisteredNode("WebServer", fw.Nats)
	if err != nil {
		fmt.Println("Error with registering node: ", err)
		panic(err)
	}

	fw.Node.StatusObserver.AddFunction(PlayStructures.AnyStatusUpdate,
		fw.AnyStatusUpdate)

}

// AnyStatusUpdate will be called whenever there is a status update over NATS
func (fw *FlotillaWeb) AnyStatusUpdate() {
	fw.wsWrite <- []byte("NewStatus")
}

// GetStatus will get the current status of the flotilla server
func (fw *FlotillaWeb) GetStatus(w http.ResponseWriter, req *http.Request) {
	status := fw.Node.StatusObserver.CurrentStatus
	fmt.Println(status)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(status.CurrentAction))
}

// ChangeStatus will request a status change
func (fw *FlotillaWeb) ChangeStatus(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error ChangeStatus: Cannot ReadAll of the Body")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	blobLength, err := strconv.Atoi(req.Header.Get("blob-length"))
	if err != nil {
		fmt.Println("User did not include blob-length in header")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Include blob-length in the header"))
	}
	fmt.Println(body)
	body = body[:blobLength]
	fmt.Println(len(body))

	// turn body into action
	action := new(PlayStructures.Action)
	err = proto.Unmarshal(body, action)
	if err != nil {
		fmt.Println("Error ChangeStatus: Cannot Unmarshal proto")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Ask for action
	err = PlayStructures.ProposeAction(action.GetAction(), fw.Nats)
	if err != nil {
		fmt.Println("Error ChangeStatus: Cannot Propose Action")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}
