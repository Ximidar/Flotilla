/*
* @Author: Ximidar
* @Date:   2018-08-25 21:58:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-11 15:21:09
 */

package CommonBlocks

import (
	"errors"
	"fmt"

	"github.com/ximidar/gocui"
)

// Button is a object to create a button
type Button struct {
	name       string
	x, y       int
	w          int
	label      string
	Selectable bool
	Update     bool
	handler    func(g *gocui.Gui, v *gocui.View) error
}

// NewButton will connstruct a Button
func NewButton(name string, x, y, w int, label string, handler func(g *gocui.Gui, v *gocui.View) error) *Button {
	return &Button{name: name,
		x:          x,
		y:          y,
		w:          w,
		label:      label,
		handler:    handler,
		Selectable: true,
	}
}

// Layout is Button's Gocui Layour Function
func (b *Button) Layout(g *gocui.Gui) error {
	v, err := g.SetView(b.name, b.x, b.y, b.x+b.w, b.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err

		}
		if err := g.SetKeybinding(b.name, gocui.KeyEnter, gocui.ModNone, b.clickButton); err != nil {
			return err
		}
		if err := g.SetKeybinding(b.name, gocui.MouseLeft, gocui.ModNone, b.clickButton); err != nil {
			return err
		}
		if err := b.centerLabel(v); err != nil {
			return err
		}
	}

	if b.Update {
		if err := b.centerLabel(v); err != nil {
			return err
		}
		b.Update = false
	}

	if !b.Selectable {
		v.FgColor = gocui.ColorBlack
		v.BgColor = gocui.ColorBlack
	} else {
		v.FgColor = gocui.ColorDefault
		v.BgColor = gocui.ColorDefault
	}
	return nil
}

// UpdateSize will update the button's size
func (b *Button) UpdateSize(x, y, w int) {
	if x == b.x && y == b.y && w == b.w {
		return
	}
	b.x = x
	b.y = y
	b.w = w
	b.Update = true

}

func (b *Button) clickButton(g *gocui.Gui, v *gocui.View) error {
	if b.Selectable {
		return b.handler(g, v)
	}
	return nil
}

func (b *Button) fitText(text string, v *gocui.View) string {
	maxX, _ := v.Size()
	if len(text) > maxX {
		length := maxX - 4
		startmess := len(text) - length
		mess := "..." + text[startmess:]
		return mess
	}
	return text
}

func (b *Button) centerLabel(v *gocui.View) error {
	w, _ := v.Size()
	b.label = b.fitText(b.label, v)

	offsetSize := (w - len(b.label)) / 2
	spaceOffset := ""
	for i := 0; i < offsetSize; i++ {
		spaceOffset = spaceOffset + " "
	}
	fmt.Fprint(v, fmt.Sprintf("%v%v", spaceOffset, b.label))
	return nil
}

/////////////////////////////////////////////////////////////////////
// Explode Button
/////////////////////////////////////////////////////////////////////

// ExplodeButton will create a Button that will have a selector that appears
// in the middle of the screen
type ExplodeButton struct {
	name  string
	x, y  int
	w     int
	label string

	getBody        func() []string
	selectCallback func(selection string)
}

// NewExplodeButton will Construct an Explode Button Object
func NewExplodeButton(name string, x, y, w int, label string, getBody func() []string, selectCallback func(selection string)) *ExplodeButton {
	return &ExplodeButton{name: name,
		x:              x,
		y:              y,
		w:              w,
		label:          label,
		getBody:        getBody,
		selectCallback: selectCallback,
	}
}

// Layout is ExplodeButton's Gocui Layout Function
func (b *ExplodeButton) Layout(g *gocui.Gui) error {
	v, err := g.SetView(b.name, b.x, b.y, b.x+b.w, b.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(b.name, gocui.KeyEnter, gocui.ModNone, b.explode); err != nil {
			return err
		}
		if err := g.SetKeybinding(b.name, gocui.MouseLeft, gocui.ModNone, b.explode); err != nil {
			return err
		}
		if err := b.centerLabel(v); err != nil {
			return err
		}
	}

	return nil
}

func (b *ExplodeButton) centerLabel(v *gocui.View) error {
	w, _ := v.Size()
	if len(b.label) > w {
		return errors.New("label is bigger than the button")
	}

	offsetSize := (w - len(b.label)) / 2
	spaceOffset := ""
	for i := 0; i < offsetSize; i++ {
		spaceOffset = spaceOffset + " "
	}
	fmt.Fprint(v, fmt.Sprintf("%v%v", spaceOffset, b.label))
	return nil
}

func (b *ExplodeButton) explode(g *gocui.Gui, v *gocui.View) error {
	body := b.getBody()
	midx, midy := g.Size()
	midx = midx / 2
	midy = midy / 2
	name := fmt.Sprintf("%s_explode", b.name)
	explode := NewExplode(name, midx, midy, body, b.selectCallback)
	g.Update(explode.Layout)
	return nil
}

// Explode will put a Gocui selector at the middle of the screen
type Explode struct {
	name           string
	x, y           int
	w, h           int
	body           []string
	selectCallback func(selection string)
}

// NewExplode will create an Explode Object
func NewExplode(name string, x, y int, body []string, selectCallback func(selection string)) *Explode {
	w := 0
	for _, l := range body {
		if len(l) > w {
			w = len(l)
		}
	}
	h := len(body) + 1
	w = w + 1

	return &Explode{name: name, x: x, y: y, w: w, h: h, body: body, selectCallback: selectCallback}
}

// Layout is Explode's Gocui Layout function
func (w *Explode) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.InputEsc = true
		if err := g.SetKeybinding(w.name, gocui.KeyEsc, gocui.ModNone, w.destroy); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.selectAndDestroy); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.MouseLeft, gocui.ModNone, w.selectAndDestroy); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyArrowUp, gocui.ModNone, w.moveSelectUp); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyArrowDown, gocui.ModNone, w.moveSelectDown); err != nil {
			return err
		}
		for _, line := range w.body {
			fmt.Fprintln(v, line)
		}

		// Make it selected and highlight the first choice
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		if _, err := g.SetCurrentView(v.Name()); err != nil {
			return err
		}

	}
	return nil
}

func (w *Explode) selectAndDestroy(g *gocui.Gui, v *gocui.View) error {
	//send selected item
	v.SetCursor(v.Cursor())
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		l = ""
		panic(err)
	}
	w.selectCallback(l)
	w.destroy(g, v)
	return nil
}

func (w *Explode) destroy(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView(w.name)
	g.DeleteKeybindings(w.name)
	return nil
}

func (w *Explode) moveSelectUp(g *gocui.Gui, v *gocui.View) error {
	_, cury := v.Cursor()
	orgx, orgy := v.Origin()

	desty := cury - 1

	if desty == orgy-1 {
		desty = orgy
	}

	v.SetCursor(orgx, desty)
	return nil

}

func (w *Explode) moveSelectDown(g *gocui.Gui, v *gocui.View) error {
	_, cury := v.Cursor()
	orgx, orgy := v.Origin()

	desty := cury + 1

	if desty == (orgy + w.h) {
		desty = (orgy + w.h)
	}

	v.SetCursor(orgx, desty)
	return nil

}
