/*
* @Author: Ximidar
* @Date:   2018-12-22 15:55:19
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-04 15:01:41
 */

package PlayStatusTest

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/ximidar/Flotilla/FlotillaStatus/PlayStatus"
)

func TestPlayStatusSetup(t *testing.T) {

	sendCallback := func(bstatus []byte) {
		fmt.Println(string(bstatus))
	}
	_, err := PlayStatus.NewPlayStatus(sendCallback)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)

}

func TestChangeStates(t *testing.T) {
	var intendedState string
	var actualState string
	sendCallback := func(bstatus []byte) {
		status, err := PlayStructures.NewStatusFlagsFromJSON(bstatus)
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
		actualState = status.CurrentAction
		CommonTestTools.CheckEquals(t, intendedState, actualState)
	}
	ps, err := PlayStatus.NewPlayStatus(sendCallback)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	ps.SetReady()
	if !ps.IsReady {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Play
	intendedState = PlayStructures.PLAY
	err = ps.UpdateStatus(PlayStructures.PLAY)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)

	if !ps.IsPlaying || !ps.IsReady {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Pause
	intendedState = PlayStructures.PAUSE
	err = ps.UpdateStatus(PlayStructures.PAUSE)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	if !ps.IsPlaying || !ps.IsReady || !ps.IsPaused {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Try to set play again
	err = ps.UpdateStatus(PlayStructures.PLAY)
	if err == nil {
		err = errors.New("error was not thrown")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Resume
	intendedState = PlayStructures.RESUME
	err = ps.UpdateStatus(PlayStructures.RESUME)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	if !ps.IsPlaying || !ps.IsReady || ps.IsPaused {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Cancel
	intendedState = PlayStructures.CANCEL
	err = ps.UpdateStatus(PlayStructures.CANCEL)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	if ps.IsPlaying || ps.IsReady {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Play without Readying again
	intendedState = PlayStructures.PLAY
	err = ps.UpdateStatus(PlayStructures.PLAY)
	if err == nil {
		err = errors.New("error was not thrown")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}

	// Set Play
	ps.SetReady()
	intendedState = PlayStructures.PLAY
	err = ps.UpdateStatus(PlayStructures.PLAY)
	CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)

	if !ps.IsPlaying || !ps.IsReady {
		err = errors.New("State did not change")
		CommonTestTools.CheckErr(t, "TestPlayStatusSetup", err)
	}
}

func TestRollCall(t *testing.T) {
	rc, err := PlayStatus.NewRollCall()
	if err != nil {
		t.Fatal("Could not build a new roll call object")
	}

	Register1 := "register1"
	Register2 := "register2"
	Register3 := "register3"
	Register4 := "register4"
	Register5 := "register5"

	// Test sending an update before registering
	err = rc.RogerUp(Register1, PlayStructures.PLAY)
	if err == nil {
		t.Fatal("No Error Produced")
	}

	// Register nodes
	rc.RegisterNode(Register1)
	rc.RegisterNode(Register2)
	rc.RegisterNode(Register3)
	rc.RegisterNode(Register4)
	rc.RegisterNode(Register5)

	if !rc.AllNodesEqual() {
		t.Fatal("All nodes are equal")
	}

	// Update one node
	err = rc.RogerUp(Register1, PlayStructures.PLAY)
	if err != nil {
		t.Fatal("Error produced while updating status")
	}

	// Check if all nodes are equal, They should not be
	if rc.AllNodesEqual() {
		t.Fatal("All nodes are not equal.")
	}

	// Update all nodes
	err = rc.RogerUp(Register2, PlayStructures.PLAY)
	err = rc.RogerUp(Register3, PlayStructures.PLAY)
	err = rc.RogerUp(Register4, PlayStructures.PLAY)
	err = rc.RogerUp(Register5, PlayStructures.PLAY)
	if err != nil {
		t.Fatal("Error produced while updating status")
	}

	if !rc.AllNodesEqual() {
		t.Fatal("All nodes are supposed to be equal")
	}

	// Try to update a node with a bogus mode
	err = rc.RogerUp(Register1, "BogusMode")
	if err == nil {
		t.Fatal("Error was not thrown")
	}

}

func TestLargeRollCall(t *testing.T) {
	rc, err := PlayStatus.NewRollCall()
	if err != nil {
		t.Fatal("Could not build a new roll call object")
	}

	// Create a million nodes
	var millnodes []string
	for i := 0; i < 1000000; i++ {
		millnodes = append(millnodes, fmt.Sprintf("Node%v", i))
	}

	for index := range millnodes {
		rc.RegisterNode(millnodes[index])
	}

	if !rc.AllNodesEqual() {
		t.Fatal("All nodes are supposed to be equal")
	}

	exitSignal := make(chan bool, 10)

	checkNodes := func() {
		time.After(20 * time.Millisecond)
		for {
			<-time.After(10 * time.Millisecond)
			if rc.AllNodesEqual() {
				exitSignal <- true
				return
			}
		}
	}

	go checkNodes()
	// Update all nodes
	for index := range millnodes {
		select {
		case <-exitSignal:
			t.Fatal("Exit signal was reached before all nodes updated")
		default:
			err = rc.RogerUp(millnodes[index], PlayStructures.PLAY)
			if err != nil {
				t.Fatal("Error could not update node because:", err)
			}
		}

	}
	<-exitSignal

}
