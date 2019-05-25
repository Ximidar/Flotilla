/*
* @Author: Ximidar
* @Date:   2019-02-04 15:21:22
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-22 17:15:38
 */

package Monitor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"syscall"
)

// Monitor will create a Process, then monitor all of the output and store it in a log.
// It will also tie the output to stdout/err
type Monitor struct {
	Name     string
	FullPath string
	Args     []string
	Exec     *exec.Cmd
	Output   chan string
	mux      sync.Mutex

	lf LoggingFunc
}

// LoggingFunc will be used to log output from the monitor
type LoggingFunc func(name string, message string)

// NewMonitor will take a string that defines a path to an executable and
// will return a monitor object
func NewMonitor(ExecPath string, lf LoggingFunc, args ...string) (*Monitor, error) {
	monitor := new(Monitor)
	monitor.FullPath = path.Clean(ExecPath)
	monitor.Args = args
	err := monitor.InitMonitor()
	if err != nil {
		return nil, err
	}
	monitor.lf = lf
	return monitor, nil
}

// InitMonitor will fill in the rest of the monitor data based upon the FullPath variable
func (mon *Monitor) InitMonitor() error {
	// Get the basename for the program we are running
	mon.Name = path.Base(mon.FullPath)

	// #nosec
	// Make the cmd to run
	mon.Exec = exec.Command(mon.FullPath, mon.Args...)

	// Tie the output to a channel
	mon.Output = make(chan string)

	// Check the FullPath variable to see if it exists
	return mon.checkPaths()
}

func (mon *Monitor) checkPaths() error {
	if _, err := os.Stat(mon.FullPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

// StartProcessInGoroutine will start the process in a goroutine
func (mon *Monitor) StartProcessInGoroutine() {
	mon.lf(mon.Name, fmt.Sprintf("Starting Process %v with args %v\n", mon.Name, strings.Join(mon.Args, " ")))
	go mon.startProcess()
	go mon.consumeLines()
}

func (mon *Monitor) consumeLines() {
	for {
		select {
		case l := <-mon.Output:
			mon.lf(mon.Name, l)
		}
	}
}

func (mon *Monitor) captureOutput(outReader, errReader io.ReadCloser) {

	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		if bytes.Contains(data, []byte("\n")) {
			advance = bytes.IndexByte(data, byte('\n')) + 1

			if advance < len(data) {
				token = append(token, data[:advance]...)

			} else {
				token = append(token, data...)
			}

		} else if atEOF {
			if len(data) > 1 {
				token = data
				advance = len(token)
				return
			}
			err = io.EOF
			return

		}

		return
	}

	outScanner := bufio.NewScanner(outReader)
	outScanner.Split(scanLines)

	errScanner := bufio.NewScanner(errReader)
	errScanner.Split(scanLines)

	contScan := func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			read := scanner.Text()

			if !strings.HasSuffix(read, "\n") {
				read += "\n"
			}

			mon.mux.Lock()
			mon.Output <- read
			mon.mux.Unlock()

		}
		mon.lf(mon.Name, "Scanner Exit")
	}

	go contScan(outScanner)
	go contScan(errScanner)

}

func (mon *Monitor) startProcess() error {

	// create a pipe for the output of the Process
	outReader, err := mon.Exec.StdoutPipe()
	if err != nil {
		mon.lf(mon.Name, fmt.Sprintf("Error creating StdoutPipe for Cmd %v", err))
		return err
	}
	errReader, err := mon.Exec.StderrPipe()
	if err != nil {
		mon.lf(mon.Name, fmt.Sprintf("Error creating StderrPipe for Cmd %v", err))
		return err
	}
	// capture the live output
	mon.captureOutput(outReader, errReader)

	// start the process
	err = mon.Exec.Start()
	if err != nil {
		mon.lf(mon.Name, fmt.Sprintf("Could Not Start %v! %v\n", mon.Name, err))
		return err
	}

	// wait for process to exit
	err = mon.Exec.Wait()
	if err != nil {
		mon.lf(mon.Name, fmt.Sprintf("Program Exited: %v! %v\n", mon.Name, err))
		return err
	}

	return nil
}

// KillMe will kill the current running process
func (mon *Monitor) KillMe() error {
	mon.lf(mon.Name, fmt.Sprint("Asking Process ", mon.Name, " to exit\n"))
	err := syscall.Kill(mon.Exec.Process.Pid, syscall.SIGINT)
	if err != nil {
		mon.lf(mon.Name, err.Error())
		return err
	}
	return mon.Exec.Wait()
}

// Exit will ask the process to exit
func (mon *Monitor) Exit() error {
	err := mon.KillMe()
	return err
}
