/*
* @Author: Ximidar
* @Date:   2018-12-20 16:19:40
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-02 13:01:52
 */

package TemperatureStructures

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
	"github.com/Ximidar/Flotilla/FlotillaStatus/StatusMonitor"
)

const (
	// Name of Nats region
	Name = "FlotillaTemperature."

	// SetTemp will set the temperature
	SetTemp = Name + "SetTemp"

	// GetTemp will return the current temperature once
	GetTemp = Name + "GetTemp"

	// PubTemp is the publisher for the current temp
	PubTemp = Name + "PublishTemp"

	// GetTempHistory will return the history of temperature
	GetTempHistory = Name + "GetTempHistory"
)

////////////////////////////////////////
// Set Temperature
////////////////////////////////////////

// SetTemperature is a structure that will ask Nats to send the temp
type SetTemperature struct {
	Tool string `json:"tool"`
	Temp uint64 `json:"temp"`
}

// NewSetTemperature will construct a SetTemperature struct
func NewSetTemperature(tool string, temp uint64) *SetTemperature {
	st := new(SetTemperature)
	st.Tool = tool
	st.Temp = temp

	return st
}

// NewSetTemperatureFromMSG will construct a SetTemperature struct from a nats msg
func NewSetTemperatureFromMSG(msg *nats.Msg) (*SetTemperature, error) {
	bst := new(SetTemperature)
	err := json.Unmarshal(msg.Data, bst)
	if err != nil {
		return nil, err
	}
	return bst, nil
}

// Send Will send the SetTemperature request
func (st *SetTemperature) Send(nc *nats.Conn) (*nats.Msg, error) {
	stb, err := json.Marshal(st)
	if err != nil {
		return nil, err
	}
	resp, err := nc.Request(SetTemp, stb, nats.DefaultTimeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

////////////////////////////////////////
// GetTemperature
////////////////////////////////////////

// GetTemperature will query the NatsServer for the current Temperature
func GetTemperature(nc *nats.Conn) (*StatusMonitor.Temperature, error) {

	// Request the temperature
	resp, err := nc.Request(GetTemp, []byte{}, nats.DefaultTimeout)
	if err != nil {
		return nil, err
	}

	temp, err := StatusMonitor.NewTemperatureFromMSG(resp)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// GetTemperatureHistory will get the temperature history
func GetTemperatureHistory(nc *nats.Conn) ([]*StatusMonitor.Temperature, error) {
	resp, err := nc.Request(GetTempHistory, []byte{}, nats.DefaultTimeout)
	if err != nil {
		return nil, err
	}

	tempHistory := []*StatusMonitor.Temperature{}
	err = json.Unmarshal(resp.Data, tempHistory)
	if err != nil {
		return nil, err
	}

	return tempHistory, nil
}
