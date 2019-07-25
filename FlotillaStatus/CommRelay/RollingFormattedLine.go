/*
* @Author: Ximidar
* @Date:   2018-12-28 19:40:11
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 17:07:15
 */

package CommRelay

import (
	"errors"
	"fmt"
	"sync"
)

var (
	//ErrLineNotKnown will be returned if the element in Slice is not there.
	ErrLineNotKnown = errors.New("line not known")
	// ErrLineNotPopulated will be returned if the line is not actually populated yet
	ErrLineNotPopulated = errors.New("line not populated")
)

// RollingFormattedLine will only keep the last MaxSize of FormattedLine as
// new lines are added the onld ones will be deleted.
type RollingFormattedLine struct {
	Slice       []FormattedLine
	MaxSize     uint64
	CurrentRead uint64
	offset      uint64
	update      chan bool
	mux         sync.Mutex
}

// NewRollingFormattedLine will create a RollingFormattedLine struct
func NewRollingFormattedLine(size uint64) *RollingFormattedLine {
	rs := new(RollingFormattedLine)
	rs.MaxSize = size
	rs.Slice = make([]FormattedLine, 0, rs.MaxSize)
	rs.CurrentRead = 0
	rs.offset = 0
	rs.update = make(chan bool)
	return rs
}

// AppendLine will append a FormattedLine to the Slice, If the slice is at capacity it will
// shift all values down and add to the end
func (rs *RollingFormattedLine) AppendLine(line FormattedLine) {
	rs.mux.Lock()
	defer rs.mux.Unlock()
	caps := cap(rs.Slice)

	if len(rs.Slice) < caps {

		rs.Slice = append(rs.Slice, line)
	} else {
		// Shift down and put in last place
		rs.Slice = rs.Slice[:copy(rs.Slice[0:], rs.Slice[1:])]
		rs.Slice = append(rs.Slice, line)
		rs.offset++
	}
	rs.CurrentRead++
}

// GetLine will get the asked for line in the history.
func (rs *RollingFormattedLine) GetLine(lineNum uint64) (FormattedLine, error) {

	rs.mux.Lock()
	defer rs.mux.Unlock()

	for x := len(rs.Slice) - 1; x >= 0; x-- {
		if rs.Slice[x].LineNumber == lineNum {
			temp := rs.Slice[x]
			rs.Slice = append(rs.Slice[:x], rs.Slice[x+1:]...)
			return temp, nil
		}
	}
	fmt.Println("Could not find line", lineNum)
	return FormattedLine{}, ErrLineNotKnown

}

// GetMostRecentLine will get the most recent added line
func (rs *RollingFormattedLine) GetMostRecentLine() (FormattedLine, error) {
	mrl, err := rs.GetLine(rs.CurrentRead - 1)
	return mrl, err
}

// Filled will return a percentage that the slice is full to
func (rs *RollingFormattedLine) Filled() int {
	fmt.Println(float64(len(rs.Slice)), " / ", float64(rs.MaxSize), " * 100.00")

	return int((float64(len(rs.Slice)) / float64(rs.MaxSize)) * 100.00)
}

// FilledTo will return a boolean that signifies if the buffer is at least N% full
func (rs *RollingFormattedLine) FilledTo(filledpercent int) bool {
	return float64(len(rs.Slice)) >= float64(rs.MaxSize)*(float64(filledpercent)/100.00)
}
