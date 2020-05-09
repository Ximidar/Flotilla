/*
* @Author: Ximidar
* @Date:   2019-02-07 15:48:51
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 17:17:35
 */

package RootFolder

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	// ARM32 Denotes an arch
	ARM32 = "arm"
	// ARM64 Denotes an arch
	ARM64 = "arm64"
	// AMD64 Denotes an arch
	AMD64 = "amd64"
	// ALLARCH Denotes to build all arches
	ALLARCH = "all"

	// Arches contains all the available arches
	Arches = []string{ARM32, ARM64, AMD64}

	BuildPak = []string{
		"Commango",
		"Flotilla_File_Manager",
		"FlotillaStatus",
		"Flotilla_CLI",
		"FlotillaWeb",
	}

	// Tags for Version, Author, and Date
	VERSION     = "0.0.1"
	DATE        = `date '+%d %b %y at %H:%M:%S %p'`
	AUTHOR      = "Matt Pedler"
	COMMIT_HASH = "git rev-parse HEAD"

	VERSION_PATH = "github.com/Ximidar/Flotilla/CommonTools/versioning"
)

// PopulateFolder will be used in conjunction with the root folder to populate the executables
type PopulateFolder struct {
	RootFolder *RootFolder

	// populate from local GOPATH
	GOPATH       string
	FlotillaPath string
	Muster       map[string]string
	Arch         string

	// TODO create the required systems to populate from the internet
}

// NewPopulateFolder will create a new instance of PopulateFolder
func NewPopulateFolder(rf *RootFolder, architecture string) (*PopulateFolder, error) {

	pf := new(PopulateFolder)
	pf.RootFolder = rf
	pf.GOPATH = os.Getenv("GOPATH")
	if pf.GOPATH == "" {
		return nil, errors.New("GOPATH does not exist")
	}
	err := pf.evaluateFlotillaPath()
	if err != nil {
		return nil, err
	}
	pf.Arch = architecture
	if !pf.evaluateArch() {
		return nil, fmt.Errorf("Architecture must be one of these: %v", strings.Join(Arches, " | "))
	}

	pf.evaluateMuster()

	return pf, nil
}

// This function only exists to populate the Folder names for flotilla and their ultimate destinations
func (pf *PopulateFolder) evaluateMuster() {
	pf.Muster = make(map[string]string)

	// Core Binaries
	pf.Muster["Commango"] = pf.RootFolder.BinPath
	pf.Muster["Flotilla_File_Manager"] = pf.RootFolder.BinPath
	pf.Muster["FlotillaStatus"] = pf.RootFolder.BinPath
	pf.Muster["Flotilla_CLI"] = pf.RootFolder.BinPath
	pf.Muster["Flotilla_Web"] = pf.RootFolder.BinPath

}

func (pf *PopulateFolder) evaluateFlotillaPath() error {
	// Evaluate Flotilla Path
	FlotillaRelPath := "/src/github.com/Ximidar/Flotilla/"

	pf.FlotillaPath = path.Clean(pf.GOPATH + FlotillaRelPath)

	if _, err := os.Stat(pf.FlotillaPath); os.IsNotExist(err) {
		return fmt.Errorf("Flotilla Folder Does not Exist: %v", pf.FlotillaPath)
	}
	return nil
}

func (pf *PopulateFolder) evaluateArch() bool {
	for _, arch := range Arches {
		if pf.Arch == arch {
			return true
		}
	}
	return false
}

// FindFlotillaBinFolder will attempt to find the bin of a supplied executable name
// This is for finding executables for flotilla.
func (pf *PopulateFolder) FindFlotillaBinFolder(progname string) (string, error) {

	// Find folder for progname
	FlotillaFolder, err := pf.FindFolder(progname, pf.FlotillaPath)
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

// FindBinForArch will return the binary for the specified Arch
func (pf *PopulateFolder) FindBinForArch(binpath string) (string, string, error) {
	files, err := ioutil.ReadDir(binpath)

	if err != nil {
		return "", "", err
	}

	prefix := pf.Arch + "_"

	// Find Binary
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			binpath := path.Clean(binpath + "/" + file.Name())
			fmt.Println(prefix, "Found bin", binpath)
			return binpath, file.Name(), nil
		}
	}

	return "", "", errors.New("could not find binary")
}

// FindFolder will attempt to find the folder at a specified directory and return it's
// path
func (pf *PopulateFolder) FindFolder(folderName string, dir string) (string, error) {

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

	return "", fmt.Errorf("Could not find Folder: %v in Dir: %v", folderName, dir)
}

// FindFile will find a file in a specific Directory
func (pf *PopulateFolder) FindFile(name, dir string) (string, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.Name() == name {
			return path.Clean(dir + "/" + file.Name()), nil
		}
	}

	return "", fmt.Errorf("cannot find file %v in path %v", name, dir)

}

// FindStaticFile will attempt to find a static file in the NodeLauncher dir
func (pf *PopulateFolder) FindStaticFile(name string) (string, error) {
	NodePath, err := pf.FindFolder("NodeLauncher", pf.FlotillaPath)
	if err != nil {
		return "", err
	}

	// If it didn't error out previously this directory probably exists
	staticPath := path.Clean(NodePath + "/StaticFiles/")

	// find file
	return pf.FindFile(name, staticPath)

}

// BuildPackages will build the different flotilla packages
func (pf *PopulateFolder) BuildPackages() error {
	for _, flotpack := range BuildPak {

		// Get the package path
		packagePath, err := pf.FindFolder(flotpack, pf.FlotillaPath)
		if err != nil {
			errMsg := fmt.Sprintf("Package %s has failed to build because %v", flotpack, err)
			panic(errMsg)
		}

		entry, _ := pf.FindFile(flotpack+".go", packagePath)
		entry = path.Base(entry)

		for _, arch := range Arches {
			ldflags := []string{"`-X",
				fmt.Sprintf("'%s.Version=%s'", VERSION_PATH, VERSION),
				"-X",
				fmt.Sprintf("'%s.CompiledBy=%s'", VERSION_PATH, AUTHOR),
				"-X",
				fmt.Sprintf("'%s.CompiledDate=$(%s)'", VERSION_PATH, DATE),
				"-X",
				fmt.Sprintf("'%s.CommitHash=$(%s)'`", VERSION_PATH, COMMIT_HASH),
			}
			arguments := []string{
				"build",
				"-ldflags",
				strings.Join(ldflags, " "),
				"-o",
				fmt.Sprintf("%s%s", pf.RootFolder.ArchPath[arch], flotpack),
				entry,
			}

			// LOG build vars
			fmt.Println("BUILDING ", flotpack)
			fmt.Println("PACKAGE LOCATION", packagePath)
			fmt.Println("ENTRY FILE ", entry)
			fmt.Println("ARCH ", arch)

			// build executable
			logCom := "go " + strings.Join(arguments, " ")
			fmt.Println("RUNNING ", logCom)
			fmt.Println("")

			shell := exec.Command("go", arguments...)
			shell.Dir = packagePath
			shell.Env = append(os.Environ(),
				"GOOS=linux",
				fmt.Sprintf("GOARCH=%s", arch))
			shell.Stdout = os.Stdout
			shell.Stderr = os.Stderr
			err := shell.Run()
			if err != nil {
				return err
			}

		}

	}
	return nil
}

// Populate will populate the RootFolder with items
func (pf *PopulateFolder) Populate(skipMakeAll bool) error {

	// Evaluate Muster and create a list of package paths

	// Make All
	if !skipMakeAll {
		err := pf.BuildPackages()
		if err != nil {
			return err
		}
	}

	return nil
}
