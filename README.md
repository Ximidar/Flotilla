# Flotilla

### What is this?
Flotilla is a collection of tools for 3D printing connected over a NATS interface. Right now this is not much to look at, but as I progress the code base we will see more and more tools and features.

### What tools are included?

- Commango
Daemon for controlling the Printer Serial Port. This tool will share the serial port with other devices over DBUS. Any other program can tap into the Serial Feed and Write lines to the printer.

- Flotilla_CLI
CLI tool for Executing commands to Flotilla. Currently it is just a CLIGUI tool for connecting to Commango

### What tools are coming?

- FM
File Manager that will organize and keep track of gcode files.

- Web Interface
Go Buffalo Server to host a webpage users can use on their home networks. This will use the Go Buffalo Package.

- API
An API Server that will act as the main point of communication for any services that want to connect over cloud infrastructures. It will provide a secure connection over ssl to your print server. This may be tied into BuffaloMango using Go Buffalo as well

- Tango
GUI for interacting with your printer. I haven't decided what framework to use. Go doesn't have very many GUI frameworks yet.

### Why are you writing this in go?
Go strikes a balance between developing rapidly and developing clean code that is easy to read. C++ is too heavy for this project and a python printer server has already been made. Go is the easiest language to cross compile I have ever experienced, and it is extremely easy to write tests for it. 

