/*
* @Author: Ximidar
* @Date:   2018-12-18 11:44:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-13 16:35:54
 */

package CommonTestTools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

// CheckErr will take in an error and a message to display for any failed
// errors
func CheckErr(t *testing.T, mess string, err error) {
	if err != nil {
		t.Fatalf("Failed Check from %v, Error: %v", mess, err)
	}
}

// CheckEquals will check if two object equal eachother
func CheckEquals(t *testing.T, object1 interface{}, object2 interface{}) {
	if object1 != object2 {
		t.Fatal("Objects do not match OB1:", object1, "OB2:", object2)
	}
}

// MakeProcess will make a proess ready to run. If it returns nil
// then the process failed to create
func MakeProcess(prog string, args ...string) *exec.Cmd {
	cmd := new(exec.Cmd)
	recoverFunc := func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered From Error:", err)
			cmd = nil
		}
	}
	defer recoverFunc()

	fmt.Printf("Calling Command %s %s\n", prog, strings.Join(args, " "))
	cmd = exec.Command(prog, args...)
	return cmd
}

// FindFlotillaBinFolder will attempt to find the bin of a supplied executable name
// This is for finding executables for flotilla.
func FindFlotillaBinFolder(progname string) (string, error) {

	// Evaluate the gopath environment variable
	GoPath := os.Getenv("GOPATH")

	if GoPath == "" {
		return "", errors.New("No environment variable by that name")
	}

	// Evaluate Flotilla Path
	FlotillaRelPath := "/src/github.com/Ximidar/Flotilla/"

	FlotillaPath := path.Clean(GoPath + FlotillaRelPath)

	if _, err := os.Stat(FlotillaPath); os.IsNotExist(err) {
		return "", fmt.Errorf("Flotilla Folder Does not Exist: %v", FlotillaPath)
	}

	// Find folder for progname
	FlotillaFolder, err := FindFolder(progname, FlotillaPath)
	if err != nil {
		return "", err
	}

	FFBinPath := path.Clean(FlotillaFolder + "/bin/")

	// Every FlotillaFolder should have a bin folder
	if _, err := os.Stat(FFBinPath); os.IsNotExist(err) {
		return "", errors.New("bin Folder does not exist for this Program. Try typing \"make\" in the folder")
	}

	return FFBinPath, nil
}

// FindFolder will attempt to find the folder at a specified directory and return it's
// path
func FindFolder(folderName string, dir string) (string, error) {

	// try to stat the folder
	_, err := os.Stat(dir)
	if err != nil {
		return "", err
	}
	// list directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	lowerFolderName := strings.ToLower(folderName)
	for _, f := range files {
		lowerDirName := strings.ToLower(f.Name())
		//fmt.Printf("Comparing %v to %v\n", lowerFolderName, lowerDirName)
		if lowerFolderName == lowerDirName {
			return path.Clean(dir + "/" + f.Name()), nil
		}
	}

	return "", errors.New("Could not find Folder")
}
