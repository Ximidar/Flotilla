/*
* @Author: Ximidar
* @Date:   2018-12-04 17:25:32
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-11 15:53:58
 */

package FileSystemTab

import (
	"fmt"

	"github.com/ximidar/gocui"
)

// FileViewInterface will be returned when a FileView object is created.
// The Parent will give it files to display and tell the view when to clear
// This will allow this view to be extremely basic
type FileViewInterface interface {
	AddFileNames(files ...string)
	ClearFiles()
	Layout(g *gocui.Gui) error
}

// FileView is the UI element that will show the file tree of the flotilla system
type FileView struct {
	Name            string
	X, Y, W, H      int
	Structure       []string
	updateStructure bool
	SelectFile      func(file string)
}

// NewFileView will create a new FileView Object
func NewFileView(name string, x int, y int, callback func(file string)) FileViewInterface {
	fv := new(FileView)
	fv.Name = name
	fv.X = x
	fv.Y = y
	fv.SelectFile = callback

	return fv
}

// Layout tells the gocui system how to lay out the ui element
func (fv *FileView) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	fv.W = maxX / 2
	fv.H = maxY - 1
	v, err := g.SetView(fv.Name, fv.X, fv.Y, fv.X+fv.W, fv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Update KeyBindings
		if err := g.SetKeybinding(fv.Name, gocui.KeyEnter, gocui.ModNone, fv.CallSelectFile); err != nil {
			return err
		}
		if err := g.SetKeybinding(fv.Name, gocui.MouseLeft, gocui.ModNone, fv.HighlightFile); err != nil {
			return err
		}
		if err := g.SetKeybinding(fv.Name, gocui.KeyArrowUp, gocui.ModNone, fv.moveSelectUp); err != nil {
			return err
		}
		if err := g.SetKeybinding(fv.Name, gocui.KeyArrowDown, gocui.ModNone, fv.moveSelectDown); err != nil {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		if _, err := g.SetCurrentView(v.Name()); err != nil {
			return err
		}
	}

	return fv.UpdateFileList(g, v)
}

// UpdateFileList will detirmine if an update to the current buffer is needed or not.
func (fv *FileView) UpdateFileList(g *gocui.Gui, v *gocui.View) error {

	if fv.updateStructure {
		v.Clear()
		for _, file := range fv.Structure {
			fmt.Fprintln(v, file)
		}
	}
	return nil
}

// AddFileNames will add file names to the view
func (fv *FileView) AddFileNames(files ...string) {
	fv.Structure = append(fv.Structure, files...)
	fv.updateStructure = true
}

// ClearFiles will clear the view of any filenames
func (fv *FileView) ClearFiles() {
	fv.Structure = []string{}
	fv.updateStructure = true
}

// Functions for moving and selecting files

// HighlightFile will select the file at the cursor position
func (fv *FileView) HighlightFile(g *gocui.Gui, v *gocui.View) error {
	//send selected item
	err := v.SetCursor(v.Cursor())

	if err != nil {
		return err
	}

	_, err = g.SetCurrentView(fv.Name)
	if err != nil {
		return err
	}
	return nil
}

// CallSelectFile will select the file at the current cursor position
func (fv *FileView) CallSelectFile(g *gocui.Gui, v *gocui.View) error {
	err := v.SetCursor(v.Cursor())
	if err != nil {
		return err
	}
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		if err.Error() != "invalid point" {
			return err
		}
		l = ""
		return nil
	}
	fv.SelectFile(l)
	return nil
}

func (fv *FileView) moveSelectUp(g *gocui.Gui, v *gocui.View) error {
	_, cury := v.Cursor()
	orgx, orgy := v.Origin()

	desty := cury - 1

	if desty == orgy-1 {
		desty = orgy
	}

	v.SetCursor(orgx, desty)
	return nil

}

func (fv *FileView) moveSelectDown(g *gocui.Gui, v *gocui.View) error {
	_, cury := v.Cursor()
	orgx, orgy := v.Origin()

	desty := cury + 1

	if desty == (orgy + fv.H) {
		desty = (orgy + fv.H)
	}

	v.SetCursor(orgx, desty)
	return nil

}
