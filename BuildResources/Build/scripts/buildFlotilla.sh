#!/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 22:12:32
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-17 15:27:48

# Paths to different important locations
BINDIR=/home/flotilla/bin
ARCH_FOLDER=/home/flotilla/build
AMD64=$ARCH_FOLDER/AMD64
ARM=$ARCH_FOLDER/ARM
ARM64=$ARCH_FOLDER/ARM64
NODELAUNCHER=$FLOTILLA_DIR/NodeLauncher

# Function for detecting failure
GLOBAL_RET=0
check_retval(){
    echo $1
    if [ "$1" != "0" ] ; then
    	GLOBAL_RET=$(($GLOBAL_RET+1))
    	exit $GLOBAL_RET
    fi
}

# Make Directories
mkdir -p $BINDIR
mkdir -p $AMD64
mkdir -p $ARM64
mkdir -p $ARM

# Build Node Launcher
echo "Making Node Launcher"
cd $NODELAUNCHER
make
echo "Done making Node Launcher"

# Copy NodeLauncher to BINDIR
echo "Copying $NODELAUNCHER/bin/ to $BINDIR/"
cp -r $NODELAUNCHER/bin/* $BINDIR/

# Use NodeLauncher to build Flotilla Folder
echo "Building Flotilla Root Folders for all Arches"
cd $BINDIR/

echo "Building Flotilla For AMD64"
./NodeLauncher CreateRoot -p $AMD64/Flotilla -a amd64 -l false
check_retval $?

echo "Building Flotilla For ARM64"
./NodeLauncher CreateRoot -p $ARM64/Flotilla -a arm64 -l false
check_retval $?

echo "Building Flotilla For ARM"
./NodeLauncher CreateRoot -p $ARM/Flotilla -a arm -l false
check_retval $?


# Place NATS Server into Folder with built binaries
echo "Copying Nats Server to Flotilla Packages"
cp /usr/local/natsAMD64/gnatsd $AMD64/Flotilla/bin/CoreFlotilla/
cp /usr/local/natsARM64/gnatsd $ARM64/Flotilla/bin/CoreFlotilla/
cp /usr/local/natsARM6/gnatsd $ARM/Flotilla/bin/CoreFlotilla/

echo "Placing all made files under user and group 1000"
chown 1000:1000 -R /home/flotilla

echo "Done Building Flotilla!"