/*
* @Author: Ximidar
* @Date:   2018-08-25 22:00:30
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-26 17:29:39
 */

package commtab

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	nats "github.com/nats-io/go-nats"
	CS "github.com/ximidar/Flotilla/DataStructures/CommStructures"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
)

// ConnectionInfo will monitor the NATS system for the current connection status
type ConnectionInfo struct {
	X, Y, W, H    int
	Name          string
	FI            *FlotillaInterface.FlotillaInterface
	gui           *gocui.Gui
	currentStatus *CS.CommStatus
}

// NewConnectionInfo is a constructor and it will return a ConnectionInfo object
func NewConnectionInfo(FI *FlotillaInterface.FlotillaInterface, name string, x int, y int, w int, h int) (*ConnectionInfo, error) {
	ci := new(ConnectionInfo)
	ci.FI = FI
	ci.Name = name
	ci.X = x
	ci.Y = y
	ci.W = w
	ci.H = h
	err := ci.SetupNatsSubs()
	if err != nil {
		log.Println("Could not setup subs")
		return nil, err
	}
	return ci, nil
}

// SetupNatsSubs will setup the subscribtions that Connection info will be connected to
func (ci *ConnectionInfo) SetupNatsSubs() error {
	_, err := ci.FI.NC.Subscribe(CS.StatusUpdate, ci.UpdateStatus)
	if err != nil {
		return err
	}

	status, err := ci.FI.CommGetStatus()
	if err != nil {
		return err
	}
	ci.currentStatus = status

	return nil
}

// UpdateStatus is a function that will be called automatically by the nats system when an
// update to the Comm status occurs
func (ci *ConnectionInfo) UpdateStatus(msg *nats.Msg) {

	// alter view with current connection data
	newStatus, err := ci.FI.DeconstructStatus(msg.Data)
	if err != nil {
		log.Println("Could not get status")
	}
	ci.currentStatus = newStatus
}

// Layout will tell gocui how to layout this widget
func (ci *ConnectionInfo) Layout(g *gocui.Gui) error {
	// grab the most recent pointer to the gui.
	ci.gui = g

	// setup screen 30 width, 7 height
	v, err := g.SetView(ci.Name, ci.X, ci.Y, ci.X+ci.W, ci.Y+ci.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			fmt.Println(g.Size())
			return err
		}
	}
	v.Title = "Connection Info"
	v.Clear()
	fmt.Fprintln(v, fmt.Sprintf("Port: %v\nBaud: %v\nConnected: %v", ci.currentStatus.GetPort(), ci.currentStatus.GetBaud(), ci.currentStatus.GetConnected()))

	return nil
}
