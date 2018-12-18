/*
* @Author: Ximidar
* @Date:   2018-12-18 11:39:47
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-18 13:42:26
 */

package CommRelayTest

import (
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/CommRelayStructures"
	"github.com/ximidar/Flotilla/FlotillaStatus/CommRelay"
)

// TestCommRelaySetup tests if NewCommRelay will return an error
func TestCommRelaySetup(t *testing.T) {
	callback := func(line string) {
		fmt.Println(line)
	}
	_, err := CommRelay.NewCommRelay(callback)
	CommonTestTools.CheckErr(t, "TestCommRelay setup failed", err)
}

func TestCommRelayLines(t *testing.T) {
	readLines := 0
	callback := func(line string) {
		readLines++
		fmt.Println(line)
	}

	commRelay, err := CommRelay.NewCommRelay(callback)
	CommonTestTools.CheckErr(t, "TestCommRelayLines setup failed", err)

	// Reset Lines to zero
	commRelay.ResetLines()

	// test sending a line and see if it parses correctly
	line0, err := CommRelayStructures.NewLine("G1 X0 Y0", 0, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	commRelay.FormatLine(line0)

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

	fmt.Println(commRelay.LineBuffer)

	<-time.After(150 * time.Millisecond)
	if readLines != 9 {
		t.Fatal("TestCommRelayLines Did not read all lines")
	}

}
