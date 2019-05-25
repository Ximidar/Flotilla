/*
* @Author: Ximidar
* @Date:   2018-12-19 12:12:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-05-07 22:43:15
 */

package NatsStatus

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/DataStructures"
	CS "github.com/ximidar/Flotilla/DataStructures/CommStructures"
	"github.com/ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/CommRelayStructures"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	PS "github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	TS "github.com/ximidar/Flotilla/DataStructures/StatusStructures/TemperatureStructures"
	"github.com/ximidar/Flotilla/FlotillaStatus/CommRelay"
	"github.com/ximidar/Flotilla/FlotillaStatus/PlayStatus"
	"github.com/ximidar/Flotilla/FlotillaStatus/StatusMonitor"
)

// TermChannel will monitor for an exit signal
var TermChannel chan os.Signal

// NatsStatus will be the interface for Flotilla Status to the Nats server
type NatsStatus struct {
	NC             *nats.Conn
	StatusMonitor  *StatusMonitor.StatusMonitor
	CommRelay      *CommRelay.CommRelay
	PlayStatusNats *PlayStatus.PlayStatusNats
	RNode          *PlayStructures.RegisteredNode
}

// NewNatsStatus will construct a NatsStatus object
func NewNatsStatus() (*NatsStatus, error) {
	ns := new(NatsStatus)

	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	var err error
	// Construct the different objects
	ns.StatusMonitor, err = StatusMonitor.NewStatusMonitor(ns.PublishTemperature)
	checkErr(err)

	ns.CommRelay, err = CommRelay.NewCommRelay(ns.SendComm, ns.AskForLine, ns.FinishedStream)
	checkErr(err)
	ns.CommRelay.StartEventHandler()

	// Make the connection to NATS
	ns.NC, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		// TODO reconnect
		panic(err)
	}

	// Create the PlayStatusNats
	ns.PlayStatusNats, err = PlayStatus.NewPlayStatusNats(ns.NC)
	checkErr(err)

	// Register Node
	ns.RNode, err = PlayStructures.NewRegisteredNode(PlayStructures.Name+"HeartBeat", ns.NC)
	checkErr(err)

	// Register functions to use with RNode
	err = ns.RegisterFunctionstoRnode()
	checkErr(err)

	//TODO check these errors for now we will assume NATS works 100%
	err = ns.registerSubscriptions()
	checkErr(err)

	return ns, nil
}

func (ns *NatsStatus) registerSubscriptions() error {
	var err error
	_, err = ns.NC.Subscribe(CS.ReadLine, ns.GetComm)
	if err != nil {
		return err
	}
	_, err = ns.NC.Subscribe(CommRelayStructures.WriteLine, ns.WriteFormattedComm)
	if err != nil {
		return err
	}
	_, err = ns.NC.Subscribe(TS.GetTemp, ns.GetTemperature)
	if err != nil {
		return err
	}
	_, err = ns.NC.Subscribe(TS.SetTemp, ns.SetTemperature)
	if err != nil {
		return err
	}
	_, err = ns.NC.Subscribe(TS.GetTempHistory, ns.GetTemperatureHistory)
	if err != nil {
		return err
	}
	_, err = ns.NC.Subscribe(PlayStructures.PublishStatus, ns.RNode.StatusObserver.UpdateStatusFromNats)
	return err
}

