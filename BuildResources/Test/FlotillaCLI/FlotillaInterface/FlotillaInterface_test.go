/*
* @Author: Ximidar
* @Date:   2018-08-25 10:51:03
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-21 17:49:03
 */
package FlotillaInterface_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/Flotilla_CLI/FlotillaInterface"
)

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
