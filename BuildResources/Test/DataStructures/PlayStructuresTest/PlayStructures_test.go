/*
* @Author: Ximidar
* @Date:   2018-12-22 21:11:06
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-22 22:07:39
 */

package PlayStructuresTest

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	PS "github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
)

func TestPlayStructures(t *testing.T) {
	so, _ := PS.NewStatusObserver()
	// Make funcs
	var helloOn bool
	var hiOn bool
	var alohaOn bool
	var herroOn bool
	hello := func() {
		fmt.Println("Hello, playground")
		helloOn = true
	}

	hi := func() {
		fmt.Println("Hi, playground")
		hiOn = true
	}

	aloha := func() {
		fmt.Println("aloha, playground")
		alohaOn = true
	}

	herro := func() {
		fmt.Println("herro, playground")
		herroOn = true
	}

	resetBools := func() {
		helloOn = false
		hiOn = false
		alohaOn = false
		herroOn = false
	}
	register1, err := so.AddFunction(PS.PLAY, hello)
	CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	register2, err := so.AddFunction(PS.PLAY, hi)
	CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	register3, err := so.AddFunction(PS.PLAY, aloha)
	CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	register4, err := so.AddFunction(PS.PLAY, herro)
	CommonTestTools.CheckErr(t, "TestPlayStructures", err)

	fmt.Println(register1, register2, register3, register4)
	fmt.Println()

	// Update the play Status
	sf := getStatus(PS.PLAY)
	so.UpdateStatus(sf)
	<-time.After(25 * time.Millisecond)
	fmt.Println()
	fmt.Println()

	if !helloOn || !hiOn || !alohaOn || !herroOn {
		err = errors.New("Callbacks did not fire")
		CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	}

	// Delete a key and update again

	err = so.DeleteFunction(PS.PLAY, register2)
	resetBools()
	CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	so.UpdateStatus(sf)
	<-time.After(25 * time.Millisecond)

	if !helloOn || hiOn || !alohaOn || !herroOn {
		err = errors.New("callback hi did not get deleted")
		CommonTestTools.CheckErr(t, "TestPlayStructures", err)
	}
}

func getStatus(status string) []byte {
	sf := PS.StatusFlags{CurrentAction: status}
	encoded, _ := json.Marshal(sf)
	return encoded

}
