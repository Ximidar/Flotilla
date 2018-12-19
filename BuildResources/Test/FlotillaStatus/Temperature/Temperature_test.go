/*
* @Author: Ximidar
* @Date:   2018-12-18 19:06:46
* @Last Modified by:   Ximidar
* @Last Mod(ified time: 2018-12-18 19:06:46
 */

package TemperatureTests

import (
	"errors"
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/FlotillaStatus/StatusMonitor"
)

func TestTemperatureMonitor(t *testing.T) {
	temps := new(StatusMonitor.TemperatureMonitor)
	err := temps.CompileRegex()
	if err != nil {
		CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)
	}

	// Test a sample temp text
	sampletemp := "T:21.25 /23.01 B:22.50 /50.72 @:0 B@:0"
	err = stringTestTB(temps, sampletemp, 21.25, 23.01, 22.50, 50.72)
	CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)

	// Test for negatives
	sampletemp = "T:-21.25 /-23.01 B:-22.50 /-50.72 @:0 B@:0"
	err = stringTestTB(temps, sampletemp, -21.25, -23.01, -22.50, -50.72)
	CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)

	// Test for changed spaces
	sampletemp = "T:-21.25/23.01B:22.50     /-50.72@:0B@:0"
	err = stringTestTB(temps, sampletemp, -21.25, 23.01, 22.50, -50.72)
	CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)
}

func stringTestTB(temps *StatusMonitor.TemperatureMonitor, rawtest string, tempT, targetT float64, tempB, targetB float64) error {
	var err error
	temps.UpdateTemperature(rawtest)

	if temps.CurrentTemperature.Temp["T"] != tempT {
		err = errors.New("T Temp Was not Set Correctly")
	} else if temps.CurrentTemperature.Target["T"] != targetT {
		err = errors.New("T Target Was not Set Correctly")
	} else if temps.CurrentTemperature.Temp["B"] != tempB {
		err = errors.New("B Temp Was not Set Correctly")
	} else if temps.CurrentTemperature.Target["B"] != targetB {
		err = errors.New("B Target Was not Set Correctly")
	}
	return err
}
