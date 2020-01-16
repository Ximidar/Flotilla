package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js" // build with GOOS=js GOARCH=wasm

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/golang/protobuf/proto"
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
		fmt.Println("this ", this)
		fmt.Println("args", args)
		callback := args[len(args)-1:][0]

		url := "http://127.0.0.1:5000/api/getfiles"
		result, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(result.Body)
		result.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		files := new(FS.File)
		err = proto.Unmarshal(body, files)
		if err != nil {
			fmt.Println("Could not unmarshal the returned data", err)
			return
		}

		fmt.Println(files)

	}()
	return nil
}
