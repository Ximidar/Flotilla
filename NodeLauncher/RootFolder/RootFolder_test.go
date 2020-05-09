/*
* @Author: Ximidar
* @Date:   2019-02-04 16:32:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 17:19:33
 */

package RootFolder

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/Ximidar/Flotilla/BuildResources/Test/CommonTestTools"
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
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName+"/amd64"), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName+"/arm64"), 0755)
	os.Mkdir(path.Clean(TestLocation+"/Flotilla"+"/"+BinName+"/armhf"), 0755)
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

func Test_GeneratingRootFolderFromNothing(t *testing.T) {
	Clean()

	root, err := GenerateRootFolder(TestLocation + "/Flotilla")
	CommonTestTools.CheckErr(t, "Generate Root Folder From Nothing", err)
	fmt.Println(root.RootPath)
}
