/*
* @Author: Ximidar
* @Date:   2018-10-29 20:36:43
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-05-07 19:53:30
 */

package FileStreamer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	"github.com/Ximidar/Flotilla/Flotilla_File_Manager/FileManager"
)

type fileStreamerAdapter struct {
	t           *testing.T
	callback    func()
	DoneChannel chan bool
}

func (fsa *fileStreamerAdapter) NotifyDone() {
	fmt.Println("Done!")
	fsa.DoneChannel <- true
}

func (fsa *fileStreamerAdapter) SendError(mess string) {
	fmt.Println("Error!", mess)
}

func (fsa *fileStreamerAdapter) LineReader(line *CommRelayStructures.Line) {
	fmt.Println("Reading: ", line.Line)
	go fsa.callback()

}

func (fsa *fileStreamerAdapter) ProgressUpdate(file *FS.File, currentLine uint64, readBytes uint64) {
	fmt.Printf("File: %v\nCurrent Line: %v\nReadBytes: %v\n", file.Name, currentLine, readBytes)
	progress := float64(readBytes) / float64(file.Size) * 100
	fmt.Println("Progress: ", progress)
}

func (fsa *fileStreamerAdapter) SendStatus(status string) {
	fmt.Println(status)
}

// TestSetup will test the basic setup of the filestreamer
func TestSetup(t *testing.T) {
	fmt.Println("Testing Setup")
	fsa := new(fileStreamerAdapter)
	_, err := NewFileStreamer(fsa)

	if err != nil {
		t.Fatal(err)
	}
}

var TestLocation = "/tmp/testing/FileStreamer"

// TestGetLines will test getting amounts of lines
func TestGetLinesFunctionality(t *testing.T) {

	fm, fs := MakeTestingInstance(t)

	// get the file to select
	benchy, err := GetBenchy(fm)
	CommonTestTools.CheckErr(t, "TestFileStreamer getting benchy", err)
	fmt.Println(benchy.Path)

	// Select the benchy
	fs.SelectFile(benchy)

	lengthsCheck := []int{1, 2, 1000, 1000, 1000}

	for _, length := range lengthsCheck {
		fmt.Println("Getting Length:", length)
		lines, err := fs.GetLines(length)
		CommonTestTools.CheckErr(t, "Couldn't get any lines", err)

		if len(lines) != length {
			t.Fatal("Lines are not populated", lines)
		}

	}
}

func TestStreamingFile(t *testing.T) {
	t.Skip("Test takes too long to complete. Run when doing full test")
	fm, fs := MakeTestingInstance(t)

	// get the file to select
	benchy, err := GetBenchy(fm)
	CommonTestTools.CheckErr(t, "TestFileStreamer getting benchy", err)
	fmt.Println(benchy.Path)

	// Select the benchy
	fs.SelectFile(benchy)

	// Keep Selecting lines until there are none left
	for {
		lines, err := fs.GetLines(1)
		if err != nil {
			if err == io.EOF {
				fmt.Println("End of File!")
				break
			}
			CommonTestTools.CheckErr(t, "Could not get lines", err)
		}
		fmt.Print(lines[0].GetLine())
	}

}

// TestGetLines will test getting amounts of lines
func TestGetSpecificLine(t *testing.T) {

	fm, fs := MakeTestingInstance(t)

	// get the file to select
	benchy, err := GetBenchy(fm)
	CommonTestTools.CheckErr(t, "TestFileStreamer getting benchy", err)
	fmt.Println(benchy.Path)

	// Select the benchy
	fs.SelectFile(benchy)

	// line 59835 == "G1 F1800 X115.432 Y89.749 E2537.18181"
	linesToGet := []uint64{0, 5, 20, 5000, 59835}
	expected := []string{
		"M190 S60\n",                              //Line 0
		"G1 F900 X72.127 Y89.873 E0.02024\n",      //Line 5
		"G1 X99.602 Y87.332 E3.47956\n",           // Line 20
		"G1 F2400 X101.522 Y107.179 E448.44093\n", // Line 5000
		"G1 F1800 X115.432 Y89.749 E2537.18181\n", // Line 59835
	}

	// Get good lines
	for i, lineNum := range linesToGet {
		fmt.Println("Getting line ", lineNum)
		line, err := fs.GetLineAt(lineNum)
		CommonTestTools.CheckErr(t, "Could not get specific line", err)
		fmt.Print(line.GetLine())

		if line.GetLine() != expected[i] {
			t.Fatal("Lines do not match")
		}
	}

	// Get Bad lines
	linesToGet = []uint64{10000000000}
	for _, lineNum := range linesToGet {
		fmt.Println("Getting line ", lineNum)
		_, err := fs.GetLineAt(lineNum)
		if err == nil {
			t.Fatal("Got A bad line and did not pop an error")
		}

	}

}

func MakeTestingInstance(t *testing.T) (*FileManager.FileManager, *FileStreamer) {
	fm, err := FileManager.NewFileManager(TestLocation)
	CommonTestTools.CheckErr(t, "TestFileStreamer Making File Manager", err)

	// Create the File Streamer object and select the file
	adapter := fileStreamerAdapter{t: t}
	fs, err := NewFileStreamer(&adapter)
	CommonTestTools.CheckErr(t, "Couldn't make File Streamer", err)

	return fm, fs
}

func GetBenchy(fm *FileManager.FileManager) (*FS.File, error) {
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")

	if _, err := os.Stat(benchyOrgFile); !os.IsNotExist(err) {
		err := fm.AddFile(benchyOrgFile, destPath)
		if err != nil {
			fmt.Println("Could not Add Benchy", err)
			return nil, err
		}
		file, err := fm.GetFileByPath("3D_Benchy.gcode")
		if err != nil {
			fmt.Println("Could not get File", err)
			structure, _ := fm.GetJSONStructure()
			fmt.Println(string(structure))
			return nil, err
		}
		return file, nil
	}

	return nil, fmt.Errorf("Could not get file %v", ".3D_Benchy.gcode")

}
