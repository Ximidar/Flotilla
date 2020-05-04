package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

// UploadFile will take in a fileobject and store it in the gcode folder
func (fw *FlotillaWeb) UploadFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Parsing Multipart Form")
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		// return a failure
		fmt.Println("Error UploadFile: Cannot parse multipart form")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("Got an upload request from ", req.RemoteAddr)

	fmt.Println("Getting File Info")
	file, header, err := req.FormFile("file")

	if err != nil {
		// return a failure
		fmt.Println("Error UploadFile: Cannot read file info")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	fmt.Printf("Received file with name %s and size %v\n", header.Filename, header.Size)

	tfile, err := os.OpenFile("/tmp/temp.gcode", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// return a failure
		fmt.Println("Error UploadFile: Could not create temp file")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer tfile.Close()

	io.Copy(tfile, file)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}
