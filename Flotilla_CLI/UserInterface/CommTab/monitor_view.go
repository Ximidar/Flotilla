/*
* @Author: Ximidar
* @Date:   2018-08-25 22:00:52
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-12 16:19:01
 */

package commtab

import (
	"fmt"
	"os"
	"strings"

	"github.com/ximidar/gocui"
)

// MonitorInterface is an adapter for the constructor of Monitor to return. It allows control
// over the UI element
type MonitorInterface interface {
	Write(g *gocui.Gui, mess string)
	Read(g *gocui.Gui) string
	Clear(g *gocui.Gui)
	Layout(g *gocui.Gui) error
}

// Monitor is an object that will display communication happening over
// the serial line on flotilla
type Monitor struct {
	name       string
	x, y       int
	buffersize int
}

// NewMonitor will create a new monitor object
func NewMonitor(name string, x, y int) MonitorInterface {
	return &Monitor{name: name, x: x, y: y, buffersize: 0}
}

// Layout is used by gocui to organize the ui element on the screen
func (w *Monitor) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView(w.name, w.x, w.y, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			fmt.Println(err)
			return err
		}
		v.Title, err = os.Hostname()
		if err != nil {
			return err
		}
		v.Autoscroll = true
		v.Wrap = true
	}
	return nil
}

// Write will write a string to the UI Element
func (w *Monitor) Write(g *gocui.Gui, mess string) {
	w.buffersize += len(mess)
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(w.name)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			// if the view is unkown dont write to it. TODO add unkown writes to buffer
			return nil
		}

		fmt.Fprintln(v, StringCleaner(mess))
		w.CheckBuffer(g, v)
		return err

	})

}

// StringCleaner will chop off the trailing \n
func StringCleaner(mess string) string {
	if strings.HasSuffix(mess, "\n") {
		return mess[:len(mess)-1]
	}
	return mess

}

// Clear will Clear the monitor of any accrued lines
func (w *Monitor) Clear(g *gocui.Gui) {
	v, err := g.View(w.name)
	if err != nil {
		return
	}
	v.Clear()
	v.SetCursor(v.Origin())
}

// Read will return the buffer of the monitor
func (w *Monitor) Read(g *gocui.Gui) string {
	v, err := g.View(w.name)
	if err != nil {
		return ""
	}
	return v.Buffer()
}

// CheckBuffer will check the current size of the buffer and cut it as necessary
func (w *Monitor) CheckBuffer(g *gocui.Gui, v *gocui.View) {
	if w.buffersize > 1000000 {
		v.Clear()
		w.buffersize = 0
	}
}
