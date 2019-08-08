/*
* @Author: Ximidar
* @Date:   2018-11-29 13:14:25
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-26 18:22:25
 */

// Package commtab is the user interface for connecting and monitoring
// the serial line for flotilla
package commtab

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/jroimartin/gocui"
	"github.com/nats-io/go-nats"
	CS "github.com/ximidar/Flotilla/DataStructures/CommStructures"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/Flotilla/Flotilla_CLI/UserInterface/CommonBlocks"
)

const (
	// ConnectionView Name for connection info view
	ConnectionView = "connection_info"

	// MonitorView Name for monitor view
	MonitorView = "monitor_view"

	// SendView Name for send view
	SendView = "send_view"

	// BaudButton Name for Baud Button view
	BaudButton = "baud_button"

	// PortButton Name for Port Button View
	PortButton = "port_button"

	// ConnectButton Name for Connect Button view
	ConnectButton = "connect_button"

	// DisconnectButton Name for Disconnect Button View
	DisconnectButton = "disconnect_button"

	// InfoView Name for info view
	InfoView = "info_view"
)

// CommTab Creates a Command Line GUI
type CommTab struct {
	readerActive bool
	RootGUI      *gocui.Gui
	Name         string

	Monitor MonitorInterface
	x, y    int

	// widgets
	connectionInfo   *ConnectionInfo
	sendBar          *SendBar
	portButton       *CommonBlocks.ExplodeButton
	baudButton       *CommonBlocks.ExplodeButton
	connectButton    *CommonBlocks.Button
	disconnectButton *CommonBlocks.Button
	CycleViews       []string
	SelectedView     int

	port string
	baud int

	FlotillaInterface *FlotillaInterface.FlotillaInterface

	mux sync.Mutex
}

// NewCommTab will Create a CommTab object
func NewCommTab(x, y int, g *gocui.Gui) (*CommTab, error) {
	gui := new(CommTab)

	// CycleViews is an array of strings to cycle the selected view
	gui.CycleViews = append(gui.CycleViews, SendView, PortButton, BaudButton, ConnectButton, DisconnectButton)
	gui.SelectedView = 9999 // Select the send bar

	var err error
	gui.FlotillaInterface, err = FlotillaInterface.NewFlotillaInterface()
	if err != nil {
		log.Println("Could not setup flotilla interface")
		return nil, err
	}

	gui.RootGUI = g
	gui.x = x
	gui.y = y
	gui.readerActive = false

	err = gui.setupBlocks()

	if err != nil {
		log.Println("Could not set up blocks")
		return nil, err
	}
	gui.port = ""
	gui.baud = -1

	gui.readerActive = true
	gui.CommRelay()

	return gui, nil
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

func (gui *CommTab) nextView(g *gocui.Gui, v *gocui.View) (err error) {

	gui.SelectedView++

	if gui.SelectedView >= len(gui.CycleViews) {
		gui.SelectedView = 0
	}
	_, err = setCurrentViewOnTop(g, gui.CycleViews[gui.SelectedView])
	g.Cursor = true

	return err
}

func (gui *CommTab) setupBlocks() (err error) {
	gui.Monitor = NewMonitor(MonitorView, 31+gui.x, 0+gui.y)
	gui.sendBar = NewSendBar(SendView, 31+gui.x, -3, gui.writeToComm)
	gui.connectionInfo, err = NewConnectionInfo(gui.FlotillaInterface, ConnectionView, gui.x, gui.y, 30, 7)
	gui.portButton = CommonBlocks.NewExplodeButton(PortButton, 0+gui.x, 8+gui.y, 14, "Port Select", gui.getPorts, gui.portSelect)
	gui.baudButton = CommonBlocks.NewExplodeButton(BaudButton, 15+gui.x, 8+gui.y, 15, "Baud Select", gui.getBauds, gui.baudSelect)
	gui.connectButton = CommonBlocks.NewButton(ConnectButton, 0+gui.x, 11+gui.y, 30, "Connect", gui.connectComm)
	gui.disconnectButton = CommonBlocks.NewButton(DisconnectButton, 0+gui.x, 14+gui.y, 30, "Disconnect", gui.disconnectComm)
	return
}

// Layout is CommTab's gocui Layout Function
func (gui *CommTab) Layout(g *gocui.Gui) error {
	gui.RootGUI = g
	g.Update(gui.Monitor.Layout)
	g.Update(gui.sendBar.Layout)
	g.Update(gui.connectionInfo.Layout)
	g.Update(gui.portButton.Layout)
	g.Update(gui.baudButton.Layout)
	g.Update(gui.connectButton.Layout)
	g.Update(gui.disconnectButton.Layout)

	// Update keybindings
	maxX, MaxY := g.Size()
	_, err := g.SetView(gui.Name, maxX+1, MaxY+1, maxX+2, MaxY+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			fmt.Println(err)
			return err
		}
		err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, gui.nextView)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gui *CommTab) writeToComm(mess string) {
	gui.FlotillaInterface.CommWrite(mess)
}

