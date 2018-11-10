/*
* @Author: Ximidar
* @Date:   2018-10-29 20:36:43
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-11-10 13:02:48
 */

package FileStreamer_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileManager"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileStreamer"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

type fileStreamerAdapter struct {
	t        *testing.T
	callback func()
}

func (fsa *fileStreamerAdapter) LineReader(line string) {
	fmt.Println("Reading: ", line)
	fsa.callback()

}

func (fsa *fileStreamerAdapter) ProgressUpdate(file *Files.File, currentLine int, readBytes int) {
	fmt.Printf("File: %v\nCurrent Line: %v\nReadBytes: %v\n", file.Name, currentLine, readBytes)
	progress := float64(readBytes) / float64(file.Size) * 100
	fmt.Println("Progress: ", progress)
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
		fs.MonitorFeedback("ok")
	}
	adapter.callback = requestLine

	fs.SelectFile(benchy)

	// Check if the file is actually selected
	if fs.SelectedFile != benchy {
		t.Fatal("Benchy is not the selected file")
	}

	// Stream the file
	fmt.Println("Setup Print Vars")
	fs.SetPlaying(true)
	fs.DonePlaying = false

	fmt.Println("Starting print(simulated)")
	err = fs.StreamFile()
	check_err(t, "TestFileStreamer Streaming the File", err)

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
