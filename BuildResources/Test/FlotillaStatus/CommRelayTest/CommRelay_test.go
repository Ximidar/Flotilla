/*
* @Author: Ximidar
* @Date:   2018-12-18 11:39:47
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-29 01:36:57
 */

package CommRelayTest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	"github.com/ximidar/Flotilla/FlotillaStatus/CommRelay"
)

// TestCommRelaySetup tests if NewCommRelay will return an error
func TestCommRelaySetup(t *testing.T) {
	callback := func(line string) error {
		fmt.Println(line)
		return nil
	}
	callbackok := func() {
		fmt.Println("Got OK")
	}
	finishcall := func() {
		fmt.Println("Got Finish")
	}
	_, err := CommRelay.NewCommRelay(callback, callbackok, finishcall)
	CommonTestTools.CheckErr(t, "TestCommRelay setup failed", err)
}

func TestCommRelayLines(t *testing.T) {
	readLines := 0
	callback := func(line string) error {
		readLines++
		fmt.Println(line)
		return nil
	}

	callbackok := func() {
		fmt.Println("Got OK")
	}
	finishcall := func() {
		fmt.Println("Got Finish")
	}

	commRelay, err := CommRelay.NewCommRelay(callback, callbackok, finishcall)
	CommonTestTools.CheckErr(t, "TestCommRelayLines setup failed", err)

	// Reset Lines to zero
	commRelay.ResetLines()

	// test sending a line and see if it parses correctly
	line0, err := CommRelayStructures.NewLine("G1 X0 Y0", 0, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	commRelay.FormatLine(line0)
	commRelay.ConsumeLine()

	<-time.After(50 * time.Millisecond)
	if readLines != 1 {
		t.Fatal("TestCommRelayLines Parse line did not parse a line")
	}

	// Send lines out of order and see if parse line will pick up each line

	line1, err := CommRelayStructures.NewLine("Line 1", 1, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line2, err := CommRelayStructures.NewLine("Line 2", 2, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line3, err := CommRelayStructures.NewLine("Line 3", 3, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line4, err := CommRelayStructures.NewLine("Line 4", 4, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line5, err := CommRelayStructures.NewLine("Line 5", 5, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	commRelay.FormatLine(line5)
	commRelay.FormatLine(line4)
	commRelay.FormatLine(line3)
	commRelay.FormatLine(line2)
	commRelay.FormatLine(line1)

	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()

	<-time.After(75 * time.Millisecond)
	if readLines != 6 {
		t.Fatal("TestCommRelayLines Did not read all lines")
	}

	// Test Reading out of order and adding in a few  undocumented lines
	commRelay.ResetLines()
	readLines = 0

	line6, err := CommRelayStructures.NewLine("Line Uknown 0", 0, false)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line7, err := CommRelayStructures.NewLine("Line Uknown 1", 0, false)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	line8, err := CommRelayStructures.NewLine("Line Uknown 2", 0, false)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	commRelay.FormatLine(line0)
	commRelay.FormatLine(line1)
	commRelay.FormatLine(line6)
	commRelay.FormatLine(line2)
	commRelay.FormatLine(line3)
	commRelay.FormatLine(line7)
	commRelay.FormatLine(line5)
	commRelay.FormatLine(line8)
	commRelay.FormatLine(line4)

	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()
	commRelay.ConsumeLine()

	fmt.Println(commRelay.LineBuffer)

	<-time.After(150 * time.Millisecond)
	if readLines != 9 {
		t.Fatal("TestCommRelayLines Did not read all lines")
	}

}

func TestLineParser(t *testing.T) {
	expectedLine := 0
	resendCallback := func(line int) {
		fmt.Printf("\nResend Line: %v, Expected: %v\n", line, expectedLine)
		if expectedLine != line {
			t.Fatal("Line was not expected line. line", line, "expected", expectedLine)
		}
	}
	lp, err := CommRelay.NewLineParser(resendCallback)
	if err != nil {
		CommonTestTools.CheckErr(t, "TestLineParser construction failed", err)
	}

	var line string

	// Try out a few different expected lines
	IntSlice := makeRandIntSlice()

	for _, val := range IntSlice {
		var num int
		num = int(val)
		expectedLine = num
		line = fmt.Sprintf("Resend: %v", num)
		lp.ProcessLine(line)
	}

	// Try out a few ok lines
	lines := []string{"ok", "no", "not even right at all", "ok", "     ok       ", "wait"}
	lineAnswer := []bool{true, false, false, true, true, true}

	for index, val := range lines {
		expected := lineAnswer[index]
		received := lp.OkRegex.MatchString(val)
		if expected != received {
			t.Fatal("answer is not correct Expected:", expected, "Received:", received, "Val:", val)
		}
	}

}

func makeRandIntSlice() []int {
	IntSlice := make([]int, 50)
	for index := range IntSlice {
		IntSlice[index] = rand.Intn(1000000)
	}
	return IntSlice
}
