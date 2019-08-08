package ContentBox

import (
	"errors"

	"github.com/gdamore/tcell"
)

// ContentBox will create a box that can print text
type ContentBox struct {
	x, y, w, h int
	screen     tcell.Screen
	lines      []string
	Name       string
}

// NewContentBox will create a new content box
func NewContentBox(s tcell.Screen, name string, x, y, w, h int) (*ContentBox, error) {
	cb := new(ContentBox)
	cb.screen = s
	cb.Name = name
	cb.x = x
	cb.y = y
	cb.w = w
	cb.h = h

	if err := cb.drawBox(); err != nil {
		return nil, err
	}
	return cb, nil
}

func (cb *ContentBox) drawBox() error {

	// check box size
	w, h := cb.screen.Size()
	if cb.x > w || cb.x < 0 {
		return errors.New("out of bounds")
	} else if cb.y > h || cb.y < 0 {
		return errors.New("out of bounds")
	}

	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

	for col := cb.x; col <= cb.x+cb.w; col++ {
		cb.screen.SetContent(col, cb.y, tcell.RuneHLine, nil, style)
		cb.screen.SetContent(col, cb.y+cb.y, tcell.RuneHLine, nil, style)
	}
	for row := cb.y + 1; row < cb.y+cb.y; row++ {
		cb.screen.SetContent(cb.x, row, tcell.RuneVLine, nil, style)
		cb.screen.SetContent(cb.x+cb.w, row, tcell.RuneVLine, nil, style)
	}
	if cb.y != cb.y+cb.y && cb.x != cb.x+cb.w {
		// Only add corners if we need to
		cb.screen.SetContent(cb.x, cb.y, tcell.RuneULCorner, nil, style)
		cb.screen.SetContent(cb.x+cb.w, cb.y, tcell.RuneURCorner, nil, style)
		cb.screen.SetContent(cb.x, cb.y+cb.y, tcell.RuneLLCorner, nil, style)
		cb.screen.SetContent(cb.x+cb.w, cb.y+cb.y, tcell.RuneLRCorner, nil, style)
	}
	for row := cb.y + 1; row < cb.y+cb.y; row++ {
		for col := cb.x + 1; col < cb.x+cb.w; col++ {
			cb.screen.SetContent(col, row, rune('b'), nil, style)
		}
	}

	return nil
}
