/*
* @Author: Ximidar
* @Date:   2018-12-19 16:00:13
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-20 16:17:25
 */

package StatusMonitor

import "github.com/nats-io/nats.go"

// StatusMonitor is a struct to hold all the different status monitors
type StatusMonitor struct {
	TempMonitor *TemperatureMonitor
}

// NewStatusMonitor will constuct a StatusMonitor
func NewStatusMonitor(pubtemp PublishTemperature) (*StatusMonitor, error) {
	var err error
	sm := new(StatusMonitor)
	sm.TempMonitor, err = NewTemperatureMonitor(pubtemp)

	if err != nil {
		return nil, err
	}

	return sm, nil

}

// GetComm will distribute the incoming comm messages to the different statuses
func (sm *StatusMonitor) GetComm(msg *nats.Msg) {

}
