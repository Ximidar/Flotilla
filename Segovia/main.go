package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/ximidar/Flotilla/Segovia/MainWindow"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Segovia")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	log.Println("Making MainWindow")
	mw, err := MainWindow.CreateWidget()
	if err != nil {
		log.Println("Could not create MainWindow")
		panic(err)
	}

	// Add the label to the window.
	win.Add(mw.Notebook)

	// Set the default window size.
	win.SetDefaultSize(800, 480)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
