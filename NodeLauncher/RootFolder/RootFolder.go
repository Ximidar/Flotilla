package RootFolder

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	// DefaultRootName is the default root folder name
	DefaultRootName = "Flotilla"
	// BinName is the name of the folder for binaries
	BinName = "bin"
)

// RootFolder will generate the flotilla root folder with any and all resources attached
type RootFolder struct {
	// Important Paths
	RootPath string
	BinPath  string
	ArchPath map[string]string
}

// NewRootFolder will return an object for an existing root folder structure
func NewRootFolder(rootPath string) (*RootFolder, error) {
	root := new(RootFolder)
	root.ArchPath = make(map[string]string)
	root.RootPath = path.Clean(rootPath)

	err := root.InitPaths()
	if err != nil {
		return nil, err
	}

	return root, nil
}

// GenerateRootFolder will create a new RootFolder at a specified path. Use this if you need to generate
// all the resources for a root folder.
func GenerateRootFolder(rootPath string) (*RootFolder, error) {
	root := new(RootFolder)
	root.ArchPath = make(map[string]string)
	root.RootPath = path.Clean(rootPath)

	// Generate the different files
	err := root.GeneratePaths()
	if err != nil {
		return nil, err
	}

	// Verify all paths have been successfully created
	err = root.verifyPaths()
	if err != nil {
		return nil, err
	}

	return root, nil
}

// GeneratePaths will generate the folder paths for a flotilla folder
func (root *RootFolder) GeneratePaths() error {
	root.generateDefaultPaths()

	// Create Paths
	pathsToCreate := []string{
		root.RootPath,
		root.BinPath,
		root.ArchPath["amd64"],
		root.ArchPath["arm64"],
		root.ArchPath["arm"],
	}
	var errorsOccured []error
	for _, rawpath := range pathsToCreate {
		fmt.Println("Making Path:", rawpath)
		err := os.MkdirAll(rawpath, 0750)
		if err != nil {
			if err == os.ErrExist {
				fmt.Println("Path Already Exists")
				continue
			}
			fmt.Println("Could not create path", rawpath, "Because:", err.Error())
			errorsOccured = append(errorsOccured, fmt.Errorf("could not create path: %v because: %v", rawpath, err.Error()))
		}
	}

	if len(errorsOccured) != 0 {
		errorString := "The following Errors Occured: "
		for _, err := range errorsOccured {
			errorString = errorString + err.Error() + " "
		}
		return errors.New(errorString)
	}
	return nil
}

// InitPaths will generate the paths for a root folder system. If anything doesn't exist an
// error will be thrown. This function should only work if a root system has been made
func (root *RootFolder) InitPaths() error {

	// Make paths
	root.generateDefaultPaths()

	// Check if paths exist
	err := root.verifyPaths()

	return err
}

func (root *RootFolder) generateDefaultPaths() {

	// Folders
	root.BinPath = root.JoinAndClean(root.RootPath, BinName)
	root.ArchPath["amd64"] = root.JoinAndClean(root.BinPath, "amd64/")
	root.ArchPath["arm64"] = root.JoinAndClean(root.BinPath, "arm64/")
	root.ArchPath["arm"] = root.JoinAndClean(root.BinPath, "arm/")
}

// JoinAndClean will join all paths together and clean the result
func (root *RootFolder) JoinAndClean(paths ...string) string {
	finalPath := ""

	for index, rawpath := range paths {
		if index == 0 {
			finalPath = rawpath
			continue
		}
		finalPath = finalPath + "/" + rawpath
	}

	path.Clean(finalPath)
	return finalPath
}

// verifyPaths will check all of the populated paths to make sure they are there
func (root *RootFolder) verifyPaths() error {

	// gather all paths we want to check
	checkPaths := []string{
		root.RootPath,
		root.BinPath,
		root.ArchPath["amd64"],
		root.ArchPath["arm64"],
		root.ArchPath["arm"],
	}

	for _, rawpath := range checkPaths {
		if !root.pathExists(rawpath) {
			return fmt.Errorf("Path %v Does not exist", rawpath)
		}
	}
	return nil
}

func (root *RootFolder) pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
