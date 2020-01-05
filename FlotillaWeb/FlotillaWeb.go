package main

import (
	"log"
	"os"

	"github.com/Ximidar/Flotilla/FlotillaWeb/backend"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	backend.Execute(dir)
}
