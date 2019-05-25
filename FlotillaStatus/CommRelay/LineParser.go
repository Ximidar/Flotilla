/*
* @Author: Ximidar
* @Date:   2018-12-25 19:57:15
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-02 12:16:48
 */

package CommRelay

import (
	"fmt"
	"regexp"
	"strconv"
)

// LineParser will take in lines and identify their messages
type LineParser struct {
	ResendRegex *regexp.Regexp
	OkRegex     *regexp.Regexp
	WaitRegex   *regexp.Regexp
	FindNums    *regexp.Regexp
	ResendChan  chan int
	OKChan      chan bool
	WaitChan    chan bool
}

// NewLineParser will construct a LineParser
func NewLineParser(resendChannel chan int, okChan chan bool, waitChan chan bool) (*LineParser, error) {
	lp := new(LineParser)
	lp.ResendChan = resendChannel
	lp.OKChan = okChan
	lp.WaitChan = waitChan
	err := lp.compileRegex()
	return lp, err
}

/* compile Regex will build regexp for the following conditions
RECV: Error:Line Number is not Last Line Number+1, Last Line: 8
RECV: Resend: 9
RECV: ok
*/

func (lp *LineParser) compileRegex() error {
	var err error
	lp.ResendRegex, err = regexp.Compile("Resend:([/ ]+)([0-9]+)")
	if err != nil {
		return err
	}
	lp.OkRegex, err = regexp.Compile("(ok|OK|Ok)")
	if err != nil {
		return err
	}
	lp.WaitRegex, err = regexp.Compile("(wait|WAIT|Wait)")
	if err != nil {
		return err
	}
	lp.FindNums, err = regexp.Compile("([0-9.-]*[0-9]+)")
	if err != nil {
		return err
	}
	return nil
}

// ProcessLine will process if a Resend Error has occured or if ok is returned
func (lp *LineParser) ProcessLine(line string) {
	if lp.ResendRegex.MatchString(line) {
		// Process Resend
		lp.processResend(line)
		return
	} else if lp.OkRegex.MatchString(line) {
		// Process OK
		lp.OKChan <- true
		return
	} else if lp.WaitRegex.MatchString(line) {
		// Process Wait
		lp.WaitChan <- true
		return
	}
	return
}

// processResend should be given a line to resend. It should resend that line as well as reset where
// The next sent line should be.
func (lp *LineParser) processResend(line string) {

	// Find line number
	nums := lp.FindNums.Find([]byte(line))
	intnum, err := strconv.ParseInt(string(nums), 10, 64)
	if err != nil {
		println(err.Error())
		return
	}
	intnum2 := int(intnum)
	fmt.Printf("line: %v is asking us to resend lineNum %v\n", line, intnum2)
	lp.ResendChan <- intnum2
}
