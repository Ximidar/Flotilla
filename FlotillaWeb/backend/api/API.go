package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nats-io/go-nats"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// Global

// Nats is the access variable to the nats server
var Nats *nats.Conn

//Serve will serve the api
func Serve(port int, directory string) {

	FlotillaWeb := NewFlotillaWeb(port, directory)

	http.Handle("/", FlotillaWeb)
	// http.HandleFunc("/api/ws", FlotillaWeb.websocketHandler)

	//Make CORS
	headersOK := handlers.AllowedHeaders([]string{"Accept",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"blob-length"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("Serving %s on HTTP port: %v\n", directory, port)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%v", port),
			handlers.CORS(
				headersOK,
				originsOk,
				methodsOK,
			)(FlotillaWeb.r)))
}

// NewFlotillaWeb will create a new flotilla webserver
func NewFlotillaWeb(port int, directory string) *FlotillaWeb {

	fw := new(FlotillaWeb)
	var err error

	// setup nats
	fw.Nats, err = NatsConnect.DefaultConn(NatsConnect.LocalNATS, "flotillaInterface")
	if err != nil {
		log.Fatal(err)
	}

	go fw.setupRouter()
	go fw.setupFileServer(directory)
	// TODO figure out why websockets mess with the file upload function
	// go fw.setupWebSocket()

	// setup Flotilla stuff
	go fw.setupCommRelay()
	go fw.setupStatus()

	return fw

}

// FlotillaWeb will be the main handler for incoming requests
type FlotillaWeb struct {
	// webserver
	fs http.Handler
	r  *mux.Router

	//websocket
	websockets []*websocket.Conn
	upgrader   websocket.Upgrader
	wsWrite    chan []byte
	wsRead     chan []byte

	// nats
	Nats *nats.Conn

	// flotilla
	Node *PlayStructures.RegisteredNode
	mux  sync.Mutex
}

func (fw *FlotillaWeb) setupFileServer(directory string) {
	fw.fs = http.FileServer(FileSystem{http.Dir(directory)})
	fw.r.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fw.fs)).Methods("GET")

}

func (fw *FlotillaWeb) setupRouter() {
	fw.r = mux.NewRouter()

	// Files
	fw.r.HandleFunc("/api/getfiles", fw.GetFiles).Methods("GET")
	fw.r.HandleFunc("/api/selectfile", fw.SelectFile).Methods("POST")
	fw.r.HandleFunc("/api/file", fw.UploadFile).Methods("POST")

	// Status
	fw.r.HandleFunc("/api/status", fw.GetStatus).Methods("GET")
	fw.r.HandleFunc("/api/status", fw.ChangeStatus).Methods("POST")

	// WebSockets
	fw.r.HandleFunc("/api/ws", fw.websocketHandler)

}

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

func (fw *FlotillaWeb) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// TODO add API key handling
	fmt.Println("Incoming Request!")

	// Lets Gorilla work
	fw.r.ServeHTTP(rw, req)
}

// GetFiles will get the files from Nats and return them
func (fw *FlotillaWeb) GetFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gettin Files!")
	fileRequest, err := FS.NewFileAction(FS.FileAction_GetFileStructure, "")
	if err != nil {
		w.Write([]byte("Error"))
		fmt.Printf("Error: %v\n", err)
		return
	}

	msg, err := FS.SendAction(fw.Nats, 5*time.Second, fileRequest)
	if err != nil {
		w.Write([]byte("Error"))
		fmt.Printf("Error: %v\n", err)
		return
	}

	w.Write(msg.Data)
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

func (fw *FlotillaWeb) setupCommRelay() {
	_, err := fw.Nats.Subscribe(CS.ReadLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.ReadLine, err)
	}
	_, err = fw.Nats.Subscribe(CS.WriteLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.WriteLine, err)
	}
}

// CommRelay will receive COMM messages from NATS
func (fw *FlotillaWeb) CommRelay(msg *nats.Msg) {
	cm := new(CS.CommMessage)
	err := proto.Unmarshal(msg.Data, cm)
	if err != nil {
		fmt.Println("Could not deconstruct proto message for commrelay")
	}
	fw.wsWrite <- []byte(cm.Message)
}
