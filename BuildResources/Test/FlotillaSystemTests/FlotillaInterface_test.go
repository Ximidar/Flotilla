/*
* @Author: Ximidar
* @Date:   2018-08-25 10:51:03
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-22 17:48:08
 */
package FlotillaSystemTest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

func TestGetAvailablePorts(t *testing.T) {

	exitChan, err := StartTestFlotilla()
	exitFunc := func() {
		if exitChan != nil {
			exitChan <- true
			<-exitChan
		}
	}
	// defer exit just in case
	defer exitFunc()
	CommonTestTools.CheckErr(t, "Could not start Flotilla Instance", err)

	fmt.Println("Testing Get Available Ports")
	mgo, err := FlotillaInterface.NewFlotillaInterface()
	CommonTestTools.CheckErr(t, "Could not create Flotilla Interface", err)

	ports, err := mgo.CommGetAvailablePorts()
	CommonTestTools.CheckErr(t, "Could not get available ports", err)
	fmt.Println(ports)
}

func Test_Get_Structure(t *testing.T) {
	exitChan, err := StartTestFlotilla()
	exitFunc := func() {
		if exitChan != nil {
			exitChan <- true
			<-exitChan
		}
	}
	// defer exit just in case
	defer exitFunc()
	CommonTestTools.CheckErr(t, "Could not start Flotilla Instance", err)

	fi, err := FlotillaInterface.NewFlotillaInterface()
	CommonTestTools.CheckErr(t, "Could not create Flotilla Interface", err)

	structure, err := fi.GetFileStructure()
	if err != nil {
		t.Fatal(err)
	}
	PrintStructure(structure)

	data, ok := structure["root"]
	if !ok {

		t.Fatal("Could not navigate map")
	}
	fmt.Println(data.Path)
}

func PrintStructure(structure map[string]*Files.File) {
	marshed, err := json.MarshalIndent(structure, "", "    ")
	if err != nil {
		fmt.Println("Couldn't get json structure:", err)
		return
	}
	fmt.Println(string(marshed))
}
