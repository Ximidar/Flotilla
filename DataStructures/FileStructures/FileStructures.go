/*
* @Author: Ximidar
* @Date:   2018-10-10 06:36:00
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-17 12:15:51
 */

package FileStructures

import "fmt"

const (
	// Name is the name of the program to call
	Name = "FILE_MANAGER."

	// File Structure controls

	// SelectFile will Select a file
	SelectFile = Name + "SELECT_FILE"
	// GetFileStructure will get the JSON representation of the current file structure
	GetFileStructure = Name + "GET_FILE_STRUCTURE"
	// AddFile will add a file to the file structure
	AddFile = Name + "ADD_FILE"
	// MoveFile will move a file from one place to another
	MoveFile = Name + "MOVE_FILE"
	// DeleteFile will delete a file
	DeleteFile = Name + "DELETE_FILE"

	// Print Controls

	// IsPrinting will query if the printer is printing or not
	IsPrinting = Name + "IS_PRINTING"
	// IsPaused will query if the printer is paused
	IsPaused = Name + "IS_PAUSED"
	// TogglePause will toggle the pause state
	TogglePause = Name + "TOGGLE_PAUSE"
	// StartPrint will start a print
	StartPrint = Name + "START_PRINT"
	// CancelPrint will cancel a print
	CancelPrint = Name + "CANCEL_PRINT"

	// Publishers

	// UpdateFS will be called any time there is an update to the file system
	UpdateFS = Name + "UPDATE_FS"
	// FileProgress will be called every time the file progresses (0 - 100%)
	FileProgress = Name + "FILE_PROGRESS"
)

// FileAction is an object to convey what action you want
// the File system to take
type FileAction struct {
	Action string `json:"action"`
	Path   string `json:"path"`
}

// NewFileAction will construct a FileAction object
func NewFileAction(action string, path string) (*FileAction, error) {
	fa := new(FileAction)
	if err := fa.SetAction(action); err != nil {
		return nil, err
	}
	fa.Path = path
	return fa, nil
}

// SetAction will make sure a real action is taking place.
func (fa *FileAction) SetAction(action string) error {
	var ok bool
	switch action {
	case SelectFile, GetFileStructure, AddFile, MoveFile, DeleteFile:
		ok = true
	default:
		ok = false
	}

	if !ok {
		return fmt.Errorf("File Action not available: %v", action)
	}
	fa.Action = action
	return nil
}
