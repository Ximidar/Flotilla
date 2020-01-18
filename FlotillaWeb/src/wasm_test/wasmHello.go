package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js" // build with GOOS=js GOARCH=wasm

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/golang/protobuf/proto"
)

const (
	urlBase = "http://127.0.0.1:5000"
)

func main() {
	fmt.Println("Hello WASM!")
	registerCallbacks()
	fmt.Println("Loaded WASM Callbacks")
	select {} // never exit
}

func registerCallbacks() {
	js.Global().Set("hello_wasm", js.FuncOf(Hello))
	js.Global().Set("flot_get_files", js.FuncOf(GetFiles))
}

// Hello is a simple experiment with WASM
func Hello(value js.Value, args []js.Value) interface{} {
	fmt.Println("Hello WASM not Main!")
	return nil
}

// GetFiles will eventually attempt to get some files
func GetFiles(this js.Value, args []js.Value) interface{} {
	go func() {
		callback := args[len(args)-1:][0]

		url := urlBase + "/api/getfiles"
		result, err := http.Get(url)
		if err != nil {
			printErr(err)
		}

		body, err := ioutil.ReadAll(result.Body)
		result.Body.Close()
		if err != nil {
			printErr(err)
		}

		files := new(FS.File)
		err = proto.Unmarshal(body, files)
		if err != nil {
			printErr(err)
			return
		}

		jsonFiles, err := json.Marshal(files)
		if err != nil {
			printErr(err)
			return
		}

		callback.Invoke(string(jsonFiles))

	}()
	return nil
}

func printErr(err error) {
	fmt.Println("Error!", err)
}
