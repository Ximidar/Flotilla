# Building 

### Overview
Building happens with docker. Docker will take in all of the repos and dependencies, then test, then build the binaries. It will build two sets of binaries for each program. An Arm binary and a binary for the system you are on (x86_64). 

### Building Manually
Each repository should have a makefile in it already. Simply navigate to the repo and type in the command
```
make build
```
This will drop the binaries in a folder called `bin`. 

### Testing with Docker
Build:
```
docker build --rm -t ximidar/flotilla_root:latest .
```
Then running it:
```
docker-compose up --build
```
This will also make the code be in hot reload mode. This mode is really slow for the webserver, but go files get reloaded very fast.