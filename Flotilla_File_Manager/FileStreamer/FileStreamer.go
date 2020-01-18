/*
* @Author: Ximidar
* @Date:   2018-10-17 17:14:20
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-05-07 22:47:26
 */

package FileStreamer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
)

// Adapter is an interface for other classes to be read files
// AKA This will read files to the NATS interface
type Adapter interface {
	ProgressUpdate(file *FS.File, currentLine uint64, readBytes uint64)
}

var (
	//ErrFileNotSelected will be returned if there isn't a file to do any transforms on
	ErrFileNotSelected = errors.New("no file selected")

	// ErrIgnoreLine means that the line is a comment or a line to be ignored
	ErrIgnoreLine = errors.New("ignore line")
)

// FileStreamer Takes a file and streams it to the NATS Comm Object
type FileStreamer struct {
	SelectedFile *FS.File
	OpenFile     *os.File
	LineNumber   uint64 // This is the raw line number read in the file
	Adapter      Adapter

	currentBytes uint64
}

// NewFileStreamer will construct a FileStreamer object
func NewFileStreamer(adapter Adapter) (*FileStreamer, error) {
	fs := new(FileStreamer)
	fs.Adapter = adapter
	fs.LineNumber = 0
	fs.currentBytes = 0

	return fs, nil
}

// SelectFile will select a file
func (fs *FileStreamer) SelectFile(file *FS.File) error {
	fmt.Printf("Selecting File: %v\n", file.Name)
	fs.SelectedFile = file

	return nil
}

func (fs *FileStreamer) isFileSelected() bool {
	return !(fs.SelectedFile == nil)
}

// Cancel will reset the current instance
func (fs *FileStreamer) Cancel() {
	if fs.isFileSelected() {
		fs.OpenFile.Close()
	}

	fs.OpenFile = nil
	fs.SelectedFile = nil
	fs.LineNumber = 0
	fs.currentBytes = 0

}

// GetLines will return an amount of requested lines where the file last left off
func (fs *FileStreamer) GetLines(amount int) ([]*CommRelayStructures.Line, error) {
	if !fs.isFileSelected() {
		return nil, errors.New("File Not Selected")
	}

	freader, err := fs.GetReaderForCurrentFile(int64(fs.currentBytes))
	if err != nil {
		fmt.Println("Could not get the file reader", err)
		return nil, err
	}

	defer fs.OpenFile.Close()

	// request lines
	receivedLines := make([]*CommRelayStructures.Line, 0)

	for i := 0; i < amount; i++ {
		line, err := fs.GetNextLine(freader)
		if err != nil {
			return receivedLines, err
		}
		receivedLines = append(receivedLines, line)

	}

	return receivedLines, nil
}

// GetNextLine will be used to get the next gcode line in the file. It will ignore any commented lines
func (fs *FileStreamer) GetNextLine(reader *bufio.Reader) (*CommRelayStructures.Line, error) {

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		// Update Progress of reading file
		fs.updateProgress(line)

		err = fs.checkLine(line)
		if err != nil {
			continue // This will only error out if the line is a commented line
		}

		fline, err := fs.formatLine(line)
		if err != nil {
			// Something went horribly wrong
			fmt.Println("Could not make formatted line", err)
			return nil, err
		}

		return fline, nil
	}

}

// GetLineAt will return the specified line number
func (fs *FileStreamer) GetLineAt(lineNum uint64) (*CommRelayStructures.Line, error) {
	if !fs.isFileSelected() {
		return nil, errors.New("File Not Selected")
	}

	// Get the file from the start
	freader, err := fs.GetReaderForCurrentFile(0)
	if err != nil {
		fmt.Println("Could not get the file reader", err)
		return nil, err
	}

	defer fs.OpenFile.Close()
	counter := uint64(0)

	// Scan through file until we find the line number
	for {
		line, err := freader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read from file")
			return nil, err
		}

		err = fs.checkLine(line)
		if err != nil {
			continue // This will only error out if the line is a commented line
		}

		if counter == lineNum {
			fline, err := CommRelayStructures.NewLine(line, lineNum, true)
			if err != nil {
				fmt.Println("ERROR Could not send line because someone couldn't build a stable struct with three variables...")
				return nil, err
			}
			return fline, nil
		}

		counter++

	}
}

// GetReaderForCurrentFile will return a reader for the currently selected file or an error
func (fs *FileStreamer) GetReaderForCurrentFile(openLoc int64) (*bufio.Reader, error) {
	// Check if the selected file is there
	if !fs.checkSelectedFile() {
		fmt.Println("No File Selected!")
		return nil, ErrFileNotSelected
	}
	// open file

	if fs.OpenFile != nil {
		// We don't care about errors here because it might already be closed.
		fs.OpenFile.Close()
	}

	var err error
	fs.OpenFile, err = fs.openSelectedFile()
	if err != nil {
		return nil, err
	}

	// Seek File to most recent
	readSeeker := io.ReadSeeker(fs.OpenFile)
	location, err := readSeeker.Seek(openLoc, 0)
	if err != nil {
		fmt.Println("Could not seek to ", fs.currentBytes,
			"Current location is at:", location)
		return nil, err
	}

	// Make reader
	reader := bufio.NewReader(readSeeker)

	return reader, nil
}

func (fs *FileStreamer) checkSelectedFile() bool {
	if fs.SelectedFile == nil {
		return false
	}

	if _, err := os.Stat(fs.SelectedFile.Path); os.IsNotExist(err) {
		return false
	}

	return true
}

func (fs *FileStreamer) openSelectedFile() (*os.File, error) {
	f, err := os.Open(fs.SelectedFile.Path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// checkLine will filter out any comments not to be sent
func (fs *FileStreamer) checkLine(line string) error {

	// Check if the line is a comment or a blank line and clean it up
	trimline := strings.Replace(line, " ", "", -1)
	trimline = strings.Replace(trimline, "\n", "", -1)
	if len(trimline) == 0 {
		return ErrIgnoreLine
	} else if trimline[:1] == ";" {
		return ErrIgnoreLine
	}

	// Check for inline comments and take them out
	if strings.Contains(line, ";") {
		index := strings.Index(line, ";")
		line = line[:index]
	}
	return nil
}

// this function will read the line to the adapter. This is also where we should modify the line for "printing mode" ie G1 X10 turns into n350 G1 X10!9
func (fs *FileStreamer) formatLine(line string) (*CommRelayStructures.Line, error) {
	fline, err := CommRelayStructures.NewLine(line, fs.LineNumber, true)
	fs.LineNumber++
	if err != nil {
		fmt.Println("ERROR Could not send line because someone couldn't build a stable struct with three variables...")
		return nil, err
	}

	return fline, nil

}

// updateProgress will update the progression through the file
func (fs *FileStreamer) updateProgress(line string) {
	fs.currentBytes += uint64(len(line))
	fs.Adapter.ProgressUpdate(fs.SelectedFile, fs.LineNumber, fs.currentBytes)
}
