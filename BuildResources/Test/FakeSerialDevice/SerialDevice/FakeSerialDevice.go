/*
* @Author: Ximidar
* @Date:   2018-12-12 14:33:07
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 16:03:20
 */

package FakeSerialDevice

import (
	"fmt"
	"io"
	"os"
	"syscall"

	"golang.org/x/sys/unix"

	"github.com/pkg/term/termios"
	fsEvents "github.com/tywkeene/go-fsevents"
)

// SerialState is an interface that can be passed to a register function for state events
type SerialState interface {
	SerialOpened()
	SerialClosed()
}

const (
	// SerialName is the name we are going to assign to our fake serial device
	SerialName = "/dev/fakeprinter"
)

// FakeSerial is an object that will emulate a serial device
type FakeSerial struct {
	ptyMaster   *os.File
	ptySlave    *os.File
	ptySettings *unix.Termios

	// Streams
	ReceiveStream chan byte
	SendStream    chan byte

	// Address
	Address string

	// Watcher
	watcher       *fsEvents.Watcher
	SlaveOpen     bool
	SlaveOpenChan chan bool
}

// NewFakeSerial will construct a new fake serial device
func NewFakeSerial() *FakeSerial {
	os.RemoveAll(SerialName)
	fs := new(FakeSerial)

	fs.Address = SerialName

	var err error
	fs.ptyMaster, fs.ptySlave, err = termios.Pty()
	if err != nil {
		panic(err)
	}
	fmt.Println("Master: ", fs.ptyMaster.Name())
	fmt.Println("Slave: ", fs.ptySlave.Name())

	err = os.Chmod(fs.ptySlave.Name(), 0660)
	if err != nil {
		panic(err)
	}
	err = os.Symlink(fs.ptySlave.Name(), fs.Address)
	if err != nil {
		panic(err)
	}

	// Set up fake device
	//setNonBlock(fs.ptyMaster)
	fs.ptySettings = new(unix.Termios)
	termios.Tcgetattr(fs.ptyMaster.Fd(), fs.ptySettings)
	termios.Tcsetattr(fs.ptyMaster.Fd(), termios.TCSADRAIN, fs.ptySettings)

	// make streams
	fs.ReceiveStream = make(chan byte, 1000)
	fs.SendStream = make(chan byte, 1000)
	fs.SlaveOpenChan = make(chan bool, 10)

	// Add watcher to slave so we know when it is opened
	fmt.Println("Setting up watcher for slave")
	fs.watcher, err = fsEvents.NewWatcher()
	if err != nil {
		fmt.Println("Error could not make a new watcher ", err)
	}
	_, err = fs.watcher.AddDescriptor(fs.ptySlave.Name(), fsEvents.AllEvents)
	if err != nil {
		fmt.Println("Error could not Add a Descriptor ", err)
	}
	_, err = fs.watcher.AddDescriptor(SerialName, fsEvents.AllEvents)
	if err != nil {
		fmt.Println("Error could not Add a Descriptor ", err)
	}
	err = fs.watcher.RegisterEventHandler(fs)
	if err != nil {
		fmt.Println("Error could not register event handler ", err)
	}
	fmt.Println("Descriptors Added")
	fmt.Println(fs.watcher.ListDescriptors())
	go fs.watcher.WatchAndHandle()
	err = fs.watcher.StartAll()
	if err != nil {
		fmt.Println("Could not start watcher: ", err)
	}
	fs.SlaveOpen = false

	return fs
}

// Handle is a interface to fsEvents
func (fs *FakeSerial) Handle(w *fsEvents.Watcher, event *fsEvents.FsEvent) error {
	switch mask := event.RawEvent.Mask; mask {
	case fsEvents.Accessed:
	case fsEvents.Open:
		fmt.Println("Slave was Opened")
		fs.SlaveOpen = true
		fs.SlaveOpenChan <- fs.SlaveOpen
		return nil
	case fsEvents.CloseWrite:
	case fsEvents.CloseRead:
		fmt.Println("Disconnect from Slave")
		fs.SlaveOpen = false
		fs.SlaveOpenChan <- fs.SlaveOpen
		return nil
	case fsEvents.Modified:
		// someone sent a line to the serial device
		return nil
	default:
		eventids := make(map[string]uint32)
		eventids["Accessed"] = fsEvents.Accessed
		eventids["Modified"] = fsEvents.Modified
		eventids["AttrChange"] = fsEvents.AttrChange
		eventids["CloseWrite"] = fsEvents.CloseWrite
		eventids["CloseRead"] = fsEvents.CloseRead
		eventids["Open"] = fsEvents.Open
		eventids["MovedFrom"] = fsEvents.MovedFrom
		eventids["MovedTo"] = fsEvents.MovedTo
		eventids["Move"] = fsEvents.Move
		eventids["Create"] = fsEvents.Create
		eventids["Delete"] = fsEvents.Delete
		eventids["RootDelete"] = fsEvents.RootDelete
		eventids["RootMove"] = fsEvents.RootMove
		eventids["IsDir"] = fsEvents.IsDir
		fmt.Println("Slave had something else happen to it")
		for ev, id := range eventids {
			if mask == id {
				fmt.Println("Event was ", ev)
			}
		}
	}

	return nil
}

// Check is a interface to fsEvents
func (fs *FakeSerial) Check(event *fsEvents.FsEvent) bool {
	return true
}

// GetMask is a interface to fsEvents
func (fs *FakeSerial) GetMask() uint32 {
	return fsEvents.AllEvents
}

// Close will close the pty connection
func (fs *FakeSerial) Close() {
	os.RemoveAll(SerialName)
	fs.ptyMaster.Close()
	fs.ptySlave.Close()
}

func setNonBlock(fd *os.File) {
	err := syscall.SetNonblock(int(fd.Fd()), true)
	if err != nil {
		panic(err)
	}
}

// ReadMaster will read all available bytes coming in over the serial device
func (fs *FakeSerial) ReadMaster() {
	for {
		buf := make([]byte, 1)
		_, err := io.ReadAtLeast(fs.ptyMaster, buf, 1)
		if err != nil {
			fmt.Println("ReadMaster got Err: ", err)
			return
		}

		fs.ReceiveStream <- buf[0]
	}
}

func (fs *FakeSerial) SendBytes(buf []byte) {
	// if the slave isn't open then don't send anything
	if !fs.SlaveOpen {
		return
	}
	_, err := fs.ptyMaster.Write(buf)
	if err != nil {
		fmt.Println("SEND BYTES ERR: ", err.Error())
		return
	}
}

// SendMaster will send any bytes that come over the stream
func (fs *FakeSerial) SendMaster() {
	// if the slave isn't open then don't send anything
	if !fs.SlaveOpen {
		return
	}
	for buf := range fs.SendStream {

		buffy := make([]byte, 1)
		buffy[0] = buf
		_, err := fs.ptyMaster.Write(buffy)
		if err != nil {
			fmt.Println("SEND MASTER ERR:", err.Error())
			return
		}

	}
}
