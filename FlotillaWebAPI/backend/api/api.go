package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nats-io/go-nats"
)

// Global

// Nats is the access variable to the nats server
var Nats *nats.Conn

//Serve will serve the api
func Serve(port int, directory string) {

	FlotillaWeb := NewFlotillaWeb(port, directory)

	http.Handle("/", FlotillaWeb)
	http.HandleFunc("/api/ws", FlotillaWeb.websocketHandler)

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

	// TODO figure out why websockets mess with the file upload function
	go fw.setupWebSocket()

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

func (fw *FlotillaWeb) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// TODO add API key handling
	fmt.Println("Incoming Request!")

	// Lets Gorilla work
	fw.r.ServeHTTP(rw, req)
}
