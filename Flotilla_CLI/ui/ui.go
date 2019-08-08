package ui

import (
	"fmt"
	"os"
	"sort"

	"github.com/gdamore/tcell"
)

// MainScreen is the root element for the ui
type MainScreen struct {
	Screen   tcell.Screen
	QuitKeys []string
}

// NewMainScreen will create a new root element
func NewMainScreen() (*MainScreen, error) {
	ms := new(MainScreen)
	var err error
	ms.Screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}
	if err = ms.Screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}

	ms.Screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	ms.Screen.Clear()

	ms.QuitKeys = []string{}

	return ms, nil
}

// Run will keep the screen alive
func (ms *MainScreen) Run() {

	for {
		ev := ms.Screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				fmt.Println(ev.Key())
				return
			case tcell.KeyCtrlL:
				ms.Screen.Sync()
			}
			if ms.quit(ev.Rune()) {
				return
			}
		case *tcell.EventResize:
			ms.Screen.Sync()
		}
	}
}

func (ms *MainScreen) quit(key rune) bool {
	skey := string(key)
	i := sort.Search(len(ms.QuitKeys),
		func(i int) bool {
			return ms.QuitKeys[i] >= skey
		})
	if i < len(ms.QuitKeys) && ms.QuitKeys[i] == skey {
		return true
	}
	return false
}

// AddQuitKey will add a key that will be used for quitting
func (ms *MainScreen) AddQuitKey(key ...string) {
	ms.QuitKeys = append(ms.QuitKeys, key...)
	sort.Strings(ms.QuitKeys)
}
