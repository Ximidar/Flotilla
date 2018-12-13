#!/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 22:12:32
# @Last Modified by:   Ximidar
# @Last Modified time: 2018-12-12 23:10:27

COMMANGO=$FLOTILLA_DIR/Commango/
FLOTILLA_CLI=$FLOTILLA_DIR/Flotilla_CLI/
FLOTILLA_FILE_MANAGER=$FLOTILLA_DIR/Flotilla_File_Manager/
BINDIR=$FLOTILLA_DIR/bin/
mkdir $BINDIR


#Build all assets
cd $COMMANGO
make
cd $FLOTILLA_CLI
make
cd $FLOTILLA_FILE_MANAGER
make

#Copy all built binaries to the bindir
cp -r $COMMANGO/bin/ $BINDIR/
cp -r $FLOTILLA_CLI/bin/ $BINDIR/
cp -r $FLOTILLA_FILE_MANAGER/bin/ $BINDIR/
