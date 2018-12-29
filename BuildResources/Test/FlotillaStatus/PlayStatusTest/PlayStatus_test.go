/*
* @Author: Ximidar
* @Date:   2018-12-22 15:55:19
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-22 16:26:07
 */

package PlayStatusTest

import (
	"errors"
	"fmt"
	"testing"

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
