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
	"io"
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

	// Arches contains all the available arches
	Arches = []string{ARM32, ARM64, AMD64}
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
	pf.Muster["Commango"] = pf.RootFolder.CoreBins
	pf.Muster["Flotilla_File_Manager"] = pf.RootFolder.CoreBins
	pf.Muster["FlotillaStatus"] = pf.RootFolder.CoreBins

	// Extra Binaries
	pf.Muster["Flotilla_CLI"] = pf.RootFolder.ExtraBins
	pf.Muster["NodeLauncher"] = pf.RootFolder.ExtraBins

}

func (pf *PopulateFolder) evaluateFlotillaPath() error {
	// Evaluate Flotilla Path
	FlotillaRelPath := "/src/github.com/ximidar/Flotilla/"

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

// MakePackage will run make on the package specified
func (pf *PopulateFolder) MakePackage(packageName string) error {

	packagePath, err := pf.FindFolder(packageName, pf.FlotillaPath)
	if err != nil {
		return err
	}

	// #nosec
	// Start a command from the package path
	cmd := exec.Command("make")

	// Add the directory to run it from
	cmd.Dir = packagePath

	// Output the output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (pf *PopulateFolder) simpleCopy(src, dest string) error {

	fmt.Printf("Copying %v to %v\n", src, dest)
	buf := make([]byte, 5*1000)
	// #nosec
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}

	// #nosec
	// Change Mode of destination file so it is executable to a normal user
	os.Chmod(dest, 0755)
	return nil
}

// MakeAll will make all packages
func (pf *PopulateFolder) MakeAll() error {
	for name := range pf.Muster {
		err := pf.MakePackage(name)
		if err != nil {
			return err
		}
	}
	return nil
}

// PackageAllBinaries will send all built binaries to their destinations
func (pf *PopulateFolder) PackageAllBinaries() error {
	for name, dest := range pf.Muster {
		// Find bin folder
		binFolder, err := pf.FindFlotillaBinFolder(name)
		if err != nil {
			return err
		}

		// extract binary
		bin, name, err := pf.FindBinForArch(binFolder)
		if err != nil {
			return err
		}

		// Copy to dest
		prefix := pf.Arch + "_"
		base := strings.Replace(name, prefix, "", 1)
		fulldest := path.Clean(dest + "/" + base)
		err = pf.simpleCopy(bin, fulldest)
		if err != nil {
			return err
		}
	}
	return nil
}

// PackageStaticFiles will send all static files to their destinations
func (pf *PopulateFolder) PackageStaticFiles() error {
	staticFiles := []string{pf.RootFolder.ConfigFile, pf.RootFolder.StartServerFile}

	for _, sf := range staticFiles {
		basename := path.Base(sf)
		src, err := pf.FindStaticFile(basename)
		if err != nil {
			return err
		}
		pf.simpleCopy(src, sf)

	}
	return nil
}

// GenerateCerts will generate the TLS config files
func (pf *PopulateFolder) GenerateCerts() error {

	// Get the static file for the config file
	tlsCNFPath, err := pf.FindStaticFile("flotillaTLS.cnf")
	if err != nil {
		fmt.Println("Could not find config file for TLS")
		return err
	}

	// Make a monitor for a command we want to run
	command := "openssl"
	args := []string{"req",
		"-new",
		"-x509",
		"-newkey",
		"rsa:2048",
		"-config",
		tlsCNFPath,
		"-keyout",
		"flotilla.key",
		"-out",
		"flotilla.pem",
		"-outform",
		"PEM"}

	cmd := exec.Command(command, args...)
	cmd.Dir = pf.RootFolder.TLS

	err = cmd.Run()
	if err != nil {
		out, _ := cmd.Output()
		fmt.Println("Could not run openssl command", string(out), err)
		return err
	}

	return nil

}

// Populate will populate the RootFolder with items
func (pf *PopulateFolder) Populate(skipMakeAll bool) error {

	// Evaluate Muster and create a list of package paths

	// Make All
	if !skipMakeAll {
		err := pf.MakeAll()
		if err != nil {
			return err
		}
	}

	// Package all binaries
	err := pf.PackageAllBinaries()
	if err != nil {
		return err
	}

	// Package all static files
	err = pf.PackageStaticFiles()
	if err != nil {
		return err
	}

	// Generate Certificates
	err = pf.GenerateCerts()
	if err != nil {
		return err
	}

	return nil
}
