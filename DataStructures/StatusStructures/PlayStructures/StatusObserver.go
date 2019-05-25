/*
* @Author: Ximidar
* @Date:   2019-01-16 15:23:59
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-17 14:29:16
 */

package PlayStructures

import (
	"encoding/json"
	"fmt"
	"math"

	nats "github.com/nats-io/go-nats"
)

// StatusFlags is a convinient struct to quickly find the current status
type StatusFlags struct {
	CurrentAction string `json:"status"`
	IsPlaying     bool   `json:"isPlaying"`
	IsPaused      bool   `json:"isPaused"`
	IsError       bool   `json:"isError"`
	IsReady       bool   `json:"isReady"`
	Error         error  `json:"error"`
}

// NewStatusFlagsFromJSON Construct the statusflags object
func NewStatusFlagsFromJSON(status []byte) (StatusFlags, error) {
	sf := StatusFlags{}
	err := json.Unmarshal(status, &sf)
	return sf, err
}

// UpdateUknownStatus is a function that will be fired off if a callback is not
// registered on any status changes
type UpdateUknownStatus func(action string)

// StatusObserver will fire functions if a status is fired off
type StatusObserver struct {
	CurrentStatus        StatusFlags
	CallbackUknownStatus UpdateUknownStatus
	callbackFuncs        map[string]map[int]func()
}

// NewStatusObserver will create a new status observer
func NewStatusObserver() (*StatusObserver, error) {
	so := new(StatusObserver)
	so.callbackFuncs = make(map[string]map[int]func())
	so.CallbackUknownStatus = nil
	return so, nil
}

// AddUknownStatusFunc will add the Callback for anytime there is not a func that is known
func (so *StatusObserver) AddUknownStatusFunc(callback UpdateUknownStatus) {
	so.CallbackUknownStatus = callback
}

// AddFunction will add a function to be called when a certain status is set
// It will also return a id number to deregister your callback
func (so *StatusObserver) AddFunction(subj string, callback func()) (int, error) {
	if err := CheckSubj(subj); err != nil {
		return -1, err
	}
	// get the next number to add
	maxNum := 0

	if so.callbackFuncs[subj] == nil {
		so.callbackFuncs[subj] = make(map[int]func())
	}

	for key := range so.callbackFuncs[subj] {
		maxNum = int(math.Max(float64(maxNum), float64(key)))
	}
	maxNum++

	// Add funcs to callback
	so.callbackFuncs[subj][maxNum] = callback

	return maxNum, nil
}

// DeleteFunction will deregister any functions
func (so *StatusObserver) DeleteFunction(subj string, id int) error {
	var err error
	err = CheckSubj(subj)
	if err != nil {
		return err
	}
	// This will do nothing if the key does not exist
	delete(so.callbackFuncs[subj], id)
	return err
}

// UpdateStatus will update the status and call the callbacks
func (so *StatusObserver) UpdateStatus(bstatus []byte) {
	newstatus, err := NewStatusFlagsFromJSON(bstatus)
	if err != nil {
		return
	}
	so.CurrentStatus = newstatus
	fmt.Println("UPDATED STATUS!", so.CurrentStatus.CurrentAction)
	so.call(so.CurrentStatus.CurrentAction)
}

// UpdateStatusFromNats is a simple function to plug into the nats server to monitor for changes
func (so *StatusObserver) UpdateStatusFromNats(msg *nats.Msg) {
	so.UpdateStatus(msg.Data)
}

func (so *StatusObserver) call(action string) {
	callbacks, ok := so.callbackFuncs[action]

	if !ok {
		if so.CallbackUknownStatus != nil {
			so.CallbackUknownStatus(action)
		}
		return
	}

	for _, val := range callbacks {
		go val()
	}
}
