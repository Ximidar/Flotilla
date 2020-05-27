package GetNats

import (
	"os"
	"testing"

	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
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

func Clean() {
	os.RemoveAll(TestLocation)
}

func Test_GetNats(t *testing.T) {
	Setup()

	root, _ := RootFolder.GenerateRootFolder(TestLocation + "/Flotilla")

	DownloadNats(root)
	UnzipAndClean(root)
}