// Serve will keep the program open
func (ns *NatsStatus) Serve() {
	// Function for waiting for exit on the main loop
	// Wait for termination
	TermChannel = make(chan os.Signal)
	signal.Notify(TermChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Flotilla Status Started")
	<-TermChannel
	fmt.Println("Recieved Interrupt Sig, Now Exiting.")
	os.Exit(0)
}

// SendComm will send the comm to the commango instance
func (ns *NatsStatus) SendComm(command string) error {
	// Get expected write bytes
	expectedBytes := len(command)

	// Construct CommMessage to send
	mess := new(CS.CommMessage)
	mess.Message = command
	mMess, err := proto.Marshal(mess)
	if err != nil {
		fmt.Println("Could not marshal a Comm Message for", command)
		return err
	}

	// Get the reply
	reply, err := ns.NC.Request(CS.WriteComm, mMess, nats.DefaultTimeout)
	if err != nil {
		fmt.Println("Could not Write Comm")
		return err
	}

	// construct the receipt
	commReceipt := new(CS.WrittenBytes)
	err = proto.Unmarshal(reply.Data, commReceipt)
	if err != nil {
		fmt.Println("Could not unmarshal the returned data", err)
		return err
	}

	if expectedBytes+1 != int(commReceipt.GetBytes()) {
		fmt.Println(fmt.Sprintf("Expected %v != Written %v", expectedBytes, int(commReceipt.GetBytes())))
		return fmt.Errorf("Expected %v != Written %v", expectedBytes, int(commReceipt.GetBytes()))
	}

	return nil
}

// AskForLine will ask for more lines
func (ns *NatsStatus) AskForLine(numOfLines int) {
	rl := new(CommRelayStructures.RequestLines)
	rl.Amount = int32(numOfLines)
	brl, err := proto.Marshal(rl)
	if err != nil {
		fmt.Println("Could not Marshal Request Lines Protobuf", err)
	}
	reply, err := ns.NC.Request(FileStructures.RequestLines, brl, 5*time.Second)

	if err != nil {
		fmt.Println("Error! Could not request more lines", err)
	}

	lines := new(CommRelayStructures.ReturnLines)
	err = proto.Unmarshal(reply.Data, lines)
	if err != nil {
		fmt.Println("Error could not unmarshal returnlines", err)
	}

	if lines.GetEOF() {
		ns.CommRelay.NotifyWhenEmpty()
	}

	for _, line := range lines.GetLines() {
		ns.CommRelay.FormatLine(line)
	}
}

// FinishedStream will request a Done Status be applied to the server
func (ns *NatsStatus) FinishedStream() {
	pstatus, err := PS.NewPlayAction(PS.DONE)
	if err != nil {
		fmt.Println("Cannot send DONE status because", err.Error())
		return
	}
	err = pstatus.Send(ns.NC)
	if err != nil {
		fmt.Println("Could not send DONE status because", err.Error())
	}

	// Set our status to done
	ns.RNode.UpdateNode(PlayStructures.DONE)

}

// SetTemperature will accept a message to set the current temperature
func (ns *NatsStatus) SetTemperature(msg *nats.Msg) {
	temp, err := TS.NewSetTemperatureFromMSG(msg)
	if err != nil {

		reply := DataStructures.ConstructNegativeReplyString(err.Error())
		pack, _ := json.Marshal(reply)
		ns.NC.Publish(msg.Reply, pack)
	}

	command, err := ns.StatusMonitor.TempMonitor.GetTempCommand(temp.Tool, temp.Temp)
	if err != nil {
		reply := DataStructures.ConstructNegativeReplyString(err.Error())
		pack, _ := json.Marshal(reply)
		ns.NC.Publish(msg.Reply, pack)
	}

	line, _ := CommRelayStructures.NewLine(command, 0, false)
	ns.CommRelay.FormatLine(line)

	// Reply Positive
	reply := DataStructures.ReplyString{Success: true, Message: "command sent"}

	pack, _ := json.Marshal(reply)
	ns.NC.Publish(msg.Reply, pack)

}

// GetTemperature will send a JSON object of the current temperature and history if the
// requestor asks for it
func (ns *NatsStatus) GetTemperature(msg *nats.Msg) {
	bTemp, err := ns.StatusMonitor.TempMonitor.CurrentTemperature.PackageJSON()
	if err != nil {
		ns.NC.Publish(msg.Reply, []byte("Could not Package Temperature"))
		return
	}
	ns.NC.Publish(msg.Reply, bTemp)

}

// GetTemperatureHistory will be used to send the Temp History to anyone who asks.
func (ns *NatsStatus) GetTemperatureHistory(msg *nats.Msg) {
	tempHistory := ns.StatusMonitor.TempMonitor.GetTempHistory()
	byteTempHistory, err := json.Marshal(tempHistory)
	if err != nil {
		fmt.Println("Could not return Temp History")
		ns.NC.Publish(msg.Reply, []byte("Could not package History"))
	}
	ns.NC.Publish(msg.Reply, byteTempHistory)
}

// PublishTemperature will publish the temperature to the nats server as it comes in
func (ns *NatsStatus) PublishTemperature(temp *StatusMonitor.Temperature) {
	bytepackage, err := temp.PackageJSON()
	if err != nil {
		return
	}

	ns.NC.Publish(TS.PubTemp, bytepackage)

}

// WriteFormattedComm will intercept messages meant for the comm then format them as necessary
func (ns *NatsStatus) WriteFormattedComm(msg *nats.Msg) {
	line, err := CommRelayStructures.NewLineFromMsg(msg.Data)
	if err != nil {
		reply := DataStructures.ConstructNegativeReplyString(err.Error())
		DataStructures.PackageAndSendReplyString(reply, ns.NC, msg.Reply)
	}

	// If we are playing, then format the line
	if ns.RNode.StatusObserver.CurrentStatus.IsPlaying {
		//Send it to the formatter
		ns.CommRelay.FormatLine(line)
	} else {
		// If we are not playing then just send the line
		err = ns.SendComm(line.GetLine())
		if err != nil {
			reply := DataStructures.ConstructNegativeReplyString(err.Error())
			DataStructures.PackageAndSendReplyString(reply, ns.NC, msg.Reply)
		}
	}

}

// GetComm will recieve Comm Messages
func (ns *NatsStatus) GetComm(msg *nats.Msg) {
	cm := new(CS.CommMessage)
	err := proto.Unmarshal(msg.Data, cm)
	if err != nil {
		fmt.Println("Error could not get comm message!", err)
	}
	go ns.CommRelay.RecieveComm(cm.GetMessage())
	go ns.StatusMonitor.TempMonitor.UpdateTemperature(cm.GetMessage())
}

// RNode funcs

// RegisterFunctionstoRnode will register all functions to use with RNode then return an error
func (ns *NatsStatus) RegisterFunctionstoRnode() error {
	// Add Functions for status updates
	var err error
	_, err = ns.RNode.StatusObserver.AddFunction(PlayStructures.CANCEL, ns.GetCancelCommand)
	_, err = ns.RNode.StatusObserver.AddFunction(PlayStructures.DONE, ns.DonePrinting)
	_, err = ns.RNode.StatusObserver.AddFunction(PlayStructures.PLAY, ns.GetPlayCommand)
	return err
}

// DonePrinting will Capture the DONE status and then set up Status to return a done command
func (ns *NatsStatus) DonePrinting() {
	fmt.Println("Setting done notification!")
	ns.CommRelay.NotifyWhenEmpty()
}

// GetPlayCommand will get the play command and update CommRelay to start playing
func (ns *NatsStatus) GetPlayCommand() {
	ns.CommRelay.Playing = true
	ns.RNode.UpdateNode(PlayStructures.PLAY)
}

// GetCancelCommand will get the CANCEL status update and set up comm relay to cancel printing
func (ns *NatsStatus) GetCancelCommand() {
	ns.CommRelay.ResetLines()
	ns.RNode.UpdateNode(PlayStructures.CANCEL)
}
