/*
* @Author: matt
* @Date:   2018-05-25 15:58:30
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 14:04:36
 */

package commango

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	"go.bug.st/serial" //https://godoc.org/go.bug.st/serial
)

// NatsAdapter adapter will be used by other packages when they need to emit statuses
type NatsAdapter interface {
	ReadLineEmitter(line string) error
	WriteLineEmitter(line string) error
	PublishStatus(status *CS.CommStatus) error
}

// Comm handles the Serial Connection
type Comm struct {
	options        *serial.Mode
	AvailablePorts *CS.Ports
	PortPath       string
	Port           serial.Port
	connected      bool
	ReadStream     chan string
	ByteStream     chan byte
	ErrorStream    chan error

	adapter         NatsAdapter
	finishedReading bool
}

// NewComm will Construct a Comm Object
func NewComm(adapter NatsAdapter) *Comm {
	comm := new(Comm)
	comm.PortPath = "None"
	comm.connected = false
	comm.adapter = adapter
	comm.options = &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	comm.ReadStream = make(chan string, 10)
	comm.ByteStream = make(chan byte, 100)
	comm.ErrorStream = make(chan error, 10)
	return comm
}

// Connected is a function to check the current connection to the serial port
func (comm *Comm) Connected() bool {
	return comm.connected
}

// SetConnected will set the connection status
func (comm *Comm) SetConnected(value bool) {
	comm.connected = value
	status := comm.GetCommStatus()
	err := comm.adapter.PublishStatus(status)
	if err != nil {
		fmt.Println("Could not set Connected", err)
	}
}

// InitComm will initialize the comm object with a defined Serial Connection
func (comm *Comm) InitComm(init *CS.InitComm) error {

	comm.PortPath = init.GetPort()
	comm.options = &serial.Mode{
		BaudRate: int(init.GetBaud()),
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	comm.PrintOptions()
	return nil
}

// PrintOptions will print out the current connection options
func (comm Comm) PrintOptions() {
	fmt.Println("Comm Options:")
	fmt.Println("|  Port Path:", comm.PortPath)
	fmt.Println("|  Serial Options:")
	fmt.Println("|  |  Baud Rate:", comm.options.BaudRate)
	fmt.Println("|  |  Parity:", comm.options.Parity)
	fmt.Println("|  |  Data Bits:", comm.options.DataBits)
	fmt.Println("|  |  Stop Bits:", comm.options.StopBits)
}

// GetCommStatus will return the current connection status
func (comm Comm) GetCommStatus() *CS.CommStatus {
	cs := new(CS.CommStatus)
	cs.Port = comm.PortPath
	cs.Baud = int32(comm.options.BaudRate)
	cs.Connected = comm.Connected()

	return cs
}

// GetAvailablePorts will query the system for all available ports we can connect to
func (comm *Comm) GetAvailablePorts() (*CS.Ports, error) {
	// Reset the available ports
	comm.AvailablePorts = new(CS.Ports)

	portlist, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}
	// insert fake printer into spot 0
	portlist = append(portlist, "")
	copy(portlist[0+1:], portlist[0:])
	portlist[0] = "/dev/fakeprinter"

	for _, p := range portlist {
		port := new(CS.Port)
		port.Address = p

		comm.AvailablePorts.Ports = append(comm.AvailablePorts.Ports, port)
	}

	return comm.AvailablePorts, nil
}

// PreCheck Do a little error checking for good start
func (comm Comm) PreCheck() (ready bool) {
	ready = true
	if comm.PortPath == "" {
		ready = false
	}
	return
}

// This initiates a reset for most microcontrollers
func (comm *Comm) cycledtr() {
	comm.Port.SetDTR(false)
	time.Sleep(100 * time.Millisecond)
	comm.Port.SetDTR(true)
	time.Sleep(100 * time.Millisecond)
	comm.Port.SetDTR(false)
}

// OpenComm will attempt to open a Serial Connection
func (comm *Comm) OpenComm() error {

	// recover in case this all goes tits up
	defer func() {
		if recovery := recover(); recovery != nil {
			fmt.Println("Could not Open Port.", recovery)
			comm.SetConnected(false)
		}
	}()

	// Do a Precheck before starting
	if !comm.PreCheck() {
		fmt.Println("Precheck Failed!")
		return errors.New("precheck Failed, Please initialize the comm before trying to open it")
	}

	if comm.Connected() {
		fmt.Println("Comm is already connected")
		return errors.New("Comm is already Connected")
	}

	var err error
	fmt.Printf("Attempting to open port with address %v\n", comm.PortPath)
	comm.Port, err = serial.Open(comm.PortPath, comm.options)
	if err != nil {
		fmt.Println("Error Could not open port", err)
		return err
	}
	comm.cycledtr()
	// Sleep to allow the port to start up
	time.Sleep(20 * time.Millisecond)
	comm.SetConnected(true)
	fmt.Println("Port Opened")

	// Start up a reader
	//go comm.Read_Forever()
	go comm.ReadOKForever()
	return nil
}

