/*
* @Author: Ximidar
* @Date:   2018-12-22 16:33:09
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 20:27:16
 */

package CommRelay

import (
	"fmt"
	"strings"
)

// FormattedLine is a struct for processed lines
type FormattedLine struct {
	FormattedLine string
	LineNumber    uint64
}

// GetFormattedLine will output the formatted line with a correct offset
func (fl *FormattedLine) GetFormattedLine(offset uint64) string {
	// Format the line
	currentline := offset + fl.LineNumber
	formLine := fmt.Sprintf("N%v %v", currentline, fl.FormattedLine)
	formLine = strings.Replace(formLine, "\n", "", -1)
	formLine = strings.Replace(formLine, "\r", "", -1)
	checksum := fl.GetChecksum(formLine)
	formLine = fmt.Sprintf("%v*%v", formLine, checksum)

	return formLine
}

// GetOverwrittenFormattedLine will create a formatted line with the overwritten line on it
func (fl *FormattedLine) GetOverwrittenFormattedLine(lineNumber uint64) string {
	// Format the line
	formLine := fmt.Sprintf("N%v %v", lineNumber, fl.FormattedLine)
	formLine = strings.Replace(formLine, "\n", "", -1)
	checksum := fl.GetChecksum(formLine)
	formLine = fmt.Sprintf("%v*%v", formLine, checksum)

	return formLine
}

// GetChecksum will calculate the checksum for the line
func (fl *FormattedLine) GetChecksum(line string) int {
	byteslice := []byte(line)
	checksum := 0
	for _, val := range byteslice {
		checksum ^= int(val)
	}
	return checksum
}

// ArrangedLines is the sorting method for an array of FormattedLine(s)
type ArrangedLines []FormattedLine

func (al ArrangedLines) Len() int {
	return len(al)
}
func (al ArrangedLines) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}
func (al ArrangedLines) Less(i, j int) bool {
	return al[i].LineNumber < al[j].LineNumber
}
