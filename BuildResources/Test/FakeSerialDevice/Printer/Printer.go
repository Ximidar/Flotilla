/*
* @Author: Ximidar
* @Date:   2019-03-01 14:09:12
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 15:54:27
 */

package Printer

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"

	FakeSerialDevice "github.com/ximidar/Flotilla/BuildResources/Test/FakeSerialDevice/SerialDevice"
)

var OK = "ok\n"
var WAIT = "wait\n"

type position struct {
	X int
	Y int
	Z int
	E int
}

type temperature struct {
	tool   string
	set    int
	actual int
}

// Printer will act as a basic marlin printer by doing more than replying ok
// Printer will respond with position updates, Temperature updates, and resend errors
type Printer struct {
	fsd *FakeSerialDevice.FakeSerial
	position
	temperature
	lineNumber int

	exitChan    chan bool
	lineIn      chan []byte
	responseOut chan []byte
}

// NewPrinter will construct a new printer
func NewPrinter() (*Printer, error) {
	printer := new(Printer)
	printer.lineNumber = -1

	// Make Channels for communication
	printer.exitChan = make(chan bool, 10)
	printer.lineIn = make(chan []byte, 100)
	printer.responseOut = make(chan []byte, 100)

	return printer, nil
}

func (printer *Printer) Run(fsd *FakeSerialDevice.FakeSerial, exit chan bool) {
	printer.fsd = fsd
	printer.exitChan = exit
	printer.CollectSerialInfo()
}

// CollectSerialInfo will collect info from the FSD and return responses to it.
func (printer *Printer) CollectSerialInfo() {
	var buffer []byte
	for {
		select {
		case buf := <-printer.fsd.ReceiveStream:
			buffer = append(buffer, buf)
			if buf == 10 { // if we detect a newline
				fmt.Println("RECV:", string(buffer[:len(buffer)-1]))
				printer.lineIn <- buffer
				buffer = []byte{}
			}
		case response := <-printer.responseOut:
			fmt.Println("SENT:", string(response))
			printer.fsd.SendBytes(response)
		case line := <-printer.lineIn:
			// Process Line
			printer.LineHandler(line)
		case <-time.After(2 * time.Second):
			printer.responseOut <- []byte(WAIT)
		case <-printer.exitChan:
			return
		}
	}
}

func (printer *Printer) temperatureSetter(line string) {

}

func (printer *Printer) positionSetter(line string) {

}

// lineChecker will check the line for the line number and the checksum.
// if the checksum does not match it will return an error
func (printer *Printer) lineChecker(line []byte) ([]byte, error) {
	// Check for line number
	if line[0] != byte('N') {
		return []byte(nil), errors.New("No line number")
	}

	firstSpace := bytes.Index(line, []byte(" "))
	asterisk := bytes.Index(line, []byte("*"))

	if firstSpace == -1 || asterisk == -1 {
		return []byte(nil), errors.New("cannot detect line")
	}

	rawline := line[:asterisk]
	rawline = bytes.TrimLeft(rawline, " ")
	rawline = bytes.TrimRight(rawline, "*")
	lineNum := line[:firstSpace]
	lineNum = bytes.Trim(lineNum, "N ")
	realNum, err := strconv.Atoi(string(lineNum))
	if err != nil {
		return []byte(nil), fmt.Errorf("Could not convert \"%v\" to integer", string(lineNum))
	}
	if realNum != printer.lineNumber+1 {
		return []byte(nil), errors.New("line number does not match last line")
	}

	checksum := line[asterisk:]
	checksum = bytes.Trim(checksum, "*\n")
	realChecksum, err := strconv.Atoi(string(checksum))
	if err != nil {
		return []byte(nil), fmt.Errorf("Could not convert \"%v\" to integer", string(checksum))
	}

	sumcheck := printer.GetChecksum(rawline)
	if sumcheck != realChecksum {
		return []byte(nil), fmt.Errorf("checksum does not match given: %v calculated %v Line: %v", realChecksum, sumcheck, rawline)
	}

	firstSpace = bytes.Index(rawline, []byte(" "))
	rawline = rawline[firstSpace:]
	rawline = bytes.TrimRight(rawline, " ")
	rawline = bytes.TrimLeft(rawline, " ")
	// increment the line number
	printer.lineNumber++
	return rawline, nil
}

// GetChecksum will calculate the checksum for the line
func (printer *Printer) GetChecksum(line []byte) int {
	checksum := 0
	for _, val := range line {
		checksum ^= int(val)
	}
	return checksum
}

func (printer *Printer) ReturnOK() {
	printer.responseOut <- []byte(OK)
}

// LineHandler will take in lines and figure them out
func (printer *Printer) LineHandler(line []byte) {
	defer printer.ReturnOK() // always respond OK to every command
	if len(line) == 0 {
		return
	}

	strippedLine, err := printer.lineChecker(line)
	if err != nil {
		// Resend error
		fmt.Println(err)
		resend := fmt.Sprintf("Resend: %v\n", printer.lineNumber+1)
		printer.responseOut <- []byte(resend)
		return
	}

	switch strippedLine[0] {
	case byte('G'):
		printer.positionSetter(string(line))
	case byte('M'):
		printer.temperatureSetter(string(line))
	default:
		fmt.Println("Unknown Command:", string(line))
	}
}
