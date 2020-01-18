/*
* @Author: Ximidar
* @Date:   2018-12-21 17:45:16
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-29 15:58:33
 */

package PlayStatus

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
)

//Errors
var (
	ErrAlreadyPlaying = errors.New("already playing")
	ErrNotPlaying     = errors.New("not playing")
	ErrNotReady       = errors.New("not ready")
	ErrAlreadyPaused  = errors.New("already paused")
	ErrInErrorState   = errors.New("in error state")
)

// SendStatus is a function that will be used to relay the status to nats
type SendStatus func(bstatus []byte)

// PlayStatus will keep track of the current status of play
type PlayStatus struct {
	proposedAction    string
	sendStatus        SendStatus
	currentStatusJSON []byte
	CurrentAction     string `json:"status"`
	IsPlaying         bool   `json:"isPlaying"`
	IsPaused          bool   `json:"isPaused"`
	IsError           bool   `json:"isError"`
	IsReady           bool   `json:"isReady"`
	Error             error  `json:"error"`

	doneCounter int
}

// NewPlayStatus will construct a PlayStatus
func NewPlayStatus(statusCallback SendStatus) (*PlayStatus, error) {
	ps := new(PlayStatus)
	ps.sendStatus = statusCallback
	ps.reset()
	return ps, nil
}

// UpdateStatus will update the status
func (ps *PlayStatus) UpdateStatus(newStatus string) error {
	fmt.Println("Setting Status to", newStatus)
	// Check the subj
	err := PlayStructures.CheckSubj(newStatus)
	if err != nil {
		return err
	}

	ps.proposedAction = newStatus
	switch newStatus {
	case PlayStructures.PLAY:
		err = ps.setPlay()
	case PlayStructures.CANCEL:
		err = ps.setCancel()
	case PlayStructures.PAUSE:
		err = ps.setPause()
	case PlayStructures.RESUME:
		err = ps.setResume()
	case PlayStructures.DONE:
		err = ps.setDone()
	case PlayStructures.RESET:
		ps.reset()
	case PlayStructures.IDLE:
		// TODO make functions for IDLE
		err = errors.New("Not implemented yet")
	default:
		return fmt.Errorf("invalid status: %v", newStatus)
	}

	if err != nil {
		fmt.Println("Error!", err)
		return err
	}

	err = ps.PackageAndSendStatus()
	if err != nil {
		return err
	}

	return nil
}

// CanSwitchToStatus will return a bool based on if we can switch to a proposed status
func (ps *PlayStatus) CanSwitchToStatus(status string) error {
	// Check the subj
	err := PlayStructures.CheckSubj(status)
	if err != nil {
		return err
	}

	switch status {
	case PlayStructures.PLAY:
		err = ps.CanSetPlay()
	case PlayStructures.CANCEL:
		err = ps.CanSetCancel()
	case PlayStructures.PAUSE:
		err = ps.CanSetPause()
	case PlayStructures.RESUME:
		err = ps.CanSetResume()
	case PlayStructures.DONE:
		err = ps.CanSetDone()
	case PlayStructures.RESET:
		err = nil
	case PlayStructures.IDLE:
		// TODO make functions for IDLE
		err = errors.New("Not implemented yet")
	default:
		err = fmt.Errorf("invalid status: %v", status)
	}

	if err != nil {
		fmt.Println("Error!", err)
		return err
	}

	return nil
}

// PackageAndSendStatus will take the current status and send it out
func (ps *PlayStatus) PackageAndSendStatus() error {
	statusb, err := ps.PackageJSON()
	if err != nil {
		return err
	}
	ps.currentStatusJSON = statusb
	ps.sendStatus(statusb)

	return nil

}

// PackageJSON will package the current status into a JSON object
func (ps *PlayStatus) PackageJSON() ([]byte, error) {
	bps, err := json.Marshal(ps)
	return bps, err
}

// GetCurrentJSON will return the saved json state
func (ps *PlayStatus) GetCurrentJSON() []byte {
	if len(ps.currentStatusJSON) == 0 {
		pjson, err := ps.PackageJSON()
		if err != nil {
			return ps.currentStatusJSON
		}
		ps.currentStatusJSON = pjson

	}
	return ps.currentStatusJSON
}