func (gui *CommTab) getBauds() []string {
	return []string{"250000", "230400", "115200", "57600", "38400", "19200", "9600"}
}

func (gui *CommTab) baudSelect(selection string) {
	gui.Monitor.Write(gui.RootGUI, fmt.Sprintf("Selection %v ", selection))
	if tempBaud, err := strconv.Atoi(selection); err == nil {
		gui.baud = int(tempBaud)
	} else {
		gui.Monitor.Write(gui.RootGUI, "default to 115200")
		gui.baud = 115200
	}
}

func (gui *CommTab) connectComm(g *gocui.Gui, v *gocui.View) error {
	gui.Monitor.Write(g, "connect!")
	gui.FlotillaInterface.CommSetConnectionOptions(gui.port, gui.baud)
	gui.FlotillaInterface.CommConnect()
	return nil
}

func (gui *CommTab) disconnectComm(g *gocui.Gui, v *gocui.View) error {
	gui.Monitor.Write(g, "disconnect!")
	gui.FlotillaInterface.CommDisconnect()
	return nil
}

func (gui *CommTab) getPorts() (portlist []string) {
	ports, err := gui.FlotillaInterface.CommGetAvailablePorts()

	if err != nil {
		portlist = append(portlist, err.Error())
		return
	}

	for _, p := range ports.GetPorts() {
		portlist = append(portlist, p.GetAddress())
	}

	return
}

func (gui *CommTab) portSelect(selection string) {
	gui.Monitor.Write(gui.RootGUI, fmt.Sprintf("Selection %v ", selection))
	gui.port = selection
}

// CommRelay will subscribes functions to incoming data from Nats
func (gui *CommTab) CommRelay() {

	gui.mux.Lock()
	defer gui.mux.Unlock()
	_, err := gui.FlotillaInterface.NC.Subscribe(CS.ReadLine, gui.CommReadSub)
	if err != nil {
		gui.Monitor.Write(gui.RootGUI, fmt.Sprintf("Error %v", err.Error()))
	}
	_, err = gui.FlotillaInterface.NC.Subscribe(CS.WriteLine, gui.CommWriteSub)

	if err != nil {
		gui.Monitor.Write(gui.RootGUI, fmt.Sprintf("Error %v", err.Error()))
	}

}

// CommReadSub will reveive Comm Messages from the Nats Server
func (gui *CommTab) CommReadSub(msg *nats.Msg) {

	read, err := gui.deconstructCommMessage(msg.Data)
	if err != nil {
		return
	}
	data := read.GetMessage()
	data = fmt.Sprintf("\u001b[46mRECV:\u001b[0m \n%v", data)
	data = strings.Replace(data, "\n", "\n\t", -1)
	data = strings.Replace(data, "echo:", "", -1)
	gui.mux.Lock()
	gui.Monitor.Write(gui.RootGUI, data)
	gui.mux.Unlock()
}

// CommWriteSub will Revieve updates from the Nats server on Written Messages
func (gui *CommTab) CommWriteSub(msg *nats.Msg) {
	wrote, err := gui.deconstructCommMessage(msg.Data)
	if err != nil {
		return
	}
	data := wrote.GetMessage()
	data = strings.Replace(data, "\n", "", -1)
	data = fmt.Sprintf("\u001b[44mSENT: %v\u001b[0m", data)
	gui.mux.Lock()
	gui.Monitor.Write(gui.RootGUI, data)
	gui.mux.Unlock()
}

func (gui *CommTab) deconstructCommMessage(cmb []byte) (cm *CS.CommMessage, err error) {
	cm = new(CS.CommMessage)
	err = proto.Unmarshal(cmb, cm)
	return
}
