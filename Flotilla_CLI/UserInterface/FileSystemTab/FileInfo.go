/*
* @Author: Ximidar
* @Date:   2018-12-06 13:57:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 16:39:04
 */

package FileSystemTab

import (
	"fmt"
	"strings"
	"time"

	readable "github.com/dustin/go-humanize"
	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/jroimartin/gocui"
)

// FileInfo will show the current info about the selected file
type FileInfo struct {
	X, Y, W, H   int
	Name         string
	CurrentFile  *FS.File
	UpdateInfo   bool
	RootFilePath string
}

// NewFileInfo will create a FileInfo instance
func NewFileInfo(y int, name string) *FileInfo {
	fi := new(FileInfo)
	fi.X = 0
	fi.W = 10
	fi.H = 2
	fi.Y = y
	fi.Name = name

	return fi
}

// Layout will tell gocui how to layout FileInfo
func (fi *FileInfo) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	fi.W = maxX - 1
	fi.H = maxY - 1
	fi.X = (maxX / 2) + 1
	v, err := g.SetView(fi.Name, fi.X, fi.Y, fi.W, fi.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// Update keybindings
	}

	if fi.UpdateInfo {
		fi.UpdateFileInfo(g, v)
	}
	return nil
}

// DeliverFile will be called by an outside source to deliver a file to be shown
func (fi *FileInfo) DeliverFile(file *FS.File) {
	fi.CurrentFile = file
	fi.UpdateInfo = true
}

// UpdateFileInfo will update the info on the file info screen
func (fi *FileInfo) UpdateFileInfo(g *gocui.Gui, v *gocui.View) {
	v.Clear()

	fmt.Fprintln(v, fi.CurrentFile.GetName())
	fmt.Fprintln(v, readable.Bytes(fi.CurrentFile.GetSize()))
	unix := fi.CurrentFile.GetUnixTime()
	fileTime := time.Unix(unix, 0)
	fmt.Fprintln(v, fileTime.Format("Mon Jan _2 2006"))

	maxX, _ := v.Size()
	path := strings.Replace(fi.CurrentFile.Path, fi.RootFilePath, "root", 1)
	if len(path) > maxX {
		length := maxX - 4
		startmess := len(path) - length
		mess := "..." + path[startmess:]
		fmt.Fprintln(v, mess)
	} else {
		fmt.Fprintln(v, path)
	}

}
