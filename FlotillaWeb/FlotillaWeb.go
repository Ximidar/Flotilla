package main

import (
	"log"
	"os"

	"github.com/ximidar/Flotilla/FlotillaWeb/backend"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	backend.Execute(dir)
}
