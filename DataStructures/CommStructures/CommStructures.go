/*
* @Author: Ximidar
* @Date:   2018-09-02 01:36:21
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-26 15:05:14
 */

package CommStructures

const (
	// Name is the programs name
	Name = "commango."

	// reply subs

	// ListPorts will list all available ports to connect to
	ListPorts = Name + "list_ports"
	// ListOptions will list all available ports and suggested speeds
	ListOptions = Name + "list_options"
	// InitializeComm will initialize the comm object
	InitializeComm = Name + "init_comm"
	// ConnectComm will connect the Comm Object to the Serial Connection
	ConnectComm = Name + "connect_comm"
	// DisconnectComm will disconnect the Comm Object
	DisconnectComm = Name + "disconnect_comm"
	// WriteComm will write a message to the Serial Line
	WriteComm = Name + "write_comm"
	// GetStatus will get the current connection status
	GetStatus = Name + "get_status"

	// Print Controls

	// IsPrinting will query if the printer is printing or not
	IsPrinting = Name + "IS_PRINTING"
	// IsPaused will query if the printer is paused
	IsPaused = Name + "IS_PAUSED"
	// TogglePause will toggle the pause state
	TogglePause = Name + "TOGGLE_PAUSE"
	// StartPrint will start a print
	StartPrint = Name + "START_PRINT"
	// CancelPrint will cancel a print
	CancelPrint = Name + "CANCEL_PRINT"

	// pubs

	// ReadLine will receive all Read Lines from the Comm Object
	ReadLine = Name + "read_line"
	// WriteLine will receive all Written Lines from the Comm Object
	WriteLine = Name + "write_line"
	// StatusUpdate will receive all Connection Updates
	StatusUpdate = Name + "status_update"
)

// These structures are now protobuffers
// InitComm Use this to define a connection to connect to
// type InitComm struct {
// 	Port string `json:"port"`
// 	Baud int    `json:"baud"`
// }

// // CommStatus Use this to get the Comm Status
// type CommStatus struct {
// 	Port      string `json:"port"`
// 	Baud      string `json:"baud"`
// 	Connected bool   `json:"connected"`
// }
