/*
* @Author: Ximidar
* @Date:   2019-02-21 01:36:15
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-21 02:47:45
 */

package Config_test

import (
	"encoding/json"
	"testing"

	"github.com/nats-io/gnatsd/test"
	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/CommonTools/NatsConnect"
	"github.com/ximidar/Flotilla/DataStructures/ConfigStructures"
	"github.com/ximidar/Flotilla/NodeLauncher/Config"
)

func SetupServer() {
	test.RunDefaultServer()
}

func TestConfigNatsInterface(t *testing.T) {
	// Start a default Nats Server
	test.RunDefaultServer()

	// Make a nats.Conn
	conn, err := NatsConnect.DefaultConn(nats.DefaultURL, "configTest")
	CommonTestTools.CheckErr(t, "Cannot Create basic Connection", err)

	// Create Prerequisite components
	err, rf, _ := SetupRootFolder()
	CommonTestTools.CheckErr(t, "Test Loader Couldn't make the test location", err)
	defer CleanRootFolder()

	ConfigLoader, err := Config.NewLoader(rf)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't make the Loader object", err)

	// Create the object we are testing
	config, err := Config.NewNats(ConfigLoader)
	CommonTestTools.CheckErr(t, "Cannot Create Config Nats", err)
	err = config.StartNatsInteraction(conn)
	CommonTestTools.CheckErr(t, "Cannot Start Nats Interaction", err)

	// Request a value for the gcode and the path variables
	type fileconfig struct {
		RootPath  string `json:"path.root"`
		GcodePath string `json:"file_manager.gcode_folder"`
	}

	raw, err := ConfigStructures.RequestConfigJSON(conn, "path.root", "file_manager.gcode_folder")
	CommonTestTools.CheckErr(t, "Cannot Request configuration", err)

	conf := new(fileconfig)
	err = json.Unmarshal(raw, conf)
	CommonTestTools.CheckErr(t, "Cannot Unmarshal json", err)

	// Check that things are kosher
	CommonTestTools.CheckEquals(t, conf.GcodePath, "/GCODE")
}
