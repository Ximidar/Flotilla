# !/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 18:57:47
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-26 17:12:00

# NATS
echo "Getting nats-io"
go get github.com/nats-io/go-nats

# Flotilla_CLI
echo "Getting Cobra"
go get github.com/spf13/cobra
echo "Getting Gocui"
go get github.com/jroimartin/gocui
echo "Getting TCell"
go get github.com/gdamore/tcell
echo "Getting go-humanize"
go get github.com/dustin/go-humanize

# Flotilla_File_Manager
echo "Getting fsnotify and golang sys"
go get -u golang.org/x/sys/...
go get github.com/fsnotify/fsnotify

# Commango
echo "Getting bug.st Serial"
# mkdir -p $GOPATH/src/go.bug.st/
# git clone https://github.com/bugst/go-serial.git -b v1 $GOPATH/src/go.bug.st/serial.v1
go get go.bug.st/serial
echo "Getting goselect"
go get github.com/creack/goselect

# Fake Serial Device
echo "Getting termios"
go get github.com/pkg/term/termios
echo "Getting fsevents"
go get github.com/tywkeene/go-fsevents

# NodeLauncher
echo "Getting lumberjack"
go get gopkg.in/natefinch/lumberjack.v2
echo "Getting viper"
go get github.com/spf13/viper

# Test
echo "Getting go-nats"
go get github.com/nats-io/go-nats
echo "Getting nats server"
go get github.com/nats-io/gnatsd

# Flotilla Web
echo "Getting Gorilla Mux"
go get github.com/gorilla/mux
echo "Getting Gorilla websocket"
go get github.com/gorilla/websocket
echo "Getting Gorilla handlers"
go get github.com/gorilla/handlers


# Setup Protobuffer
# wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-cpp-3.6.1.zip
echo "Getting Protobuf"
go get -u github.com/golang/protobuf/protoc-gen-go