/*
* @Author: Ximidar
* @Date:   2019-03-01 20:22:21
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 17:11:34
 */

package CommRelay

import (
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
)

func lineCaller(line string) error {
	fmt.Println("Got line:", line)
	return nil
}

func askforLine(lineAmount int) {
	fmt.Println(lineAmount, "Lines Were Asked For")
}

func finished() {
	fmt.Println("Comm Is Finished")
}

func TestCommRelayEvents(t *testing.T) {

	comm, err := NewCommRelay(lineCaller, askforLine, finished)
	CommonTestTools.CheckErr(t, "Could not create Comm Relay", err)

	comm.Playing = true

	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 0,
	})
	select {
	case line := <-comm.IncomingLines:
		fmt.Println("Alright", line.FormattedLine)
	case <-time.After(2 * time.Second):
		t.Fatal("Could not get event for an incomming line")
	}

	// Test parsing the incoming comm line
	comm.RecieveComm("ok\n")
	select {
	case <-comm.OKEvent:
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not get OK Event")
	}
	// Test parsing the incoming comm line
	comm.RecieveComm("wait\n")
	select {
	case <-comm.WaitEvent:
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not get Wait Event")
	}
	// Test parsing the incoming comm line
	comm.RecieveComm("Resend: 10\n")
	select {
	case ResendNum := <-comm.ResendEvent:
		fmt.Println("Alright", ResendNum)
	case <-time.After(2 * time.Second):
		t.Fatal("Could not get Resend Event")
	}

	// Test the write to comm
	comm.WriteToComm <- FormattedLine{
		LineNumber:    10,
		FormattedLine: "Aloha",
	}
	select {
	case line := <-comm.WriteToComm:
		fmt.Println("Alright", line.FormattedLine, line.LineNumber)
	case <-time.After(2 * time.Second):
		t.Fatal("Could not get Resend Event")
	}

	comm.ResetLines()
	comm.makechannels()
	comm.Playing = true

	comm.StartEventHandler()
	<-time.After(2 * time.Second)

	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 0,
	})
	<-time.After(100 * time.Millisecond)
	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 1,
	})
	<-time.After(100 * time.Millisecond)
	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 2,
	})
	<-time.After(100 * time.Millisecond)
	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 3,
	})
	<-time.After(100 * time.Millisecond)
	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 4,
	})
	<-time.After(100 * time.Millisecond)
	comm.FormatLine(&CommRelayStructures.Line{
		Line:       "Hello",
		LineNumber: 5,
	})
	<-time.After(100 * time.Millisecond)

	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)
	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)
	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)
	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)
	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)
	comm.RecieveComm("ok\n")
	<-time.After(100 * time.Millisecond)

	<-time.After(2 * time.Second)

}

// TestCommRelaySetup tests if NewCommRelay will return an error
func TestCommRelaySetup(t *testing.T) {
	callback := func(line string) error {
		fmt.Println(line)
		return nil
	}
	askforline := func(lineAmount int) {
		fmt.Println("Asked for", lineAmount, "Lines")
	}
	finishcall := func() {
		fmt.Println("Got Finish")
	}
	_, err := NewCommRelay(callback, askforline, finishcall)
	CommonTestTools.CheckErr(t, "TestCommRelay setup failed", err)
}

func TestCommRelayLines(t *testing.T) {
	readLines := 0
	callback := func(line string) error {
		readLines++
		fmt.Println(line)
		return nil
	}

	askforline := func(lineAmount int) {
		fmt.Println("Asked for", lineAmount, "Lines")
	}
	finishcall := func() {
		fmt.Println("Got Finish")
	}

	commRelay, err := NewCommRelay(callback, askforline, finishcall)
	CommonTestTools.CheckErr(t, "TestCommRelayLines setup failed", err)
	<-time.After(75 * time.Millisecond)

	// test sending a line and see if it parses correctly
	line0, err := CommRelayStructures.NewLine("First Line Tester!", 0, true)
	CommonTestTools.CheckErr(t, "TestCommRelayLines newline failed", err)

	commRelay.FormatLine(line0)
	<-time.After(100 * time.Millisecond)
	commRelay.RecieveComm("ok\n")
	err = commRelay.ConsumeLine()
	CommonTestTools.CheckErr(t, "Couldn't consume the line", err)
	<-time.After(75 * time.Millisecond)
	if readLines != 1 {
		fmt.Println("readLines is at", readLines)
		t.Fatal("TestCommRelayLines Parse line did not parse a line")
		return
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

	<-time.After(150 * time.Millisecond)
	if readLines != 9 {
		t.Fatal("TestCommRelayLines Did not read all lines")
	}

}
