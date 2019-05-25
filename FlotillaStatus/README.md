# Flotilla Status

[![Go Report Card](https://goreportcard.com/badge/github.com/Ximidar/FlotillaStatus)](https://goreportcard.com/report/github.com/Ximidar/FlotillaStatus)

This repository will query a status from the comm layer. It will also serve as the intermediary between any program wanting to send something to Commango. This is necessary because of how sending lines to the printer works while printing.

While printing users will send lines like this:
```
M105
```
This seems innocuous but it will cause an error while printing. This repo will keep track of all incomming lines and transform them to this:
```
N### M105 *####
```
Or however they need to be transformed

This will assure that all incomming data gets transformed to the printer correctly.

This repo will also keep track of things like:
- Temperature
- Incomming Errors
- all "ok" and "wait" and "processing" signals Marlin sends
- Pause/Cancel/Resume Commands
- Print Start Command
- Any other signals from marlin that need to be monitored
