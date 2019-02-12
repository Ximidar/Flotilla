# Building 

### Overview
Building happens with docker. Docker will take in all of the repos and dependencies, then test, then build the binaries. It will build two sets of binaries for each program. An Arm binary and a binary for the system you are on (x86_64). 

### Building Manually
Each repository should have a makefile in it already. Simply navigate to the repo and type in the command
```
make
```
This will drop the two binaries in a folder called `bin`. This way will fail if you do not have all the dependencies installed. 

### Building with Docker
To test building:
```
docker build -t ximidar/flotilla:latest .
```
This command will run all tests then build the Flotilla Package.

To get the binaries out:
```
mkdir buildFlotilla

# Run this command inside of Flotilla directory
docker build -t ximidar/flotilla:latest .

# Run this command anywhere
docker run -v ./buildFlotilla:/home/flotilla/ ximidar/flotilla:latest
```
This will drop a Flotilla Package Folder into `./buildFlotilla`

### Note
Currently this will only make the x86 Flotilla Package. You will have to construct an arm package yourself for right now. This will change soon. 

### Running Flotilla
To start Flotilla use the `NodeLauncher` Program in the `Flotilla/bin/Extras` folder. You will need to remember the path to where you placed your Flotilla Folder. For simplicity I placed mine in my home folder. To start Flotilla run this command:

```
~/Flotilla/bin/Extras/NodeLauncher Start -p ~/Flotilla/
```

This will start a Flotilla Instance using the binaries in the `~/Flotilla` folder