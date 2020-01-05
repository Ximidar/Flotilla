package main

import (
	"fmt"
	"syscall/js" // build with GOOS=js GOARCH=wasm
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Hello WASM!")
	registerCallbacks()
	fmt.Println("Loaded WASM Callbacks")
	<-c
}

func registerCallbacks() {
	js.Global().Set("hello_wasm", js.FuncOf(Hello))
}

// Hello is a simple experiment with WASM
func Hello(value js.Value, args []js.Value) interface{} {
	fmt.Println("Hello WASM not Main!")
	return nil
}
