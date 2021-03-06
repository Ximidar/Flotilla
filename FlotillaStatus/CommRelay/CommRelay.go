/*
* @Author: Ximidar
* @Date:   2018-12-17 10:31:55
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-05-07 20:19:23
 */

package CommRelay

import (
	"fmt"
	"strings"
	"sync"

	CRS "github.com/Ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
)

// LineCallback is a function that other packages provide to CommRelay
type LineCallback func(line string) error

// FillBuffer will ask to fill the buffer
type FillBuffer func()

// FinishedStreamCallback is for telling the upper division that we have received the last
// line of the stream
type FinishedStreamCallback func()

// CommRelay will take in all commands meant for Commango then organize them
type CommRelay struct {
	Offset        uint64 // Offset from the read line
	CurrentLine   uint64 // The current read line
	bufferedLines []*CRS.Line

	FinalLineBuffer *RollingFormattedLine
	CurrentReadLine uint64

	LineCallback LineCallback
	FillBuffer   FillBuffer
	SkipNextOK   bool
	LineParser   *LineParser
	SaveLines    map[uint64]FormattedLine
	mux          sync.Mutex

	Playing            bool
	NotifyWhenFinished bool
	Finished           bool
	FinishedStream     func()

	// Event Driven Channels
	IncomingLines chan FormattedLine
	WriteToComm   chan FormattedLine
	OKEvent       chan bool
	WaitEvent     chan bool
	ResendEvent   chan int
	Consume       chan bool
}

// NewCommRelay will return a new Comm Object
func NewCommRelay(lineCallbackFunction LineCallback,
	fillbuffer FillBuffer,
	finishedStreamCallback FinishedStreamCallback) (*CommRelay, error) {

	comm := new(CommRelay)
	// reset all lines
	comm.makechannels()
	comm.ResetLines()

	// Assign Callbacks
	comm.LineCallback = lineCallbackFunction
	comm.FillBuffer = fillbuffer
	comm.FinishedStream = finishedStreamCallback

	// Make the line parser
	var err error
	comm.LineParser, err = NewLineParser(comm.ResendEvent, comm.OKEvent, comm.WaitEvent)
	if err != nil {
		return nil, err
	}

	// Start up the event handler
	return comm, err
}

// StartEventHandler will start the event handler in a new goroutine
func (CR *CommRelay) StartEventHandler() {
	fmt.Println("Starting event handler")
	go CR.EventHandler()
	go CR.CommInputHandler()
}

func (CR *CommRelay) makechannels() {
	CR.IncomingLines = make(chan FormattedLine, 1000)
	CR.WriteToComm = make(chan FormattedLine, 1000)
	CR.OKEvent = make(chan bool, 100)
	CR.WaitEvent = make(chan bool, 100)
	CR.ResendEvent = make(chan int, 100)
}

// FormatLine will intake a line and then be transformed into the correct format
func (CR *CommRelay) FormatLine(line *CRS.Line) {
	var formatted FormattedLine
	// If the line number isn't known add it to the line buffer
	if !line.GetKnownNumber() {
		formatted = FormattedLine{
			FormattedLine: line.GetLine(),
			LineNumber:    0,
		}
	} else {
		formatted = FormattedLine{
			FormattedLine: line.GetLine(),
			LineNumber:    line.GetLineNumber(),
		}
	}

	CR.IncomingLines <- formatted

}

// ResetLines will reset the line number to 0
func (CR *CommRelay) ResetLines() {
	fmt.Println("Reseting Comm Relay")
	CR.Offset = 0
	CR.CurrentLine = 0
	CR.FinalLineBuffer = NewRollingFormattedLine(1000)
	CR.CurrentReadLine = 0
	CR.SkipNextOK = false
	CR.Playing = false
	CR.NotifyWhenFinished = false
	CR.Finished = false
	CR.SaveLines = make(map[uint64]FormattedLine)
}

// ConsumeLine will send the next line when an OK signal is received
func (CR *CommRelay) ConsumeLine() error {
	// If there's nothing to give we cannot give it
	line, err := CR.FinalLineBuffer.GetLine(CR.CurrentReadLine)
	if err != nil {
		fmt.Println(err)
		// If we have the flag to notify when finished and we cannot grab the next line
		// Then end the print.
		if CR.NotifyWhenFinished && !CR.Finished {
			if CR.CurrentReadLine != CR.FinalLineBuffer.CurrentRead {
				return err
			}
			fmt.Println("Finished Printing!")
			CR.Finished = true
			CR.ResetLines()
			CR.FinishedStream()
			return nil
		}

		// try to look for the line in the save lines dictionary
		CR.mux.Lock()
		line, ok := CR.SaveLines[CR.CurrentReadLine]
		CR.mux.Unlock()
		if !ok {
			return err
		}
		CR.WriteToComm <- line
		return nil

	}
	CR.WriteToComm <- line
	return nil
}

