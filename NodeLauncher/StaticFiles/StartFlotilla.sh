# This file will start the Flotilla Instance in the present working directory

PWD=$(pwd)
NodeLauncher=$PWD/bin/Extras/NodeLauncher

# Check if NodeLauncher exists
if [[ ! -f $NodeLauncher ]]; then
	echo "Node Launcher was not found at position $NodeLauncher"
	exit 1
fi

# Launch Node Launcher
$NodeLauncher Start -p $PWD --tls