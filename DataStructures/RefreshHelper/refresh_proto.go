package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Ximidar/Flotilla/NodeLauncher/Monitor"
)

func main() {
	RefreshDataStructures()
}

func RefreshDataStructures() {
	GOPATH := os.Getenv("GOPATH")
	DataStructures := GOPATH + "/src/github.com/Ximidar/Flotilla/DataStructures"

	// Check for existance
	if _, err := os.Stat(DataStructures); os.IsNotExist(err) {
		err := fmt.Errorf("PATH: %v Does not exist", DataStructures)
		panic(err)
	}

	// Check for protoc
	monitor := ProtocWithArgs("--version")
	if monitor.Error != nil {
		err := fmt.Errorf("protoc Not installed")
		panic(err)
	}

	// Get all .proto files
	protoFiles := GetProtoFiles(DataStructures)

	// compile the proto file

	jsargs := "--js_out=import_style=commonjs,binary:"
	for _, proto := range protoFiles {
		base := path.Base(proto)
		dumpDir := strings.Replace(proto, base, "", 1)

		gargs := make([]string, 0)
		gargs = append(gargs, "--proto_path="+dumpDir, "--go_out="+dumpDir, proto)
		ProtocWithArgs(gargs...)
		jargs := make([]string, 0)
		jargs = append(jargs, "--proto_path="+dumpDir, jsargs+dumpDir, proto)
		ProtocWithArgs(jargs...)
	}

}

func logger(name string, message string) {
	fmt.Println(name, ": ", message)
}

func ProtocWithArgs(args ...string) *Monitor.Monitor {
	var err error

	execPath := "protoc"

	monitor, err := Monitor.NewMonitor(execPath, logger, args...)
	if err != nil {
		panic(err)
	}
	go monitor.ConsumeLines()
	monitor.StartProcess()
	return monitor

}

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
