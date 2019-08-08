/*
* @Author: Ximidar
* @Date:   2018-06-16 16:39:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 16:57:47
 */

package UserInterface

import (
	"fmt"

	"github.com/ximidar/Flotilla/Flotilla_CLI/UserInterface/CommTab"
	"github.com/ximidar/Flotilla/Flotilla_CLI/UserInterface/CommonBlocks"
	"github.com/ximidar/Flotilla/Flotilla_CLI/UserInterface/FileSystemTab"
	"github.com/jroimartin/gocui"
)

// CliGui is an object that will instantiate the ui
type CliGui struct {
	TabList        *CommonBlocks.Tabs
	CommTab        *commtab.CommTab
	FileTab        *FileSystemTab.FileSystemTab
	CurrentTabName string
	RootGUI        *gocui.Gui
}

// NewCliGui is the constructor for CliGui
func NewCliGui() (*CliGui, error) {
	cli := new(CliGui)
	cli.TabList = CommonBlocks.NewTabs(0, 0, "Tabs", cli)
	cli.CurrentTabName = "CommTab"

	return cli, nil
}

// ScreenInit will initialize the screen for the gui
func (gui *CliGui) ScreenInit() (err error) {

	// Make Tabs
	gui.TabList.AddTab("CommTab", "Comm")
	gui.TabList.AddTab("FileTab", "Files")

	err = gui.setupCommTab()
	if err != nil {
		fmt.Println("Could not Create Comm Tab", err)
		return err
	}
	err = gui.setupFileTab()
	if err != nil {
		fmt.Println("Could not Create File tab", err)
		return err
	}
	gui.RootGUI, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer gui.RootGUI.Close()

	// Set GUI options
	gui.RootGUI.Cursor = true
	gui.RootGUI.Mouse = true
	gui.RootGUI.Highlight = true
	gui.RootGUI.SelFgColor = gocui.ColorGreen
	err = gui.RootGUI.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.quit)
	if err != nil {
		return err
	}

	gui.RootGUI.SetManagerFunc(gui.Layout)

	if err := gui.RootGUI.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
		//return err
	}
	return

}

// UpdateTab will set the currently selected tab, then it will reset the manager func which will
// get rid of all views and keybindings.
func (gui *CliGui) UpdateTab(name string) {
	gui.CurrentTabName = name

	// Reset Gui for new tab
	gui.RootGUI.SetManagerFunc(gui.Layout)

}

// CheckSize makes sure the size of the screen is big enough to accomodate the tool
func (gui *CliGui) CheckSize(x, y int) bool {
	if x > 88 || y > 20 {
		return true
	}
	return false
}

// Layout is a function for Gocui to help layout the screen
func (gui *CliGui) Layout(g *gocui.Gui) error {
	x, y := g.Size()
	// if !gui.CheckSize(x, y) {
	// 	return nil
	// }

	var managers []func(*gocui.Gui) error

	// Add the tab to every view
	managers = append(managers, gui.TabList.Layout)

	//Update TabContents based on Selected Tab
	switch gui.CurrentTabName {
	case "CommTab":
		managers = append(managers, gui.CommTab.Layout)
	case "FileTab":
		managers = append(managers, gui.FileTab.Layout)
	}

	// Update all layouts
	for _, layout := range managers {
		g.Update(layout)
	}

	// reset root gui options and bindings
	_, err := g.SetView("FlotillaUI", x+1, y+1, x+2, y+2)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.Cursor = true
		g.Mouse = true
		g.Highlight = true
		g.SelFgColor = gocui.ColorGreen
		if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.quit); err != nil {
			return err
		}

	}

	return nil
}

// CommTabHandler will controll pulling up the CommTab Contents.
func (gui *CliGui) CommTabHandler(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (gui *CliGui) setupCommTab() error {
	var err error
	gui.CommTab, err = commtab.NewCommTab(0, 3, gui.RootGUI)
	if err != nil {
		fmt.Println("Could not set up Comm Tab", err)
		return err
	}
	gui.CommTab.Name = "CommContents"

	return nil
}

func (gui *CliGui) setupFileTab() error {
	var err error
	gui.FileTab, err = FileSystemTab.NewFileSystemTab("FileContents", 0, 3)
	return err
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	view, err := g.SetViewOnTop(name)
	if err != nil {
		view.SetCursor(view.Origin())
	}

	return view, err
}

func (gui *CliGui) quit(g *gocui.Gui, v *gocui.View) error {
	// TODO add a function here that will tell all running tabs to quit
	return gocui.ErrQuit
}
