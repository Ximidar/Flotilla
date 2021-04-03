/*
* @Author: Ximidar
* @Date:   2019-02-20 23:27:18
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-21 02:54:55
 */

package NatsFile

import (
	"encoding/json"

	"github.com/Ximidar/Flotilla/DataStructures/ConfigStructures"
	"github.com/nats-io/nats.go"
)

// Config will query for the settings it needs and store them in a struct
type config struct {
	GcodePath string `json:"path.gcode"`

	nc *nats.Conn
}

// newConfig will return a new Config object. This is lower case because it only pertains to
// FileManager and no one else
func newConfig(nc *nats.Conn) (*config, error) {
	config := new(config)
	config.nc = nc
	return config, nil
}

// GetConfig will load the config into the config struct
func (conf *config) GetConfig() error {

	raw, err := ConfigStructures.RequestConfigJSON(conf.nc, "path.gcode")
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, conf)
	if err != nil {
		return err
	}
	return nil
}
