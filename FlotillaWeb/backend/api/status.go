package api

import "net/http"

// Status will take care of play / pause / resume / cancel commands

// GetStatus will get the current status of the flotilla server
func GetStatus(w http.ResponseWriter, req *http.Request) {
	hello := []byte("Hello!")
	WriteBasicHeaders(w)
	w.Write(hello)
}

// ChangeStatus will request a status change
func ChangeStatus(w http.ResponseWriter, req *http.Request) {

}
