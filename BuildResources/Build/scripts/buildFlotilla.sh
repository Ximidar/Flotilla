#!/bin/sh
# @Author: Ximidar
# @Date:   2018-10-21 22:12:32
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-12 12:10:02

# Paths to different important locations
BINDIR=/home/flotilla/bin
ROOTFLOTILLA=/home/flotilla/Flotilla
NODELAUNCHER=$FLOTILLA_DIR/NodeLauncher

# Make Directories
mkdir -p $BINDIR
mkdir -p $ROOTFLOTILLA

# Build Node Launcher
echo "Making Node Launcher"
cd $NODELAUNCHER
make
echo "Done making Node Launcher"

# Copy NodeLauncher to BINDIR
echo "Copying $NODELAUNCHER/bin/ to $BINDIR/"
cp -r $NODELAUNCHER/bin/* $BINDIR/

# Use NodeLauncher to build Flotilla Folder
echo "Building Flotilla Root Folder at $ROOTFLOTILLA"
cd $BINDIR/
./NodeLauncher CreateRoot -p $ROOTFLOTILLA

# Place NATS Server into Folder with built binaries
echo "Copying Nats Server to Core Flotilla"
cp /usr/local/nats/gnatsd $ROOTFLOTILLA/bin/CoreFlotilla/

echo "Placing all made files under user and group 1000"
chown 1000:1000 -R /home/flotilla

echo "Done Building Flotilla!"