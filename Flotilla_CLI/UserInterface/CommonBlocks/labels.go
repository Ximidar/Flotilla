/*
* @Author: Ximidar
* @Date:   2018-12-04 15:47:49
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-11 16:06:57
 */

package CommonBlocks

import (
	"fmt"
	"time"

	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/gocui"
)

// Label is an object for displaying a temporary message on the screen
type Label struct {
	X, Y, W, H int
	Name       string
	Message    string
	FI         *FlotillaInterface.FlotillaInterface
}

// NewLabel will construct a new Filesystem object
func NewLabel(name string, message string, x int, y int, w int, h int) *Label {

	fs := new(Label)
	fs.Name = name
	fs.Message = message
	fs.X = x
	fs.Y = y
	fs.W = w
	fs.H = h

	return fs

}

// Layout will tell gocui how to layout this widget
func (fs *Label) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	middleX := (maxX / 2) - (len(fs.Message) / 2)
	middleY := maxY / 2

	v, err := g.SetView(fs.Name, middleX, middleY, middleX+fs.W, middleY+fs.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(v.Name(), gocui.MouseLeft, gocui.ModNone, fs.Destroy); err != nil {
			return err
		}
		if err := g.SetKeybinding(v.Name(), gocui.KeyEnter, gocui.ModNone, fs.Destroy); err != nil {
			return err
		}
		fmt.Fprintln(v, fs.Message)
		g.SetViewOnTop(v.Name())
		g.SetCurrentView(v.Name())

		// Check for focus will check if this view is focused on or not.
		// If it is not focused on then it will be destroyed
		checkFocusFunc := func() {
			for {
				select {
				case <-time.After(200 * time.Millisecond):

				}

				if g.CurrentView() != v {
					fs.Destroy(g, v)
					return
				}
			}
		}
		go checkFocusFunc()
	}

	return nil
}

// Destroy will destroy the view and the keybindings
func (fs *Label) Destroy(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(fs.Name)
	g.DeleteKeybindings(fs.Name)
	return nil
}
