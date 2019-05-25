/*
* @Author: Ximidar
* @Date:   2018-08-25 21:59:56
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-04 16:19:46
 */
package commtab

import (
	"fmt"

	"github.com/ximidar/gocui"
)

// SendBar is a gui object for sending strings
type SendBar struct {
	name    string
	x, y    int
	handler func(message string)
}

// NewSendBar will construct a SendBar Object
func NewSendBar(name string, x, y int, handler func(message string)) *SendBar {
	return &SendBar{name: name,
		x:       x,
		y:       y,
		handler: handler}
}

// Layout is SendBar's Gocui Layout Function
func (w *SendBar) Layout(g *gocui.Gui) error {
	maxX, MaxY := g.Size()
	v, err := g.SetView(w.name, w.x, MaxY+w.y, maxX-1, (MaxY+w.y)+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			fmt.Println(err)
			return err
		}
		v.Title = "Send"
		v.Editable = true
		err = g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.sendViewClear)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *SendBar) sendViewClear(g *gocui.Gui, v *gocui.View) error {
	contents := v.Buffer()
	v.Clear()
	v.SetCursor(v.Origin())

	// send the contents somewhere
	w.handler(contents)

	return nil
}
