#!/bin/sh
# @Author: Ximidar
# @Date:   2018-12-12 22:52:23
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-26 13:48:05

#This file is used to run all tests 

checkError(){
    if [ "$1" != "0" ] ; then
    	echo "$2"
    	exit $1
    fi
}

test_location(){
	echo "Testing $1"
	if cd $2 ; then
		go test ./...
		checkError $? "Test $1 Failed"
	else
		echo "Could not reach $1 at destination $2"
		exit 1
	fi
}

# Go to the testing location

if cd $FLOTILLA_DIR/BuildResources/Test ; then
	echo "Start Testing"
	# Store testing location
	TEST_FOLDER=$(pwd)
	echo $TEST_FOLDER
else
	echo "Could Not Reach the test locations"
	exit "1"
fi

# Run actual tests
# Go to each test directly and run tests in order. go test ./... fails the tests because nats is
# a finicky hellion when shutting down between tests. Instead travel to each folder and test individualy

test_location "Data Structures" $TEST_FOLDER/DataStructures
test_location "Flotilla Status" $TEST_FOLDER/FlotillaStatus
test_location "Node Launcher" $TEST_FOLDER/NodeLauncher
test_location "Flotilla File Manager" $TEST_FOLDER/FlotillaFileManager

echo "Skipping Flotilla System Test because it does not pass in a docker build environment due to the nature of how it spins off multiple processes. It passes fine on a regular computer"
# test_location "Flotilla System Tests" $TEST_FOLDER/FlotillaSystemTests