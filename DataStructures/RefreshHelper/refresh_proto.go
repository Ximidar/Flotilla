package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// ######################### USE PBJS!

func main() {
	RefreshDataStructures()
}

// RefreshDataStructures will find all protoc files and convert them to js and go
func RefreshDataStructures() {
	GOPATH := os.Getenv("GOPATH")
	DataStructures := GOPATH + "/src/github.com/Ximidar/Flotilla/DataStructures"
	webproto := GOPATH + "/src/github.com/Ximidar/Flotilla/FlotillaWeb/src/js_proto/"
	PBJSCLI := GOPATH + "/src/github.com/Ximidar/Flotilla/FlotillaWeb/node_modules/protobufjs/bin/pbjs"
	// Check for existance
	if _, err := os.Stat(DataStructures); os.IsNotExist(err) {
		err := fmt.Errorf("PATH: %v Does not exist", DataStructures)
		panic(err)
	}
	if _, err := os.Stat(webproto); os.IsNotExist(err) {
		err := fmt.Errorf("PATH: %v Does not exist", webproto)
		panic(err)
	}
	if _, err := os.Stat(PBJSCLI); os.IsNotExist(err) {
		fmt.Println("You might need to install protoc minimal js with npm")
		err := fmt.Errorf("PATH: %v Does not exist", PBJSCLI)
		panic(err)
	}

	// Check for protoc
	_, err := CommandWithArgs("protoc", "--version")
	if err != nil {
		err := fmt.Errorf("protoc Not installed")
		panic(err)
	}

	// Get all .proto files
	protoFiles := GetProtoFiles(DataStructures)

	// compile the proto file

	pbjsArgs := []string{"-t", "static-module", "-w", "commonjs", "-o"}
	for _, proto := range protoFiles {
		base := path.Base(proto)
		jsBase := strings.Replace(base, ".proto", "_pb.js", 1)
		dumpDir := strings.Replace(proto, base, "", 1)

		// Make go protoc files
		gargs := make([]string, 0)
		gargs = append(gargs, "--proto_path="+dumpDir, "--go_out="+dumpDir, proto)
		CommandWithArgs("protoc", gargs...)

		// Make pbjs files
		jargs := make([]string, 0)
		jargs = append(jargs, pbjsArgs...)
		jargs = append(jargs, dumpDir+jsBase, proto)
		CommandWithArgs(PBJSCLI, jargs...)

		// Make pbjs files in FlotillaWeb
		jargs = make([]string, 0)
		jargs = append(jargs, pbjsArgs...)
		jargs = append(jargs, webproto+jsBase, proto)
		CommandWithArgs(PBJSCLI, jargs...)
	}

}

func logger(name string, message string) {
	fmt.Println(name, ": ", message)
}

// CommandWithArgs Runs Protoc with associated arguments
func CommandWithArgs(bin string, args ...string) (*exec.Cmd, error) {
	var err error

	// setup the command
	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run
	err = cmd.Run()

	// return the result
	return cmd, err

}

// GetProtoFiles Finds all Protoc Files
func GetProtoFiles(dir string) []string {
	protoFiles := make([]string, 0)
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Err: ", err)
				return err
			}
			if strings.HasSuffix(path, ".proto") {
				protoFiles = append(protoFiles, path)
			}
			return nil
		})
	if err != nil {
		fmt.Println("Failed my walk")
		panic(err)
	}
	return protoFiles

}
