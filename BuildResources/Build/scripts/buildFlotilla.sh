#!/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 22:12:32
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-28 16:02:50

# Paths to different important locations
BINDIR=$HOME/bin
ARCH_FOLDER=$HOME/build
AMD64=$ARCH_FOLDER/AMD64
ARM=$ARCH_FOLDER/ARM
ARM64=$ARCH_FOLDER/ARM64
NODELAUNCHER=$FLOTILLA_DIR/NodeLauncher

# Function for detecting failure
GLOBAL_RET=0
check_retval(){
    if [ "$1" != "0" ] ; then
    	echo "$2"
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
if make; then
	echo "Done making Node Launcher"
else
	check_retval "1" "Building NodeLauncher failed"
fi

# Copy NodeLauncher to BINDIR
echo "Copying $NODELAUNCHER/bin/ to $BINDIR/"
cp -r $NODELAUNCHER/bin/* $BINDIR/

# Use NodeLauncher to build Flotilla Folder
echo "Building Flotilla Root Folders for all Arches"
cd $BINDIR/

echo "Building Flotilla For AMD64"
./NodeLauncher CreateRoot -p $AMD64/Flotilla -a=amd64 -l=true
check_retval $? "Could not create Flotilla for AMD64"

echo "Building Flotilla For ARM64"
./NodeLauncher CreateRoot -p $ARM64/Flotilla -a=arm64 -l=true -m=true
check_retval $? "Could not create Flotilla for AMD64"

echo "Building Flotilla For ARM"
./NodeLauncher CreateRoot -p $ARM/Flotilla -a=arm -l=true -m=true
check_retval $? "Could not create Flotilla for AMD64"


# Place NATS Server into Folder with built binaries
echo "Copying Nats Server to Flotilla Packages"
cp $NATS_LOC/AMD64/gnatsd $AMD64/Flotilla/bin/CoreFlotilla/
cp $NATS_LOC/ARM64/gnatsd $ARM64/Flotilla/bin/CoreFlotilla/
cp $NATS_LOC/ARM6/gnatsd $ARM/Flotilla/bin/CoreFlotilla/

echo "Making user 1000 own all files"
chown 1000:1000 -R $ARCH_FOLDER

echo "Done Building Flotilla!"