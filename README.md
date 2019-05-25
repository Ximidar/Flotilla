# Flotilla
[![Go Report Card](https://goreportcard.com/badge/github.com/Ximidar/Flotilla)](https://goreportcard.com/report/github.com/Ximidar/Flotilla)


### What is this?
Flotilla is a 3D printing server. Use this to print files and have a nice interface to your printer. Currently Flotilla is under construction in my free time. It is not stable yet.

### What tools are included?

- Commango
This controls the serial connection to the printer. It will send and receive Marlin data and share that data with the rest of the cluster

- Flotilla_CLI
CLI tool for Executing commands to Flotilla. Currently it is just a CLIGUI tool for connecting to Commango

- FM
File Manager that keeps track of files and folders, as well as serves the file up for printing.

- NodeLauncher
This will create the Flotilla Package for distribution. It also will provide a way to start and stop Flotilla and it will distribute the config to the rest of the program.

- FlotillaWeb
This is the web interface for the server (Currently under construction)

- FlotillaStatus
This will take in information from the printer and translate it into easy information for the rest of the cluster to ingest. 


