/*
* @Author: Ximidar
* @Date:   2019-02-16 22:36:27
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-17 14:13:33
 */

package Config

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/DataStructures/ConfigStructures"
)

// Nats will handle the nats interface for the config
type Nats struct {
	NC     *nats.Conn
	loader *Loader
}

// NewNats will construct a Nats object
func NewNats(loader *Loader) (*Nats, error) {
	configNats := new(Nats)
	configNats.loader = loader

	return configNats, nil
}

// StartNatsInteraction will start serving requests over nats
func (cnats *Nats) StartNatsInteraction(nc *nats.Conn) error {
	cnats.NC = nc
	return cnats.registerReqs()
}

func (cnats *Nats) registerReqs() error {
	_, err := cnats.NC.Subscribe(ConfigStructures.GetVar, cnats.GetVar)
	return err
}

// GetVar will return either a map or a json byte slice
func (cnats *Nats) GetVar(msg *nats.Msg) {

	// Convert to ConfigReply
	configReply, err := cnats.decodeConfigReply(msg.Data)
	if err != nil {
		fmt.Println("Could not decode to Config Reply Because: ", err)
		cnats.NC.Publish(msg.Reply, []byte(nil))
		return
	}

	mData, err := cnats.loader.GetVars(configReply)
	if err != nil {
		// respond with no data
		cnats.NC.Publish(msg.Reply, []byte(nil))
	}

	// respond with bytes
	cnats.NC.Publish(msg.Reply, mData)

}

func (cnats *Nats) decodeConfigReply(cr []byte) (ConfigStructures.ConfigReply, error) {

	configR := ConfigStructures.ConfigReply{}

	err := json.Unmarshal(cr, &configR)
	if err != nil {
		return ConfigStructures.ConfigReply{}, err
	}

	return configR, nil
}
