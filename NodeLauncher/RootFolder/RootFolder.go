/*
* @Author: Ximidar
* @Date:   2019-02-04 15:27:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 15:42:18
 */

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

	// GCODEFolderName is the name for the gcode file locations
	GCODEFolderName = "GCODE"

	// BinName is the name of the bin folder
	BinName = "bin"

	// BinCoreSubfolder is the name of the core binary folder
	BinCoreSubfolder = "CoreFlotilla"

	// BinExtraSubfolder is the name of the extra binary folder
	BinExtraSubfolder = "Extras"

	// ConfigName is the name of the config file
	ConfigName = "Config.yaml"

	// StartServerName is the name of the start script
	StartServerName = "StartFlotilla.sh"

	// WorkspaceName is the name of the workspace
	WorkspaceName = "Workspace"

	// WSLogs is a subfolder of workspace where logs would go
	WSLogs = "Logs"

	// WSCerts hold the certificates for TLS
	WSCerts = ".certs"
)

// RootFolder will generate the flotilla root folder with any and all resources attached
type RootFolder struct {
	// Important Paths
	RootPath        string
	GCODEPath       string
	BinPath         string
	CoreBins        string
	ExtraBins       string
	Workspace       string
	Logs            string
	TLS             string
	ConfigFile      string
	StartServerFile string
}

// NewRootFolder will return an object for an existing root folder structure
func NewRootFolder(rootPath string) (*RootFolder, error) {
	root := new(RootFolder)
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
		root.GCODEPath,
		root.BinPath,
		root.CoreBins,
		root.ExtraBins,
		root.Workspace,
		root.Logs,
		root.TLS,
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

	// Create Files
	filesToCreate := []string{
		root.ConfigFile,
		root.StartServerFile,
	}

	for _, rawpath := range filesToCreate {
		fmt.Println("Making file:", rawpath)
		file, err := os.Create(rawpath)
		if err != nil {
			errorsOccured = append(errorsOccured, err)
		}
		file.Close()
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
	root.GCODEPath = root.JoinAndClean(root.RootPath, GCODEFolderName)
	root.BinPath = root.JoinAndClean(root.RootPath, BinName)
	root.CoreBins = root.JoinAndClean(root.BinPath, BinCoreSubfolder)
	root.ExtraBins = root.JoinAndClean(root.BinPath, BinExtraSubfolder)
	root.Workspace = root.JoinAndClean(root.RootPath, WorkspaceName)
	root.Logs = root.JoinAndClean(root.Workspace, WSLogs)
	root.TLS = root.JoinAndClean(root.Workspace, WSCerts)

	// Files
	root.ConfigFile = root.JoinAndClean(root.RootPath, ConfigName)
	root.StartServerFile = root.JoinAndClean(root.RootPath, StartServerName)
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
		root.GCODEPath,
		root.BinPath,
		root.CoreBins,
		root.ExtraBins,
		root.Workspace,
		root.Logs,
		root.ConfigFile,
		root.StartServerFile,
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
