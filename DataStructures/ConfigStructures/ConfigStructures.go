/*
* @Author: Ximidar
* @Date:   2019-02-16 22:41:46
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-16 23:05:16
 */

package ConfigStructures

import "github.com/nats-io/nats.go"

const (

	// Name is the name of the config
	Name = "Config."

	// GetVar is the name of the subject that will return a config var
	GetVar = Name + "GetVar"

	// UpdateConfig is a subject that will fire if a new config is available
	UpdateConfig = Name + "UpdateConfig"
)

// CallbackFunc is a type of function for updating the callback
type CallbackFunc func()

// FlotillaConfig will be used to interact with the nats server to get variables
// This struct will be used to inform the program when a config change occurs
type FlotillaConfig struct {
	callbacks        *CallbackFunc
	variablesToWatch []string

	NC *nats.Conn
}

// NewFlotillaConfig will return a FlotillaConfig object
func NewFlotillaConfig(nc *nats.Conn, cf CallbackFunc) (*FlotillaConfig, error) {
	fc := new(FlotillaConfig)
	fc.NC = nc

	// subscribe to config updates
	err := fc.makesubs()
	if err != nil {
		return nil, err
	}

	return fc, nil
}

func (fc *FlotillaConfig) makesubs() error {
	_, err := fc.NC.Subscribe(UpdateConfig, fc.updateConfig)
	return err
}

// updateConfig will receive the push to update the config
func (fc *FlotillaConfig) updateConfig(msg *nats.Msg) {

}
