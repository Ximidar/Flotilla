/*
* @Author: Ximidar
* @Date:   2019-02-16 21:15:36
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-17 14:20:04
 */

package Config_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/ConfigStructures"
	"github.com/ximidar/Flotilla/NodeLauncher/Config"
	"github.com/ximidar/Flotilla/NodeLauncher/RootFolder"
)

var TestLocation = "/tmp/LoaderTest"

func TestLoader(t *testing.T) {
	err, rf, _ := SetupRootFolder()
	CommonTestTools.CheckErr(t, "Test Loader Couldn't make the test location", err)
	defer CleanRootFolder()

	ConfigLoader, err := Config.NewLoader(rf)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't make the Loader object", err)

	// Get the GCODE Folder default Which should be "/GCODE"
	expected := "/GCODE"
	result := ConfigLoader.Config.GetString("file_manager.gcode_folder")

	CommonTestTools.CheckEquals(t, expected, result)

	// Test the GetVar Function with json
	cr := ConfigStructures.ConfigReply{
		VarNames: []string{"commango.extra_comms", "commango.extra_bauds"},
		ToJSON:   true,
	}

	jsonMap, err := ConfigLoader.GetVars(cr)
	fmt.Println(string(jsonMap))
	CommonTestTools.CheckErr(t, "Test Loader Couldn't get the commango args", err)
	type commangoConfig struct {
		ExtraComms []string `json:"commango.extra_comms"`
		ExtraBauds []string `json:"commango.extra_bauds"`
	}

	commConfig := commangoConfig{}
	err = json.Unmarshal(jsonMap, &commConfig)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't decode the map", err)

	fmt.Println(commConfig)

	CommonTestTools.CheckEquals(t, "/tmp/fakeprinter", commConfig.ExtraComms[0])

	// Test using a map to byte structures
	cr = ConfigStructures.ConfigReply{
		VarNames: []string{"nats_server"},
		ToJSON:   false,
	}

	bmap, err := ConfigLoader.GetVars(cr)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't get the nats_server args", err)
	nats_server := make(map[string][]byte)
	err = ConfigStructures.DecodeBytes(bmap, &nats_server)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't decode the map", err)

	// try to get some values out
	nats_server2 := make(map[string]interface{})
	err = ConfigStructures.DecodeBytes(nats_server["nats_server"], &nats_server2)
	CommonTestTools.CheckErr(t, "Test Loader Couldn't decode the map", err)

	CommonTestTools.CheckEquals(t, "/workspace/certs/flotilla.cert", nats_server2["tls_cert"].(string))
	CommonTestTools.CheckEquals(t, true, nats_server2["http_server"].(bool))
	CommonTestTools.CheckEquals(t, 5070, nats_server2["http_port"].(int))
	CommonTestTools.CheckEquals(t, false, nats_server2["tls"].(bool))

}

func SetupRootFolder() (error, *RootFolder.RootFolder, *RootFolder.PopulateFolder) {
	rf, err := RootFolder.GenerateRootFolder(TestLocation + "/Flotilla")
	if err != nil {
		return err, nil, nil
	}
	pf, err := RootFolder.NewPopulateFolder(rf, "amd64") // Change this if you arent amd64
	if err != nil {
		return err, nil, nil
	}
	err = pf.PackageStaticFiles()
	if err != nil {
		return err, nil, nil
	}
	return nil, rf, pf
}

func CleanRootFolder() error {
	return os.RemoveAll(TestLocation)
}
