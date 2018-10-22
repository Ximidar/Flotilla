# !/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 18:57:47
# @Last Modified by:   Ximidar
# @Last Modified time: 2018-10-21 22:07:56

# NATS
go get github.com/nats-io/go-nats

# Flotilla_CLI
go get github.com/spf13/cobra
go get github.com/ximidar/gocui

# Flotilla_File_Manager
go get -u golang.org/x/sys/...
go get github.com/fsnotify/fsnotify

# Commango
mkdir -p $GOPATH/src/go.bug.st/
git clone https://github.com/bugst/go-serial.git -b v1 $GOPATH/src/go.bug.st/serial.v1
go get github.com/creack/goselect

