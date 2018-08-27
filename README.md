# MangoOS

### What is this?
MangoOS is a collection of tools for 3D printing connected over a common D-Bus interface. Right now this is not much to look at, but as I progress the code base we will see more and more tools and features.

### What tools are included?

- Commango
Daemon for controlling the Printer Serial Port. This tool will share the serial port with other devices over DBUS. Any other program can tap into the Serial Feed and Write lines to the printer.

- MangoCLI
CLI tool for Executing commands to MangoOS. Currently it is just a CLIGUI tool for connecting to Commango

### What tools are coming?

- MangoFM
File Manager that will organize and keep track of gcode files.

- BuffaloMango
Go Buffalo Server to host a webpage users can use on their home networks. This will use the Go Buffalo Package.

- MangoCore
An API Server that will act as the main point of communication for any services that want to connect over cloud infrastructures. It will provide a secure connection over ssl to your print server. This may be tied into BuffaloMango using Go Buffalo as well

- Tango
GUI for interacting with your printer. I haven't decided what framework to use. Go doesn't have very many GUI frameworks yet.

### Why are you writing this in go?
Go strikes a balance between developing rapidly and developing clean code that is easy to read. C++ is too heavy for this project and a python printer server has already been made. Go is the easiest language to cross compile I have ever experienced, and it is extremely easy to write tests for it. Most Go packages are written purely in Go so the entire project could compile down into a few hundered MB of information. 

### Why are you writing this at all? We already have Octoprint.
For the past two years I have developed on Octoprint. Getting anything to work in it has been a nightmare. If you add too much processing suddenly the entire server becomes unresponsive. Oh you want to connect to the webserver in a middle of a print, well let's make the current print stutter because we are pushing a webpage. Oh you made a request? maybe I'll process it now, Maybe I'll process it in a few minutes. who knows? not I. Oh you want to open another thread? Welcome to the python GIL. Your thread is a lie. Oh you want octoprint to share it's resources with other processes on the smae machine on something besides a REST API? nah. We don't do that here in octoprint land. 

I'm mostly writing this because I would like to make a server that allows other people to make programs that have small specific purposes. You don't like our fileserver? keep the DBUS functions intact and you can alter the backend for that program to your specifications. Want to do processing on the incoming data from the printerboard? Make a small program just for that and tap into the DBUS endpoints to get the info you require. Basically I just want to create a server that is a collection of small effecient programs.
