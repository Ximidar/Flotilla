package api

import (
	"net/http"
	"path"
)

func (fw *FlotillaWeb) IndexHandler(w http.ResponseWriter, req *http.Request) {
	// TODO figure out how to do this better
	index := path.Join(fw.workingDir, "/dist/index.html")
	http.ServeFile(w, req, index)

}

// SetupFileServer is a helper function to serve the frontend files
func (fw *FlotillaWeb) SetupFileServer() error {

	staticFiles := path.Join(fw.workingDir, "/dist/")
	fs := http.FileServer(http.Dir(staticFiles))
	fw.r.PathPrefix("/").Handler(fs)

	return nil
}
