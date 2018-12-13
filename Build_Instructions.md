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
This command will run all tests then build all binaries.

To get the binaries out:
```
mkdir bin

# Run this command inside of Flotilla directory
docker build -t ximidar/flotilla:latest .

# Run this command anywhere
docker run -v ~/path_to_your_bin/bin/:/home/flotilla/bin ximidar/flotilla:latest
```
This will drop all binaries into `~/path_to_your_bin/bin/`
You will still need to provide your own NATS server for all of these to run together. Go Here(https://nats.io/download/nats-io/gnatsd/) and grab the right server.

### Running Flotilla
Currently you will have to start the NATS server, Flotilla_File_Manager, and Commango in seperate command windows. Then to interface with Flotilla run `./FlotillaCLI ui` Which will start up a ui. If everything is working correctly, Which is not guarunteed, you will be able to control the serial and browse the files. 