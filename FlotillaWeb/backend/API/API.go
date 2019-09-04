package API

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"

	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/CommonTools/NatsConnect"
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
	var err error
	Nats, err = NatsConnect.DefaultConn(nats.DefaultURL, "flotillaInterface")
	if err != nil {
		log.Fatal(err)
	}

	// Create Mux
	mux := mux.NewRouter()

	// UI serve
	fileServer := http.FileServer(FileSystem{http.Dir(directory)})
	mux.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), fileServer)).Methods("GET")

	// API Endpoints
	mux.HandleFunc("/api/getfiles", GetFiles).Methods("GET")

	// make main router
	http.Handle("/", &FlotillaWeb{mux})

	log.Printf("Serving %s on HTTP port: %v\n", directory, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// FlotillaWeb will be the main handler for incoming requests
type FlotillaWeb struct {
	r *mux.Router
}

func (s *FlotillaWeb) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("root command!")
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}

func WriteBasicHeaders(rw http.ResponseWriter) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// GetFiles will get the files from Nats and return them
func GetFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gettin Files!")
	fileRequest, err := FS.NewFileAction(FS.FileAction_GetFileStructure, "")
	if err != nil {
		w.Write([]byte("Error"))
		fmt.Printf("Error: %v\n", err)
		return
	}

	msg, err := FS.SendAction(Nats, 5*time.Second, fileRequest)
	if err != nil {
		w.Write([]byte("Error"))
		fmt.Printf("Error: %v\n", err)
		return
	}

	WriteBasicHeaders(w)
	w.Write(msg.Data)
}
