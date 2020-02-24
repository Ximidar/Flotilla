package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	DS "github.com/Ximidar/Flotilla/DataStructures"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

// SelectFile will take in a file object and select it on Nats
func (fw *FlotillaWeb) SelectFile(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error SelectFile: Cannot ReadAll of the Body")
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
	body = body[:blobLength]
	fmt.Println(len(body))

	SelectFile := new(FS.File)
	err = proto.Unmarshal(body, SelectFile)
	if err != nil {
		fmt.Println("Error SelectFile: Cannot Unmarshal proto")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = fw.selectFile(SelectFile)
	if err != nil {
		fmt.Println("Error Could not Select file")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (fw *FlotillaWeb) selectFile(file *FS.File) error {
	selectAction, err := FS.NewFileAction(FS.FileAction_SelectFile, file.Path)
	if err != nil {
		return err
	}

	reply, err := FS.SendAction(fw.Nats, 5*time.Second, selectAction)
	if err != nil {
		return err
	}

	rSTR, err := fw.returnReplyString(reply)
	if err != nil {
		return err
	}
	if rSTR.Success {
		return nil
	}
	return errors.New(string(rSTR.Message))
}

func (fw *FlotillaWeb) returnReplyString(msg *nats.Msg) (*DS.ReplyString, error) {
	msgdata := DS.ReplyString{}

	// unmarshal msg data
	err := json.Unmarshal(msg.Data, &msgdata)
	if err != nil {
		return nil, err
	}

	return &msgdata, nil
}
