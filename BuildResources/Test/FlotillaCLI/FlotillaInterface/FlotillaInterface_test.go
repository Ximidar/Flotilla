/*
* @Author: Ximidar
* @Date:   2018-08-25 10:51:03
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-11-29 12:57:10
 */
package FlotillaInterface_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

// TODO add functions to turn NATS server on and off and different nodes on and off

func Test_Get_Available_Ports(t *testing.T) {
	fmt.Println("Testing Get Available Ports")
	mgo, err := FlotillaInterface.NewFlotillaInterface()

	if err != nil {
		t.Fatal(err)
	}
	ports, err := mgo.CommGetAvailablePorts()
	if err != nil {
		fmt.Println("Could not get available ports", err)
		t.Fatal(err)
	}

	fmt.Println(ports)
}

func Test_Comm_set_up_and_write(t *testing.T) {
	mgo, err := FlotillaInterface.NewFlotillaInterface()

	if err != nil {
		t.Fatal(err)
	}

	mgo.CommSetConnectionOptions("/dev/ttyACM0", 115200)
	mgo.CommConnect()
	defer mgo.CommDisconnect()

	duration := time.Duration(5 * time.Second)
	time.Sleep(duration)

	stop_reading := false

	read_func := func() {
		for read := range mgo.EmitLine {
			if stop_reading {
				break
			}

			fmt.Printf("%s", read)
		}
	}

	go read_func()

	pause_dur := time.Duration(100 * time.Millisecond)
	writes := []string{"Hello!", "My", "Name", "Is", "Matt"}

	for _, write := range writes {
		mgo.CommWrite(write)
		time.Sleep(pause_dur)
	}

	stop_reading = true

}

func Test_Get_Structure(t *testing.T) {
	fi, err := FlotillaInterface.NewFlotillaInterface()

	if err != nil {
		t.Fatal(err)
	}

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
