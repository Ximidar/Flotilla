/*
* @Author: Ximidar
* @Date:   2018-10-29 20:36:43
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-28 14:29:57
 */

package FileStreamer_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileManager"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileStreamer"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

type fileStreamerAdapter struct {
	t           *testing.T
	callback    func()
	DoneChannel chan bool
}

func (fsa *fileStreamerAdapter) LineReader(line *CommRelayStructures.Line) {
	fmt.Println("Reading: ", line.Line)
	go fsa.callback()

}

func (fsa *fileStreamerAdapter) ProgressUpdate(file *Files.File, currentLine uint64, readBytes uint64) {
	fmt.Printf("File: %v\nCurrent Line: %v\nReadBytes: %v\n", file.Name, currentLine, readBytes)
	progress := float64(readBytes) / float64(file.Size) * 100
	fmt.Println("Progress: ", progress)
}

func (fsa *fileStreamerAdapter) SendStatus(status string) {
	fmt.Println(status)
	fsa.DoneChannel <- true
}

// TestSetup will test the basic setup of the filestreamer
func TestSetup(t *testing.T) {
	fmt.Println("Testing Setup")
	fsa := new(fileStreamerAdapter)
	_, err := FileStreamer.NewFileStreamer(fsa)

	if err != nil {
		t.Fatal(err)
	}
}

func check_err(t *testing.T, mess string, err error) {
	if err != nil {
		t.Fatalf("Failed Check from %v, Error: %v", mess, err)
	}
}

// TestFileStreamer will attempt to select a file and stream it
func TestFileStreamer(t *testing.T) {
	fmt.Println("Testing the Streaming Capabilities")
	fm, err := FileManager.NewFileManager()

	check_err(t, "TestFileStreamer Making File Manager", err)

	// get the file to select
	benchy, err := GetBenchy(fm)
	check_err(t, "TestFileStreamer getting benchy", err)

	fmt.Println(benchy.Path)
	fs := new(FileStreamer.FileStreamer)

	// Create the File Streamer object and select the file
	adapter := fileStreamerAdapter{t: t}
	fs, err = FileStreamer.NewFileStreamer(&adapter)
	check_err(t, "TestFileStreamer Making the adapter and file streamer", err)
	// Create a function for the adapter to request the next line
	requestLine := func() {
		fs.MonitorFeedback()
	}
	adapter.callback = requestLine

	fs.SelectFile(benchy)

	// Check if the file is actually selected
	if fs.SelectedFile != benchy {
		t.Fatal("Benchy is not the selected file")
	}

	// callbacks for pause and play
	pauseResumeCallback := func() {
		<-time.After(2 * time.Second)
		fmt.Println("\nPAUSE")
		fs.Pause()
		<-time.After(2 * time.Second)
		fmt.Println("\nResume!")
		fs.Resume()
	}

	// Stream the file
	fmt.Println("Setup Print Vars")
	fs.DonePlaying = false
	fs.SetPlaying(false)
	adapter.DoneChannel = make(chan bool, 10)

	fmt.Println("Starting print(simulated)")
	go pauseResumeCallback()
	fs.Play()
	<-adapter.DoneChannel
	//check_err(t, "TestFileStreamer Streaming the File", err)

	if !fs.DonePlaying {
		check_err(t, "TestFileStreamer Failed to flip done bool", errors.New("Failed to flip done bool"))
	}

}

func GetBenchy(fm *FileManager.FileManager) (*Files.File, error) {
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")

	if _, err := os.Stat(benchyOrgFile); !os.IsNotExist(err) {
		err := fm.AddFile(benchyOrgFile, destPath)
		if err != nil {
			return nil, err
		}
		file, err := fm.GetFileByPath("3D_Benchy.gcode")
		if err != nil {
			structure, _ := fm.GetJSONStructure()
			fmt.Println(string(structure))
			return nil, err
		}
		return file, nil
	}

	return nil, fmt.Errorf("Could not get file %v", ".3D_Benchy.gcode")

}
