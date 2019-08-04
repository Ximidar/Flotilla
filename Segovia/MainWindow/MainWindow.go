package MainWindow

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// CreateWidget is a function that will return the MainWindow in this package
func CreateWidget() (*MainWindow, error) {
	mw := new(MainWindow)
	if err := mw.createNotebook(); err != nil {
		return nil, err
	}

	if err := mw.createFiles(); err != nil {
		return nil, err
	}

	if err := mw.createStatus(); err != nil {
		return nil, err
	}

	if err := mw.createUtil(); err != nil {
		return nil, err
	}

	return mw, nil
}

// MainWindow will hold the basic structure for the window
type MainWindow struct {
	Notebook *gtk.Notebook
}

// createWindow will construct the MainWindow
func (mw *MainWindow) createFiles() error {
	Grid, err := gtk.GridNew()
	if err != nil {
		fmt.Println("Could not create Grid")
		return err
	}
	Grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	Grid.SetHExpand(true)
	Grid.SetVExpand(true)

	label, _ := gtk.LabelNew("Segovia Files!")
	Grid.Add(label)

	tabLabel, _ := gtk.LabelNew("Files")
	tabLabel.SetSizeRequest(800/3, 50)

	mw.Notebook.AppendPage(Grid, tabLabel)

	return nil
}

func (mw *MainWindow) createStatus() error {
	Grid, err := gtk.GridNew()
	if err != nil {
		fmt.Println("Could not create Grid")
		return err
	}
	Grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	Grid.SetHExpand(true)
	Grid.SetVExpand(true)

	label, _ := gtk.LabelNew("Segovia Status!")
	Grid.Add(label)

	tabLabel, _ := gtk.LabelNew("Status")
	tabLabel.SetSizeRequest(800/3, 50)

	mw.Notebook.AppendPage(Grid, tabLabel)

	return nil
}

func (mw *MainWindow) createUtil() error {
	Grid, err := gtk.GridNew()
	if err != nil {
		fmt.Println("Could not create Grid")
		return err
	}
	Grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	Grid.SetHExpand(true)
	Grid.SetVExpand(true)

	label, _ := gtk.LabelNew("Segovia Util!")
	Grid.Add(label)

	tabLabel, _ := gtk.LabelNew("Util")
	tabLabel.SetSizeRequest(800/3, 50)

	mw.Notebook.AppendPage(Grid, tabLabel)

	return nil
}

func (mw *MainWindow) createNotebook() error {
	// Create the Main Notebook
	var err error
	mw.Notebook, err = gtk.NotebookNew()
	if err != nil {
		fmt.Println("Could not create Main Notebook")
		return err
	}

	csspro, _ := gtk.CssProviderNew()
	if err := csspro.LoadFromPath("MainWindow/MainWindow.css"); err != nil {
		return err
	}

	context, _ := mw.Notebook.GetStyleContext()
	context.AddProvider(csspro, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	context.AddClass("mainNotebook")
	mw.Notebook.SetHExpand(true)
	mw.Notebook.SetVExpand(true)

	return nil

}
