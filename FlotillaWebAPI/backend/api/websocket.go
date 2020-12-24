package api

import (
	"fmt"
	"io/ioutil"
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
	go fw.websocketReadLoop(ws)

}

func (fw *FlotillaWeb) websocketReadLoop(ws *websocket.Conn) {
	for {
		mt, reader, err := ws.NextReader()

		if err != nil {
			ws.Close()
			return
		}

		mess, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println("Got Err from attempting to read Message: ", err)
		}
		//Handle Reader
		fmt.Println("Got message type of ", mt)
		fmt.Println("Mess: ", string(mess))
		fw.wsRead <- mess

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
