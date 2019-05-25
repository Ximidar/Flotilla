/*
* @Author: Ximidar
* @Date:   2019-03-01 19:37:42
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 16:36:01
 */

package CommRelay

import (
	"fmt"
	"testing"
	"time"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
)

func TestProcessLine(t *testing.T) {
	resend := make(chan int, 10)
	ok := make(chan bool, 10)
	wait := make(chan bool, 10)
	lp, err := NewLineParser(resend, ok, wait)
	CommonTestTools.CheckErr(t, "Could not make line parser", err)

	lp.ProcessLine("ok\n")
	select {
	case <-ok:
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not process OK signal")
	}

	lp.ProcessLine("Resend: 10")
	select {
	case number := <-resend:
		if number != 10 {
			t.Fatal("Could not return correct number")
		}
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not process Resend signal")
	}

	lp.ProcessLine("wait\n")
	select {
	case <-wait:
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not process WAIT signal")
	}

	waitSig := []byte{119, 97, 105, 116, 10}
	lp.ProcessLine(string(waitSig))
	select {
	case <-wait:
		fmt.Println("Alright")
	case <-time.After(2 * time.Second):
		t.Fatal("Could not process WAIT signal")
	}

}
