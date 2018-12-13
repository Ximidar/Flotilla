#!/bin/sh
# @Author: Ximidar
# @Date:   2018-12-12 22:52:23
# @Last Modified by:   Ximidar
# @Last Modified time: 2018-12-12 23:33:56

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
cd $FLOTILLA_DIR/BuildResources/Test
check_retval $?
if [ "$GLOBAL_RET" != "0" ] ; then
	echo "Could Not Reach the test locations"
	exit $GLOBAL_RET
fi

# Store testing location
TEST_FOLDER=$(pwd)
echo $TEST_FOLDER

# Test Data Structures
echo "Testing DataStructures"
cd $TEST_FOLDER/DataStructures/file_structure_tests
go test
check_retval $?

# TODO Update this test to start a NATS server for testing otherwise this will fail
# Test FlotillaCLI
echo "Testing FlotillaCLI"
cd $TEST_FOLDER/FlotillaCLI/FlotillaInterface
go test
#check_retval $?

# Test Flotilla File Manager
echo "Testing Flotilla File Manager"
cd $TEST_FOLDER/FlotillaFileManager/FileManager
go test
check_retval $?

cd $TEST_FOLDER/FlotillaFileManager/FileStreamer
go test
check_retval $?

# Exit with the failed score
exit $GLOBAL_RET