/*
* @Author: Ximidar
* @Date:   2018-12-18 19:06:46
* @Last Modified by:   Ximidar
* @Last Mod(ified time: 2018-12-18 19:06:46
 */

package StatusMonitor

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Ximidar/Flotilla/BuildResources/Test/CommonTestTools"
)

func TestTemperatureMonitor(t *testing.T) {
	publishTemp := func(t *Temperature) {
		fmt.Println(t)
	}
	temps := new(TemperatureMonitor)
	err := temps.CompileRegex()
	if err != nil {
		CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)
	}
	temps.PublishTemperature = publishTemp
	temps.ResetTemps()

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

func stringTestTB(temps *TemperatureMonitor, rawtest string, tempT, targetT float64, tempB, targetB float64) error {
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

func TestGetTempCommand(t *testing.T) {
	publishTemp := func(t *Temperature) {
		fmt.Println(t)
	}
	temps := new(TemperatureMonitor)
	err := temps.CompileRegex()
	if err != nil {
		CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)
	}
	temps.PublishTemperature = publishTemp
	temps.ResetTemps()

	// Test bed
	expected := "M140 S100"
	result, _ := temps.GetTempCommand("B", 100)

	if expected != result {
		err := fmt.Errorf("expected did not equal result: %v", result)
		CommonTestTools.CheckErr(t, "TestGetTempCommand", err)
	}

	// test T1
	expected = "M104 T1 S100"
	result, _ = temps.GetTempCommand("T1", 100)

	if expected != result {
		err := fmt.Errorf("expected did not equal result: %v", result)
		CommonTestTools.CheckErr(t, "TestGetTempCommand", err)
	}

	// test err
	_, err = temps.GetTempCommand("weird heater", 1000)
	if err == nil {
		err = errors.New("Failed to produce error")
		CommonTestTools.CheckErr(t, "TestGetTempCommand", err)
	}
}

func TestTempHistory(t *testing.T) {
	publishTemp := func(t *Temperature) {
		//fmt.Println(t)
	}
	temps, err := NewStatusMonitor(publishTemp)
	CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)
	err = temps.TempMonitor.CompileRegex()
	CommonTestTools.CheckErr(t, "TestTemperatureMonitor", err)

	// Add a bunch of temps
	for i := 0; i < 200; i++ {
		temps.TempMonitor.UpdateTemperature(fmt.Sprintf("T:%v/%v B:%v / %v", i, i+1, i+2, i+3))
		<-time.After(10 * time.Millisecond)
	}

	history := temps.TempMonitor.GetTempHistory()
	if len(history) != 100 {
		t.Fatal("Length Does not equal 100 Actual:", len(history))
	}

	for _, line := range history {
		fmt.Printf("Target: %v Actual: %v Time %v\n", line.Target["T"], line.Temp["T"], line.Time.Unix())
	}
}
