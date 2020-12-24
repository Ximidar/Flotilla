package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func (fw *FlotillaWeb) setupWebSocket() {
	// variable for websocket
	fw.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	// make read and write chans
	fw.wsRead = make(chan []byte, 1000)
	fw.wsWrite = make(chan []byte, 1000)

	go fw.websocketReader()
	go fw.websocketWriter()

}

func (fw *FlotillaWeb) websocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket handler activated!")
	ws, err := fw.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error with websocket conn: ", err)
		return
	}

	fw.websockets = append(fw.websockets, ws)

}

// function to read and write to the websocket
func (fw *FlotillaWeb) websocketReader() {

	for { //ever
		deleteWS := make([]int, 0)
		for index, ws := range fw.websockets {
			mt, mess, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("Error with reader!", err)
				deleteWS = append(deleteWS, index)
			}

			fmt.Println("Got message type of ", mt)
			fmt.Println("Mess: ", string(mess))
			fw.wsRead <- mess
		}
		fw.mux.Lock()
		for offset, val := range deleteWS {
			fw.websockets = append(fw.websockets[:val-offset],
				fw.websockets[(val+1)-offset:]...)
		}
		fw.mux.Unlock()

	}

}

func (fw *FlotillaWeb) websocketWriter() {
	for {
		select {
		case writeMess := <-fw.wsWrite:
			fw.mux.Lock()
			deleteWS := make([]int, 0)
			for index, ws := range fw.websockets {
				err := ws.WriteMessage(websocket.TextMessage, writeMess)
				if err != nil {
					fmt.Println("Error with Writer ", err)
					deleteWS = append(deleteWS, index)
				}
			}
			for offset, val := range deleteWS {
				fw.websockets = append(fw.websockets[:val-offset],
					fw.websockets[(val+1)-offset:]...)
			}
			fw.mux.Unlock()

		}
	}
}
