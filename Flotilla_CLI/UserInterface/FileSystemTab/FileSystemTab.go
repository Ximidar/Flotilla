/*
* @Author: Ximidar
* @Date:   2018-12-02 13:26:45
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 13:21:37
 */

package FileSystemTab

import (
	"fmt"

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/Ximidar/Flotilla/Flotilla_CLI/UserInterface/CommonBlocks"
	"github.com/jroimartin/gocui"
)

const (
	// FileViewName : Name for the File View
	FileViewName = "FileView"

	// FileInfoName : Name for the File Info View
	FileInfoName = "FileInfo"

	// SelectButtonName : Name of the button to select files
	SelectButtonName = "SelectButton"

	// PathBarName : Name for the view in charge of the current path
	PathBarName = "PathBar"
)

// FileSystemTab is an object for displaying the FileSystem
type FileSystemTab struct {
	X, Y      int
	Name      string
	FI        *FlotillaInterface.FlotillaInterface
	Structure *FS.File

	// Views
	FileView     FileViewInterface
	FileInfo     *FileInfo
	SelectButton *CommonBlocks.Button

	// File Manipulation
	CurrentDirectory *FolderNode
}

// NewFileSystemTab will construct a new Filesystem object
func NewFileSystemTab(name string, x int, y int) (*FileSystemTab, error) {

	fs := new(FileSystemTab)
	fs.Name = name
	fs.X = x
	fs.Y = y

	// Set up the flotillainterface
	var err error
	fs.FI, err = FlotillaInterface.NewFlotillaInterface()
	if err != nil {
		return nil, err
	}
	fs.Structure, err = fs.FI.GetFileStructure()
	if err != nil {
		return nil, err
	}
	fs.initNode()

	// Set up the fileview
	fs.FileView = NewFileView(FileViewName, fs.X, fs.Y, fs.SelectFile)
	fs.UpdateFileList()

	// Set up FileInfo
	fs.FileInfo = NewFileInfo(fs.Y+3, FileInfoName)
	fs.FileInfo.RootFilePath = fs.Structure.GetPath()

	// Set up SelectButton
	fs.SelectButton = CommonBlocks.NewButton(SelectButtonName, fs.FileInfo.X+1, fs.FileInfo.Y, 30, "Select and Play", fs.SelectAndPlay)

	return fs, nil

}

// Layout will tell gocui how to layout this widget
func (fs *FileSystemTab) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	_, err := g.SetView(fs.Name, maxX+1, maxY+1, maxX+2, maxY+2)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Update keybindings
	}
	g.Update(fs.FileView.Layout)
	g.Update(fs.FileInfo.Layout)
	if fs.FileInfo.CurrentFile == nil {
		fs.SelectButton.Selectable = false
		fs.SelectButton.Update = true
	} else {
		fs.SelectButton.Selectable = true
		fs.SelectButton.Update = true
	}
	fs.SelectButton.UpdateSize(fs.FileInfo.X, fs.Y, 30)
	g.Update(fs.SelectButton.Layout)

	return nil
}

// SelectAndPlay will select the current file and start playing it
func (fs *FileSystemTab) SelectAndPlay(g *gocui.Gui, v *gocui.View) error {
	file := fs.FileInfo.CurrentFile
	err := fs.FI.SelectAndPlayFile(file)
	if err != nil {
		mess := fmt.Sprintf("Cannot Play: %v", err.Error())
		label := CommonBlocks.NewLabel("PlayFileMessage", mess, 10, 10, len(mess)+2, 2)
		g.Update(label.Layout)

		return nil
	}
	mess := fmt.Sprintf("Playing File: %v", file.Name)
	label := CommonBlocks.NewLabel("PlayFileMessage", mess, 10, 10, len(mess)+2, 2)
	g.Update(label.Layout)
	return nil
}

// UpdateFileList will update the files in the view
func (fs *FileSystemTab) UpdateFileList() {
	fs.FileView.ClearFiles()
	fs.FileView.AddFileNames(fs.CurrentDirectory.GetFileList()...)
}

// SelectFile will select the a file or it will navigate the filesystem
func (fs *FileSystemTab) SelectFile(file string) {
	if file == ".." {
		// Return to previous node
		fs.CurrentDirectory = fs.CurrentDirectory.PreviousNode
		fs.UpdateFileList()
		return
	}

	fileInfo, err := fs.CurrentDirectory.ReturnFileByName(file)
	if err != nil {
		return
	}

	if fileInfo.IsDir {
		tempCD, err := fs.CurrentDirectory.MoveToFolder(file)
		if err != nil {
			return
		}
		fs.CurrentDirectory = tempCD
		fs.UpdateFileList()
		return
	}

	fs.FileInfo.DeliverFile(fileInfo)
}

// initNode will apply the structure to the folder node and instantiate the root folder
func (fs *FileSystemTab) initNode() {
	var contents []*FS.File
	if fs.Structure.GetContents() != nil {
		contents = fs.Structure.GetContents()
	}

	fs.CurrentDirectory = &FolderNode{PreviousNode: nil, Contents: contents, Info: fs.Structure}
}
