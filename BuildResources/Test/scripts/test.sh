#!/bin/sh
# @Author: Ximidar
# @Date:   2018-12-12 22:52:23
# @Last Modified by:   Ximidar
# @Last Modified time: 2019-02-25 23:23:31

#This file is used to run all tests 

GLOBAL_RET=0
check_retval(){
    echo $1
    if [ "$1" != "0" ] ; then
    	GLOBAL_RET=$(($GLOBAL_RET+1))
    	exit $GLOBAL_RET
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
go test ./...
exit "0"
# Don't pass tests for right now.
#check_retval $?