func (CR *CommRelay) sendLine(line string) {
	noNewLine := strings.Replace(line, "\n", "", -1)
	if CR.LineCallback != nil {

		err := CR.LineCallback(noNewLine)
		if err != nil {
			fmt.Print("\n", err, "\n")
		}
	}
}

// ResendLine will be shared with another process to detect the need for resend lines
func (CR *CommRelay) ResendLine(lineNum int) {
	fmt.Println("Rewinding Next Read line to:", lineNum)
	CR.CurrentReadLine = uint64(lineNum)
	CR.ConsumeLine()

}

// RecieveComm will check for wait, ok, and such then send a signal to a callback
func (CR *CommRelay) RecieveComm(line string) {
	if !CR.Playing {
		//fmt.Println("Not playing so not sending line")
		return
	}
	CR.LineParser.ProcessLine(line)

}

// SaveLine will save lines to a buffer to look up when communication goes wrong
func (CR *CommRelay) SaveLine(line FormattedLine, lineNum uint64) {
	CR.mux.Lock()
	CR.SaveLines[lineNum] = line
	CR.mux.Unlock()
	go CR.cleanSaveLines(lineNum)
}

func (CR *CommRelay) cleanSaveLines(lineNum uint64) {
	lineNum += 1000

	for {
		_, ok := CR.SaveLines[lineNum]
		if ok {
			CR.mux.Lock()
			delete(CR.SaveLines, lineNum)
			CR.mux.Unlock()
			lineNum++
		} else {
			break
		}
	}
}

// NotifyWhenEmpty will set a condition to callback to FinishedStream when
// The buffer empties
func (CR *CommRelay) NotifyWhenEmpty() {
	CR.NotifyWhenFinished = true
}

//CommInputHandler will handle signals like OK, WAIT, and RESEND
func (CR *CommRelay) CommInputHandler() {
	for {
		select {
		case <-CR.OKEvent:
			// Consume Line as OK comes in
			if err := CR.ConsumeLine(); err != nil {
				fmt.Println("Could not consume line", err)

			}

			if CR.FinalLineBuffer.Filled() < 25 {
				fmt.Println("Under 25%", CR.FinalLineBuffer.Filled())
				go CR.FillBuffer()
			}
		case <-CR.WaitEvent:
			CR.OKEvent <- false
			// CR.debugChannels()
		case resendLineNumber := <-CR.ResendEvent:
			fmt.Println("Resend Event")
			CR.CurrentReadLine = uint64(resendLineNumber)
		}
	}
}

// EventHandler will direct all incoming events
func (CR *CommRelay) EventHandler() {
	for {
		select {
		case line := <-CR.IncomingLines:
			line.LineNumber = CR.CurrentLine
			CR.CurrentLine++
			CR.FinalLineBuffer.AppendLine(line)

		case line := <-CR.WriteToComm:
			sendLine := line.GetFormattedLine(0)
			CR.sendLine(sendLine)

			//save line for later lookup
			go CR.SaveLine(line, CR.CurrentReadLine)

			// up latest line to read
			CR.CurrentReadLine++

		}
	}
}

// Debug the channels to see what is filling up
func (CR *CommRelay) debugChannels() {
	// CR.IncomingLines = make(chan FormattedLine, 10000)
	// CR.WriteToComm = make(chan FormattedLine, 10000)
	// CR.OKEvent = make(chan bool, 100)
	// CR.WaitEvent = make(chan bool, 100)
	// CR.ResendEvent = make(chan int, 100)
	fmt.Printf("IncomingLines CAP %v LEN %v\n", cap(CR.IncomingLines), len(CR.IncomingLines))
	fmt.Printf("CR.WriteToComm CAP %v LEN %v\n", cap(CR.WriteToComm), len(CR.WriteToComm))
	fmt.Printf("CR.OKEvent CAP %v LEN %v\n", cap(CR.OKEvent), len(CR.OKEvent))
	fmt.Printf("CR.WaitEvent CAP %v LEN %v\n", cap(CR.WaitEvent), len(CR.WaitEvent))
	fmt.Printf("CR.ResendEvent CAP %v LEN %v\n", cap(CR.ResendEvent), len(CR.ResendEvent))

}
