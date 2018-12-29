/*
* @Author: Ximidar
* @Date:   2018-12-28 20:05:10
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-29 01:25:27
 */

package CommRelayTest

import (
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/FlotillaStatus/CommRelay"
)

func TestRollingFormattedLine(t *testing.T) {
	rofl := CommRelay.NewRollingFormattedLine(5)
	lines := makeLines(10)

	for _, line := range lines {
		rofl.AppendLine(line)
	}

	// Check line 6
	expectedLine, err := rofl.GetLine(6)
	CommonTestTools.CheckErr(t, "TestRollingFormattedLine", err)
	actualLine := lines[6]

	if expectedLine.LineNumber != actualLine.LineNumber {
		t.Fatal("Lines Dont Match", expectedLine.FormattedLine, actualLine.FormattedLine)
	}

	// Check most recent line
	expectedLine, err = rofl.GetMostRecentLine()
	CommonTestTools.CheckErr(t, "TestRollingFormattedLine", err)
	actualLine = lines[len(lines)-1]

	if expectedLine.LineNumber != actualLine.LineNumber {
		t.Fatal("Lines Dont Match", expectedLine.FormattedLine, actualLine.FormattedLine)
	}

	fmt.Println(lines)
	fmt.Println(rofl.Slice)
}

func TestRollingFormattedLine150kLines(t *testing.T) {
	expectedSize := 200
	rofl := CommRelay.NewRollingFormattedLine(uint64(expectedSize))
	lines := makeLines(150000)

	// Test without goroutines
	lineCounter := 0
	for _, line := range lines {
		rofl.AppendLine(line)
		lineCounter++
		if lineCounter == 10 {
			lineCounter = 0
			size := len(rofl.Slice)
			if size != expectedSize {
				t.Fatal("Size is not right. Expected:", expectedSize, "Actual:", size)
			}
		}

	}

}
func TestRollingFormattedLine150kLinesGoroutines(t *testing.T) {
	expectedSize := 200
	rofl := CommRelay.NewRollingFormattedLine(uint64(expectedSize))
	exitSignal := make(chan bool, 10)

	dumpLines := func() {
		lines := makeLines(150000)
		for _, line := range lines {
			rofl.AppendLine(line)
			<-time.After(5 * time.Microsecond)
		}
		exitSignal <- true
	}

	checkEvery20Milli := func() {
		for {
			<-time.After(20 * time.Millisecond)
			size := len(rofl.Slice)
			if size != expectedSize {
				t.Fatal("Size is not right. Expected:", expectedSize, "Actual:", size)
			}

			select {
			case <-exitSignal:
				return
			default:
				continue
			}

		}

	}

	grabLineEvery20Milli := func() {
		for {
			<-time.After(20 * time.Millisecond)
			if !(rofl.CurrentRead > 5) {
				continue
			}
			grabLine := rofl.CurrentRead - 5
			line, err := rofl.GetLine(grabLine)
			CommonTestTools.CheckErr(t, "TestRollingFormattedLine150kLinesGoroutines", err)
			if line.LineNumber != grabLine {
				err := fmt.Errorf("line and grabline dont match expected: %v actual %v", grabLine, line.LineNumber)
				CommonTestTools.CheckErr(t, "TestRollingFormattedLine150kLinesGoroutines", err)
			}

			select {
			case <-exitSignal:
				return
			default:
				continue
			}
		}
	}

	// Test with goroutines
	go dumpLines()
	go checkEvery20Milli()
	go grabLineEvery20Milli()
	<-exitSignal

}

func makeLines(amount int) []CommRelay.FormattedLine {

	lines := []CommRelay.FormattedLine{}
	for index := 0; index < amount; index++ {
		line := CommRelay.FormattedLine{FormattedLine: fmt.Sprintf("Line: %v", index), LineNumber: uint64(index)}
		lines = append(lines, line)
	}
	return lines

}
