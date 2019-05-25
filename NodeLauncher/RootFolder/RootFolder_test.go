/*
* @Author: Ximidar
* @Date:   2019-02-04 16:32:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 17:19:33
 */

package RootFolder

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
)

const (
	TestLocation = "/tmp/RootFolderTest/"
)

func Setup() {
	if _, err := os.Stat(TestLocation); os.IsNotExist(err) {
		os.Mkdir(TestLocation, 0755)
	} else {
		Clean()
		os.Mkdir(TestLocation, 0755)
	}

}

func CreatePaths() {
	// Folders
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+GCODEFolderName), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName+"/"+BinCoreSubfolder), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName+"/"+BinExtraSubfolder), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+WorkspaceName), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+WorkspaceName+"/"+WSLogs), 0755)

	// Files
	file, err := os.Create(path.Clean(TestLocation + "/Flotilla" + "/" + ConfigName))
	if err != nil {
		fmt.Println("ERROR Could not create config file")
		return
	}
	file.Write([]byte("Config stuff"))
	file.Close()
	file, err = os.Create(path.Clean(TestLocation + "/Flotilla" + "/" + StartServerName))
	if err != nil {
		fmt.Println("ERROR Could not create start server file")
		return
	}
	file.Write([]byte("echo \"Flotilla Start\""))
	file.Close()
}

func Clean() {
	os.RemoveAll(TestLocation)
}

func Test_joinAndClean(t *testing.T) {
	root := RootFolder{}

	expected := "hello/path"
	obtained := root.JoinAndClean("hello", "path")
	CommonTestTools.CheckEquals(t, expected, obtained)

	expected = "long/path/with/file.name"
	obtained = root.JoinAndClean("long", "path", "with", "file.name")
	CommonTestTools.CheckEquals(t, expected, obtained)

}

func Test_MakeRootFolderWithPreExistingFolder(t *testing.T) {

	// Setup the files for testing
	Setup()
	CreatePaths()

	// Create a RootFolder Object using the test location
	root, err := NewRootFolder(TestLocation + "/Flotilla")
	CommonTestTools.CheckErr(t, "Could not Make RootFolder", err)

	fmt.Println(root.RootPath)

	Clean()
}

func Test_MakeRootFolderWithoutPreExistingFolder(t *testing.T) {
	Clean()
	_, err := NewRootFolder(TestLocation + "/Flotilla")
	if err == nil {
		CommonTestTools.CheckErr(t, "Making RootFolder Without Folder", errors.New("Making RootFolder Object succeeded"))
	}

	// Setup the files for testing
	Setup()
	CreatePaths()

	// Delete one of the files
	os.Remove(path.Clean(TestLocation + "/Flotilla" + "/" + ConfigName))

	// test for setup failure
	_, err = NewRootFolder(TestLocation + "/Flotilla")
	if err == nil {
		CommonTestTools.CheckErr(t, "Making RootFolder Without config file", errors.New("Making RootFolder Object succeeded"))
	}
	Clean()
}

func Test_GeneratingRootFolderFromNothing(t *testing.T) {
	Clean()

	root, err := GenerateRootFolder(TestLocation + "/Flotilla")
	CommonTestTools.CheckErr(t, "Generate Root Folder From Nothing", err)
	fmt.Println(root.RootPath)
}
