# !/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 18:57:47
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-12 11:53:15

# NATS
echo "Getting nats-io"
go get github.com/nats-io/go-nats

# Flotilla_CLI
echo "Getting Cobra"
go get github.com/spf13/cobra
echo "Getting Gocui"
go get github.com/ximidar/gocui
echo "Getting go-humanize"
go get github.com/dustin/go-humanize

# Flotilla_File_Manager
echo "Getting fsnotify and golang sys"
go get -u golang.org/x/sys/...
go get github.com/fsnotify/fsnotify

# Commango
echo "Getting bug.st Serial"
mkdir -p $GOPATH/src/go.bug.st/
git clone https://github.com/bugst/go-serial.git -b v1 $GOPATH/src/go.bug.st/serial.v1
echo "Getting goselect"
go get github.com/creack/goselect

# Fake Serial Device
echo "Getting termios"
go get github.com/pkg/term/termios

