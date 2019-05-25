/*
* @Author: Ximidar
* @Date:   2019-01-16 16:22:16
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-29 15:31:42
 */

package PlayStatus

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrChangeInProgress = errors.New("Change in progress")

// Control will mix the PlayStatus and RollCall types and coordinate forces between the two
type Control struct {
	PlayStatus *PlayStatus
	RollCall   *RollCall

	ChangeInProgress bool
	mux              sync.Mutex
}

// NewControl will construct a control object
func NewControl(publishStatusFunc SendStatus) (*Control, error) {

	// Construct Control
	control := new(Control)
	control.ChangeInProgress = false

	var err error
	control.PlayStatus, err = NewPlayStatus(publishStatusFunc)
	if err != nil {
		return nil, err
	}
	control.PlayStatus.SetReady()
	control.RollCall, err = NewRollCall()

	if err != nil {
		return nil, err
	}
	return control, nil
}

// ProposeAction will take in an action then propose it to all nodes, then
// wait until all nodes have signaled they are ready for that action to be set.
// Then it will publish that action
func (ctrl *Control) ProposeAction(action string) error {
	if ctrl.ChangeInProgress {
		return ErrChangeInProgress
	}
	ctrl.ChangeInProgress = true
	// Check if action is valid and that we can move to it
	fmt.Println("Can we switch to this status?", action)
	err := ctrl.PlayStatus.CanSwitchToStatus(action)
	if err != nil {
		return err
	}

	// Change action
	fmt.Println("Proposing", action)
	err = ctrl.PlayStatus.UpdateStatus(action)
	if err != nil {
		return err
	}

	// Wait for all nodes to roger up
	ctrl.SyncAllNodes(action)

	ctrl.ChangeInProgress = false

	return nil

}

// SyncAllNodes will return when all nodes have changed to their prospective statuses
func (ctrl *Control) SyncAllNodes(action string) {
	for {
		if ctrl.RollCall.AllNodesEqual(action) {
			return
		}
		<-time.After(300 * time.Millisecond)
	}
}
