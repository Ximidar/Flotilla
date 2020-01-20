package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"syscall/js" // build with GOOS=js GOARCH=wasm

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/golang/protobuf/proto"
)

const (
	urlBase = "http://127.0.0.1:5000"
)

var CommCallbacks []js.Value

func main() {
	fmt.Println("Hello WASM!")
	registerCallbacks()
	fmt.Println("Loaded WASM Callbacks")
	go OpenWS()
	select {} // never exit
}

func registerCallbacks() {
	js.Global().Set("hello_wasm", js.FuncOf(Hello))
	js.Global().Set("flot_get_files", js.FuncOf(GetFiles))
	js.Global().Set("flot_register_comm_callback", js.FuncOf(RegisterCommCallback))
}

// Hello is a simple experiment with WASM
func Hello(value js.Value, args []js.Value) interface{} {
	fmt.Println("Hello WASM not Main!")
	return nil
}

func RegisterCommCallback(value js.Value, args []js.Value) interface{} {
	go func() {
		callback := args[len(args)-1:][0]
		CommCallbacks = append(CommCallbacks, callback)
	}()
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

func OpenWS() {
	// initialize comm callbacks
	CommCallbacks = make([]js.Value, 10)

	// get connection string
	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:5000",
		Path:   "/api/ws",
	}
	fmt.Println("Attempting connection to: ", u.String())

	ws := js.Global().Get("WebSocket").New(u.String())
	ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("open")

		ws.Call("send", "Hello From WASM")
		return nil
	}))

	ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mess := args[0].Get("data")

		for _, c := range CommCallbacks {
			c.Invoke(mess)
		}

		return nil
	}))

	// ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	// 	fmt.Println("Got error with ws: ", err)
	// 	return
	// }

	// for {
	// 	_, mess, err := ws.ReadMessage()
	// 	if err != nil {
	// 		fmt.Println("Got error with ws: ", err)
	// 	}
	// 	fmt.Println("RECV: ", string(mess))
	// 	ws.WriteMessage(websocket.TextMessage, []byte("Returned Aloha"))
	// }
}