// CanSetPlay checks whether we can currently set play or not
func (ps *PlayStatus) CanSetPlay() error {
	// Take care of error states
	if ps.CurrentAction == PlayStructures.PLAY || ps.IsPlaying {
		fmt.Println("Not setting play. CurrentAction:", ps.CurrentAction, "IsPlaying?:", ps.IsPlaying)
		return ErrAlreadyPlaying
	}
	return nil
}

// setPlay will respond to a request to set play
func (ps *PlayStatus) setPlay() error {
	err := ps.CanSetPlay()
	if err != nil {
		return err
	}

	// Change status to play
	ps.CurrentAction = ps.proposedAction
	ps.IsPlaying = true
	return nil
}

// CanSetPause will return an error if we cannot set pause
func (ps *PlayStatus) CanSetPause() error {
	// Take care of error states
	if ps.CurrentAction == PlayStructures.PAUSE || ps.IsPaused {
		return ErrAlreadyPaused
	}
	return nil
}

func (ps *PlayStatus) setPause() error {
	err := ps.CanSetPause()
	if err != nil {
		return err
	}
	// Change status
	ps.CurrentAction = ps.proposedAction
	ps.IsPaused = true
	return nil
}

// CanSetResume will check if we can set resume or not
func (ps *PlayStatus) CanSetResume() error {
	if ps.CurrentAction == PlayStructures.RESUME ||
		ps.CurrentAction == PlayStructures.PLAY ||
		!ps.IsPaused {
		return ErrAlreadyPlaying
	}
	return nil
}

func (ps *PlayStatus) setResume() error {
	err := ps.CanSetResume()
	if err != nil {
		return err
	}

	// Change status
	ps.CurrentAction = ps.proposedAction
	ps.IsPaused = false

	return nil
}

// CanSetCancel will check if we can set cancel or not
func (ps *PlayStatus) CanSetCancel() error {
	if ps.CurrentAction == PlayStructures.CANCEL ||
		!ps.IsPlaying {
		return ErrNotPlaying
	}
	return nil
}

func (ps *PlayStatus) setCancel() error {
	err := ps.CanSetCancel()
	if err != nil {
		return err
	}

	// Change status
	ps.CurrentAction = ps.proposedAction
	ps.reset()

	return nil
}

// CanSetDone checks if we can set done or not
func (ps *PlayStatus) CanSetDone() error {
	if ps.CurrentAction == PlayStructures.DONE ||
		!ps.IsPlaying {
		return ErrNotPlaying
	}
	return nil
}

func (ps *PlayStatus) setDone() error {
	err := ps.CanSetDone()
	if err != nil {
		return err
	}

	// Change status
	ps.CurrentAction = ps.proposedAction
	ps.reset()

	return nil
}

// SetError will be used to set the error state
func (ps *PlayStatus) SetError(errorMess string) {
	ps.reset()
	ps.IsError = true
	ps.Error = errors.New(errorMess)
}

// SetReady will be used to set the ready state
func (ps *PlayStatus) SetReady() {
	ps.reset()
	ps.IsReady = true
}

func (ps *PlayStatus) reset() {
	ps.CurrentAction = PlayStructures.DONE
	ps.IsError = false
	ps.IsPlaying = false
	ps.IsPaused = false
	ps.IsReady = false
	ps.Error = nil

	// TODO figure out a way to set IDLE correctly
	// Set the Idle action after setting the Done Action
	// time.AfterFunc(2*time.Second, func() {
	// 	ps.CurrentAction = PlayStructures.IDLE
	// 	statusb, _ := ps.PackageJSON()

	// 	ps.currentStatusJSON = statusb
	// 	ps.sendStatus(statusb)

	// })

}

// Monitor will look for start messages and reset the status accordingly
func (ps *PlayStatus) Monitor(line string) {
	line = strings.Replace(line, "\n", "", -1)
	line = strings.Replace(line, " ", "", -1)
	line = strings.ToLower(line)
	if line == "start" {
		fmt.Println("Got Start! Ready!")
		ps.reset()
		ps.SetReady()
	}
}
