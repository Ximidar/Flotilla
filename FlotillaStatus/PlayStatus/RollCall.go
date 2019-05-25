/*
* @Author: Ximidar
* @Date:   2019-01-02 22:24:25
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-29 15:06:36
 */

package PlayStatus

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
)

// RollCall will accept a call from every node out there. When a state change is issued by
// one node all nodes will need to roger up and accept the state change.
type RollCall struct {
	Roll map[string]string
	mux  sync.Mutex
}

// NewRollCall will make a new Roll Call Object
func NewRollCall() (*RollCall, error) {
	rc := new(RollCall)
	rc.Roll = make(map[string]string)
	return rc, nil
}

// RegisterNode will register the node in the Roll.
func (rc *RollCall) RegisterNode(name string) {

	// Make sure the node isn't already registered
	_, ok := rc.Roll[name]
	if ok {
		return
	}

	// Register the node
	rc.mux.Lock()
	rc.Roll[name] = "Idle"
	rc.mux.Unlock()
}

// RogerUp will change all the nodes to their current status
func (rc *RollCall) RogerUp(name, mode string) error {
	// if the node isn't registered throw an error
	_, ok := rc.Roll[name]
	if !ok {
		return errors.New("node is not registered")
	}

	// if the mode change isn't a real mode then throw an error
	if !rc.checkMode(mode) {
		return errors.New("mode is not correct")
	}
	rc.mux.Lock()
	rc.Roll[name] = mode
	rc.mux.Unlock()

	return nil
}

// checkMode will check if the mode is the real
func (rc *RollCall) checkMode(mode string) bool {
	err := PlayStructures.CheckSubj(mode)
	if err != nil {
		return false
	}
	return true

}

// AllNodesEqual will check if all nodes are the same or not.
func (rc *RollCall) AllNodesEqual(action string) bool {
	rc.mux.Lock()
	defer rc.mux.Unlock()
	for _, val := range rc.Roll {
		if val != action {
			return false
		}
	}
	return true
}

// PrintNodeStatus will print the current status
func (rc *RollCall) PrintNodeStatus() {
	rc.mux.Lock()
	for key, val := range rc.Roll {
		fmt.Printf("Name: %v Status: %v\n", key, val)
	}
	rc.mux.Unlock()
}

// DeleteNode will delete the registration for the name of a node.
// If the node doesn't exist it will not return an error.
func (rc *RollCall) DeleteNode(name string) {
	rc.mux.Lock()
	defer rc.mux.Unlock()

	delete(rc.Roll, name)
}
