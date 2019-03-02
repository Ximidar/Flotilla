/*
* @Author: Ximidar
* @Date:   2019-03-01 15:05:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 15:23:50
 */

package Printer

import (
	"fmt"
	"testing"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
)

func TestLineChecker(t *testing.T) {
	printer, _ := NewPrinter()

	line := "N0 G1 X1 X2"
	expected := "G1 X1 X2"
	checksum := printer.GetChecksum([]byte(line))

	line = fmt.Sprintf("%v*%v\n", line, checksum)

	rawline, err := printer.lineChecker([]byte(line))
	CommonTestTools.CheckErr(t, "Did not Check line correctly", err)

	if string(rawline) != expected {
		t.Fatal("Line does not match the rawline", expected, string(rawline), []byte(expected), rawline)
	}

}
