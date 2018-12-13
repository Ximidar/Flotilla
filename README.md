# Flotilla
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FXimidar%2FFlotilla.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FXimidar%2FFlotilla?ref=badge_shield)


### What is this?
Flotilla is a collection of tools for 3D printing connected over a NATS interface. Right now this is not much to look at, but as I progress the code base we will see more and more tools and features.

### What tools are included?

- Commango
Daemon for controlling the Printer Serial Port. This tool will share the serial port with other devices over NATS. Any other program can tap into the Serial Feed and Write lines to the printer.

- Flotilla_CLI
CLI tool for Executing commands to Flotilla. Currently it is just a CLIGUI tool for connecting to Commango

- FM
File Manager that will organize and keep track of files.

### What tools are coming?

- Web Interface
Go Server to host a webpage users can use on their home networks.

- Tango
GUI for interacting with your printer. Will build with Qt(Maybe the go version, maybe just straight c++. NATS is supposed to be interfacable with any language right?)

- Node Launcher
Program for keeping track of nodes and launching them. This will also keep track of any custom nodes made by third parties and launch those too. This will keep track of output from each node and keep track of the lifecycle of each node. If a node goes down this will restart it or blare warning lights or something.

- Config Node
This will control the configuration options for Flotilla. Other nodes can query this node for what mode they should be in, extra options, everything. This will also let the programs know when it is time to update their configuration.

- CommStatus
This will monitor the output from the Commango and keep track of things like temperature, error status, current position, ect.

### Why are you writing this in go?
Go strikes a balance between developing rapidly and developing clean code that is easy to read. C++ is too heavy for this project and a python printer server has already been made. Go is the easiest language to cross compile I have ever experienced, and it is extremely easy to write tests for it. 

### What Environment do you use to write this code?
I am using Sublime Text 3 and I set up my environment following this [guide](https://www.alexedwards.net/blog/streamline-your-sublime-text-and-go-workflow)



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FXimidar%2FFlotilla.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FXimidar%2FFlotilla?ref=badge_large)