// CloseComm will attempt to close the Serial Connection
func (comm *Comm) CloseComm() error {
	// recover in case this all goes tits up
	defer func() {
		if recovery := recover(); recovery != nil {
			fmt.Println("Could not close port.", recovery)
		}
	}()

	if comm.Connected() {
		fmt.Printf("Attempting to close port with address %s\n", comm.PortPath)
		err := comm.Port.Close()
		if err != nil {
			fmt.Println("Could not close port")
			return err
		}
		fmt.Println("Port Closed")
	}

	comm.SetConnected(false)
	return nil
}

// WriteComm will write a message to the Serial Connection
func (comm *Comm) WriteComm(message string) (lenWritten int, err error) {
	lenWritten = -1
	//prepare string for writing
	// Check that message has an \n after it
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	// Turn message into a bytestring
	byteMessage := []byte(message)
	expectedWrite := len(message)

	// Write the comm if we are connected
	if comm.Connected() {
		lenWritten, err = comm.Port.Write(byteMessage)

		if err != nil {
			return
		}
		if lenWritten != expectedWrite {
			fmt.Println("Didn't write expected amount of bytes")
			fmt.Printf("Written: %v Expected: %v", lenWritten, expectedWrite)
			comm.adapter.WriteLineEmitter(message + " Error on this line")
			return lenWritten, errors.New("expected Bytes did not match written")
		}
	} else {
		err = errors.New("Cannot Write Comm")
		return
	}

	comm.adapter.WriteLineEmitter(message)
	return
}

// ReadOK This function will read the serial port until it receives ok
func (comm *Comm) ReadOK() {
	var buf []byte
	for comm.connected {
		select {
		case read := <-comm.ByteStream:
			//add the bytes to the buffer
			buf = append(buf, read)
			if len(buf) == 0 {
				continue
			}

			if comm.CheckForEndline(read) {
				comm.ReadStream <- string(buf)
				buf = []byte{}
			}
		case <-comm.ErrorStream:
			return //If we are erroring out then we can't read anything
		}
	}
}

// CheckForOK Check for ok or start
func (comm *Comm) CheckForOK(buf []byte) bool {
	bufString := string(buf)

	acceptableChecks := []string{"ok", "start", "wait", "echo:busy: processing"}

	for index := range acceptableChecks {
		if strings.Contains(bufString, acceptableChecks[index]) {
			return true
		}
	}
	return false
}

// CheckForEndline will check if the line ends with \n
func (comm *Comm) CheckForEndline(buf byte) bool {
	if buf == 10 {
		return true
	}
	return false
}

// StreamBytes This function will continuously read the output from the Serial line
// It will send these bytes through a channel.
func (comm *Comm) StreamBytes() {
	for comm.connected {
		read, err := comm.ReadBytes(1)
		if err != nil {
			comm.ErrorStream <- err
			fmt.Println("Stream Bytes Erroring out")
			return
		}
		comm.ByteStream <- read[0]
	}
}

// ReadLine will Read until a newline gets read
func (comm *Comm) ReadLine() (out []byte, err error) {
	readLine := false
	for readLine == false {
		read, err := comm.ReadBytes(1)

		if err != nil {
			//fmt.Println("Readline errored out", err)
			return nil, err
		}

		if read[0] == 10 {
			out = append(out, read[0])
			readLine = true
			return out, nil
		}

		out = append(out, read[0])

	}

	return

}

// ReadBytes will read n amount of bytes from the Serial Line
func (comm *Comm) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	bytesRead, err := comm.Port.Read(buf)
	if bytesRead != n {
		fmt.Printf("Read: %v Expexted: %v\n", bytesRead, n)
		err = fmt.Errorf("Read: %v Expexted: %v", bytesRead, n)
	}
	return buf, err
}

// ReadForever will Read the serial line until there is an error or until until the program exits
func (comm *Comm) ReadForever() {

	for comm.connected {
		out, err := comm.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("%v", err)
			}
		} else {
			stringOut := string(out)
			comm.adapter.ReadLineEmitter(stringOut)
		}
	}
	fmt.Println("Stopping the reading")
}

// ReadOKForever will Read messages seperated by an OK
func (comm *Comm) ReadOKForever() (err error) {
	go comm.StreamBytes()
	go comm.ReadOK()

	for comm.connected {
		select {
		case fullString := <-comm.ReadStream:
			comm.adapter.ReadLineEmitter(fullString)
		case err = <-comm.ErrorStream:
			fmt.Println("Read Ok Erroring out:", err.Error())
			return err
		}
	}

	fmt.Println("Stopped Reading")
	return nil

}
