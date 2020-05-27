package snappy

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
)

// This file automates the packaging of multiarch snaps

// Snappy is a stuct that will build arch snaps
type Snappy struct {
	rf *RootFolder.RootFolder
}

// NewSnappy will take the root folder and package snaps
func NewSnappy(rf *RootFolder.RootFolder) *Snappy {
	snap := new(Snappy)
	snap.rf = rf
	return snap
}

// MakeSnaps will create Snaps for all arches
func (snap *Snappy) MakeSnaps() error {
	var err error
	// Keys of dictionary hold the different Arches
	for arch := range snap.rf.ArchPath {

		// create snap file
		err = snap.CopySnapFile(arch)
		if err != nil {
			return err
		}

		// snap clean
		err = snap.SnapClean()
		if err != nil {
			return err
		}

		// snap create
		err = snap.SnapCreate(arch)
		if err != nil {
			return err
		}

	}

	return nil
}

// SnapCreate will run a snapcraft Pack
func (snap *Snappy) SnapCreate(arch string) error {

	if arch == "arm" {
		arch = "armhf"
	}

	command := "snapcraft"
	args := []string{"snap", "--output", fmt.Sprintf("flotilla_%s.snap", arch), "--target-arch", arch}
	shell := exec.Command(command, args...)
	shell.Env = os.Environ()
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Dir = snap.rf.RootPath

	return shell.Run()
}

// SnapClean will run snapcraft Clean
func (snap *Snappy) SnapClean() error {

	command := "snapcraft"
	args := []string{"clean", "flot-root"}

	shell := exec.Command(command, args...)
	shell.Env = os.Environ()
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Dir = snap.rf.RootPath

	return shell.Run()
}

// CopySnapFile will create a snap file in the snap directory
func (snap *Snappy) CopySnapFile(arch string) error {
	archFile := fmt.Sprintf("Flotilla_linux_%s.zip", arch)
	archPath := snap.rf.JoinAndClean(snap.rf.RootPath, archFile)

	// does archPath Exist?
	if !snap.rf.PathExists(archPath) {
		return fmt.Errorf("Path %s Does not exist", archPath)
	}

	// get snap file text
	snapFileText := snap.generateSnapFile(archFile)

	// copy file text to the snap file
	snapFile := snap.rf.JoinAndClean(snap.rf.SnapPath, "snapcraft.yaml")

	// remove snap file if it already exists
	if snap.rf.PathExists(snapFile) {
		err := os.Remove(snapFile)
		if err != nil {
			return err
		}
	}

	fileo, err := os.Create(snapFile)
	if err != nil {
		return err
	}

	defer fileo.Close()

	_, err = io.Copy(fileo, strings.NewReader(snapFileText))
	if err != nil {
		return err
	}

	return nil
}

// generateSnapFile will give back a yaml string to use with
// the snap
func (snap *Snappy) generateSnapFile(archPath string) string {
	snapFile := fmt.Sprintf(`name: flotilla # you probably want to 'snapcraft register <name>'
base: core18 # the base snap is the execution environment for this snap
version: '0.1' # just for humans, typically '1.2+git' or '1.3.2'
summary: 3D printer server software
description: |
	Flotilla is a 3D printer server designed to be a simple set up
	and run software package. Currently only tested with Marlin.
grade: devel # must be 'stable' to release into candidate/stable channels
confinement: devmode # use 'strict' once you have the right plugs and slots


plugs:
	etc-conf:
		interface: system-files
		read:
			- /etc/flotilla
		write:
			- /etc/flotilla

apps:
	# flotcli
	cli:
		command: Flotilla_CLI
		plugs:
			- network
	
	# nats-server
	nats:
		command: nats-server
		daemon: simple
		plugs:
			- network
			- network-bind
	
	# status
	status:
		command: FlotillaStatus
		daemon: simple
		plugs:
			- network
	
	# commango
	commango:
		command: Commango
		daemon: simple
		plugs:
			- network
			- io-ports-control
			- serial-port
	
	# web
	# flot-web:
	#   command: bin/FlotillaWeb
	#   daemon: simple
	#   plugs:
	#     - network
	#     - network-bind
	#     - etc-conf
	
	# file manager
	file-manager:
		command: Flotilla_File_Manager
		daemon: simple
		plugs:
			- network
			- etc-conf
			- io-ports-control


parts:
	flot-root:
		plugin: dump
		source: %s
		source-type: zip

	`, archPath)
	snapFile = strings.ReplaceAll(snapFile, "\t", "  ")
	return snapFile
}
