/*
* @Author: Ximidar
* @Date:   2018-11-30 15:43:19
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-02 15:13:59
 */

package CommonBlocks

import (
	"errors"
	"fmt"
	"math"

	"github.com/ximidar/gocui"
)

// TabsAdapter is an interface for the Tab Class to tell the Tabs class it is selected
type TabsAdapter interface {
	SelectTab(name string) error
}

// ReportTabAdapter is an interface from a parent class that Tabs will use to report a tab change
type ReportTabAdapter interface {
	UpdateTab(name string)
}

// Tabs will organize all Tab Objects
type Tabs struct {
	X, Y          int
	Name          string
	TabList       []string
	Tabs          []*Tab
	CurrentTab    int
	ReportAdapter ReportTabAdapter
}

// NewTabs will create a new Tabs Object
func NewTabs(x, y int, Name string, reportAdapter ReportTabAdapter) *Tabs {
	tabs := new(Tabs)
	tabs.Name = Name
	tabs.X = x
	tabs.Y = y
	tabs.CurrentTab = 0
	tabs.ReportAdapter = reportAdapter

	return tabs
}

// SelectTab will select the named tab
func (tabs *Tabs) SelectTab(name string) error {
	foundTab := false
	indexTab := 0
	for index, val := range tabs.TabList {
		if val == name {
			indexTab = index
			foundTab = true
		}
	}

	if foundTab {
		for index, val := range tabs.Tabs {
			if index == indexTab {
				val.Selected = true
			} else {
				val.Selected = false
			}

		}
		tabs.ReportAdapter.UpdateTab(name)
		return nil
	}

	return errors.New("Could not find named tab")

}

// Layout Will Layout tabs as needed
func (tabs *Tabs) Layout(g *gocui.Gui) error {
	MaxX, _ := g.Size()
	NumberOfTabs := len(tabs.Tabs)
	spacing := int(math.Abs(math.Min(15, float64(MaxX)/float64(NumberOfTabs))))

	// iterate over tabs and arrange them
	x := tabs.X
	y := tabs.Y
	h := 2

	// Update the layout
	for index := range tabs.Tabs {
		tabs.Tabs[index].X = x
		tabs.Tabs[index].Y = y
		tabs.Tabs[index].W = spacing
		tabs.Tabs[index].H = h
		g.Update(tabs.Tabs[index].Layout)

		x = x + spacing + 1

	}
	return nil
}

// AddTab will add a new tab to the tab bar
func (tabs *Tabs) AddTab(Name string, Label string) {
	tab := NewTab(10, 10, 10, 10, Name, Label, tabs)

	tabs.TabList = append(tabs.TabList, Name)
	tabs.Tabs = append(tabs.Tabs, tab)

	// If this is the first tab added automatically select it
	if len(tabs.Tabs) == 1 {
		tabs.Tabs[0].Selected = true
	}
}

// Tab will be a button that will bring up a specific screen
type Tab struct {
	X, Y, W, H int
	Name       string
	Label      string
	Selected   bool
	Adapter    TabsAdapter
	Handler    func(g *gocui.Gui, v *gocui.View) error
}

// NewTab will Create a new tab
func NewTab(x int, y int, w int, h int, Name, Label string, adapter TabsAdapter) *Tab {
	tab := Tab{X: x, Y: y, W: w, H: h, Name: Name, Label: Label, Adapter: adapter}
	return &tab
}

// Layout will lay the tab in the appropriate position
func (tab *Tab) Layout(g *gocui.Gui) error {
	v, err := g.SetView(tab.Name, tab.X, tab.Y, tab.X+tab.W, tab.Y+tab.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(tab.Name, gocui.KeyEnter, gocui.ModNone, tab.ClickEvent); err != nil {
			return err
		}

		if err := g.SetKeybinding(tab.Name, gocui.MouseLeft, gocui.ModNone, tab.ClickEvent); err != nil {
			return err
		}
		if err := tab.centerLabel(v); err != nil {
			return err
		}
	}

	if tab.Selected {
		v.Highlight = true
	} else {
		v.Highlight = false
	}

	return nil
}

func (tab *Tab) centerLabel(v *gocui.View) error {
	w, _ := v.Size()
	if len(tab.Label) > w {
		return errors.New("label is bigger than the button")
	}

	offsetSize := (w - len(tab.Label)) / 2
	spaceOffset := ""
	for i := 0; i < offsetSize; i++ {
		spaceOffset = spaceOffset + " "
	}
	v.Clear()
	fmt.Fprint(v, fmt.Sprintf("%v%v", spaceOffset, tab.Label))
	return nil
}

// ClickEvent will fire with a click to the object (either click or select). This will fire
// The handler that another object handed it.
func (tab *Tab) ClickEvent(g *gocui.Gui, v *gocui.View) error {

	// Change the color of the tab to show that it is highlighted
	return tab.Adapter.SelectTab(tab.Name)

}
