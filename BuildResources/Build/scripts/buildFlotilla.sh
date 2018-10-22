#!/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 22:12:32
# @Last Modified by:   Ximidar
# @Last Modified time: 2018-10-21 22:18:46

COMMANGO=$FLOTILLA_DIR/Commango/
FLOTILLA_CLI=$FLOTILLA_DIR/Flotilla_CLI/
FLOTILLA_FILE_MANAGER=$FLOTILLA_DIR/Flotilla_File_Manager/
BINDIR=$FLOTILLA_DIR/bin/
mkdir $BINDIR

cd $COMMANGO
go build -o $BINDIR/Commango

cd $FLOTILLA_CLI
go build -o $BINDIR/Flotilla_CLI

cd $FLOTILLA_FILE_MANAGER
go build -o $BINDIR/Flotilla_File_Manager

