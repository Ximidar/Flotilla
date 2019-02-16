/*
* @Author: Ximidar
* @Date:   2019-02-15 12:16:19
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-15 12:35:10
 */

// Package ffmt will use the debug flag from common tools to print out statements
// it works exactly like fmt except it will check if we are supposed to be printing
// out statements
package ffmt

import (
	"fmt"
	"io"

	"github.com/ximidar/Flotilla/CommonTools/flags"
)

// Errorf works like it does in fmt except it checks a debug variable
func Errorf(format string, a ...interface{}) error {
	if commonFlags.Debug {
		return fmt.Errorf(format, a...)
	}
	return nil
}

// Fprint works like it does in fmt except it checks a debug variable
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fprint(w, a...)
	}
	return 0, nil
}

// Fprintf works like it does in fmt except it checks a debug variable
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fprintf(w, format, a...)
	}
	return 0, nil
}

// Fprintln works like it does in fmt except it checks a debug variable
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fprintln(w, a...)
	}
	return 0, nil
}

// Fscan works like it does in fmt except it checks a debug variable
func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fscan(r, a...)
	}
	return 0, nil
}

// Fscanf works like it does in fmt except it checks a debug variable
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fscanf(r, format, a...)
	}
	return 0, nil
}

// Fscanln works like it does in fmt except it checks a debug variable
func Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Fscanln(r, a...)
	}
	return 0, nil
}

// Print works like it does in fmt except it checks a debug variable
func Print(a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Print(a...)
	}
	return 0, nil
}

// Printf works like it does in fmt except it checks a debug variable
func Printf(format string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}

// Println works like it does in fmt except it checks a debug variable
func Println(a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Println(a...)
	}
	return 0, nil
}

// Scan works like it does in fmt except it checks a debug variable
func Scan(a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Scan(a...)
	}
	return 0, nil
}

// Scanf works like it does in fmt except it checks a debug variable
func Scanf(format string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Scanf(format, a...)
	}
	return 0, nil
}

// Scanln works like it does in fmt except it checks a debug variable
func Scanln(a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Scanln(a...)
	}
	return 0, nil
}

// Sprint works like it does in fmt except it checks a debug variable
func Sprint(a ...interface{}) string {
	if commonFlags.Debug {
		return fmt.Sprint(a...)
	}
	return ""
}

// Sprintf works like it does in fmt except it checks a debug variable
func Sprintf(format string, a ...interface{}) string {
	if commonFlags.Debug {
		return fmt.Sprintf(format, a...)
	}
	return ""
}

// Sprintln works like it does in fmt except it checks a debug variable
func Sprintln(a ...interface{}) string {
	if commonFlags.Debug {
		return fmt.Sprintln(a...)
	}
	return ""
}

// Sscan works like it does in fmt except it checks a debug variable
func Sscan(str string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Sscan(str, a...)
	}
	return 0, nil
}

// Sscanf works like it does in fmt except it checks a debug variable
func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Sscanf(str, format, a...)
	}
	return 0, nil
}

// Sscanln works like it does in fmt except it checks a debug variable
func Sscanln(str string, a ...interface{}) (n int, err error) {
	if commonFlags.Debug {
		return fmt.Sscanln(str, a...)
	}
	return 0, nil
}
