/*
* @Author: Ximidar
* @Date:   2018-07-28 11:10:37
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-28 13:14:32
 */

package NatsConn

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	commango "github.com/Ximidar/Flotilla/Commango/comm"
	"github.com/Ximidar/Flotilla/CommonTools/NatsConnect"
	DS "github.com/Ximidar/Flotilla/DataStructures"
	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	proto "github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

// TermChannel will monitor for an exit signal
var TermChannel chan os.Signal

// NatsConn is the Comm interface to the Nats Server
type NatsConn struct {
	NC *nats.Conn

	//comm connection
	Comm *commango.Comm
}

// NewNatsConn will construct a NatsConn object
func NewNatsConn() *NatsConn {

	gnats := new(NatsConn)
	var err error
	gnats.NC, err = NatsConnect.DefaultConn(NatsConnect.DockerNATS, CS.Name)

	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	gnats.Comm = commango.NewComm(gnats)
	gnats.createReqReplies()

	return gnats
}

// Serve will keep the program open
func (gnats *NatsConn) Serve() {
	// Function for waiting for exit on the main loop
	// Wait for termination
	TermChannel = make(chan os.Signal)
	signal.Notify(TermChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Commango Started")
	<-TermChannel
	fmt.Println("Recieved Interrupt Sig, Now Exiting.")
	os.Exit(1)

}

func (gnats *NatsConn) createReqReplies() (err error) {
	// req replies
	_, err = gnats.NC.Subscribe(CS.ListPorts, gnats.listPorts)
	_, err = gnats.NC.Subscribe(CS.InitializeComm, gnats.initComm)
	_, err = gnats.NC.Subscribe(CS.ConnectComm, gnats.connectComm)
	_, err = gnats.NC.Subscribe(CS.DisconnectComm, gnats.disconnectComm)
	_, err = gnats.NC.Subscribe(CS.WriteComm, gnats.writeComm)
	_, err = gnats.NC.Subscribe(CS.GetStatus, gnats.getStatus)

	return err
}

func (gnats *NatsConn) listPorts(msg *nats.Msg) {

	ports, err := gnats.Comm.GetAvailablePorts()
	if err != nil {
		reply := DS.ConstructNegativeReplyString(err.Error())
		// Marshal reply
		mreply, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, mreply)
		return
	}

	// Marshal ports
	mReply, err := proto.Marshal(ports)

	// Publish
	gnats.NC.Publish(msg.Reply, mReply)
}

func (gnats *NatsConn) getStatus(msg *nats.Msg) {
	status := gnats.Comm.GetCommStatus()
	mdata, err := proto.Marshal(status)
	if err != nil {
		fmt.Println("Could not marshal status", status)
		return
	}
	gnats.NC.Publish(msg.Reply, mdata)
}

// PublishStatus will publish the Comm status to the Nats Server
func (gnats *NatsConn) PublishStatus(status *CS.CommStatus) error {
	reply := new(DS.ReplyJSON)
	reply.Success = true
	byteStatus, err := proto.Marshal(status)
	if err != nil {
		fmt.Println("Could not marshal the status into a protobuffer")
	}

	err = gnats.NC.Publish(CS.StatusUpdate, byteStatus)
	if err != nil {
		fmt.Println("Could not publish status")
		return err
	}
	return nil

}

func (gnats *NatsConn) initComm(msg *nats.Msg) {

	reply := new(DS.ReplyString)
	init := new(CS.InitComm)
	err := proto.Unmarshal(msg.Data, init)

	// error out if we cannot unmarshal the data
	if err != nil {
		reply.Success = false
		reply.Message = "could not unmarshal data"
		repByte, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, repByte)
		return
	}

	// error out if we cannot initialize the comm
	err = gnats.Comm.InitComm(init)
	if err != nil {
		reply.Success = false
		reply.Message = "Could Not Initialize Comm: " + err.Error()
		repByte, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, repByte)
		return
	}

	// Create success response and send
	reply.Success = true
	reply.Message = "Comm Initialized"
	mReply, err := json.Marshal(reply)
	if err != nil {
		panic(err)
	} // There should be no reason it can't marshal
	gnats.NC.Publish(msg.Reply, mReply)

}

func (gnats *NatsConn) connectComm(msg *nats.Msg) {
	err := gnats.Comm.OpenComm()
	reply := new(DS.ReplyString)
	if err != nil {
		reply.Success = false
		reply.Message = err.Error()
		rep, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, rep)
		return
	}
	reply.Success = true
	reply.Message = "Connected"
	mReply, err := json.Marshal(reply)
	gnats.NC.Publish(msg.Reply, mReply)
}

func (gnats *NatsConn) disconnectComm(msg *nats.Msg) {
	err := gnats.Comm.CloseComm()
	reply := new(DS.ReplyString)
	if err != nil {
		reply.Success = false
		reply.Message = err.Error()
		rep, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, rep)
		return
	}
	reply.Success = true
	reply.Message = "Disconnected"
	mReply, _ := json.Marshal(reply)
	gnats.NC.Publish(msg.Reply, mReply)
}

func (gnats *NatsConn) constructCommMessage(cmb []byte) (*CS.CommMessage, error) {
	cm := new(CS.CommMessage)
	err := proto.Unmarshal(cmb, cm)
	return cm, err
}

func (gnats *NatsConn) constructWriteReceipt(bytesWritter int) (wr *CS.WrittenBytes) {
	wr = new(CS.WrittenBytes)
	wr.Bytes = int32(bytesWritter)
	return
}

func (gnats *NatsConn) writeComm(msg *nats.Msg) {
	// Get Message protobuf
	message, err := gnats.constructCommMessage(msg.Data)
	if err != nil {
		fmt.Println("Could not deconstruct message")
		reply := DS.ConstructNegativeReplyString(err.Error())
		// Marshal reply
		mreply, _ := json.Marshal(reply)
		gnats.NC.Publish(msg.Reply, mreply)
		return
	}
	mess := message.GetMessage()
	bytesWritten, err := gnats.Comm.WriteComm(mess)
	reply := gnats.constructWriteReceipt(bytesWritten)

	mReply, err := proto.Marshal(reply)
	if err != nil {
		fmt.Println("Could not marshal Reply Receipt for message", message, reply)
		return
	}

	gnats.NC.Publish(msg.Reply, mReply)
}

// ReadLineEmitter will publish any Read lines from Comm
func (gnats *NatsConn) ReadLineEmitter(line string) error {
	brl, err := gnats.marshalCommMessage(line)
	if err != nil {
		return err
	}
	err = gnats.NC.Publish(CS.ReadLine, brl)
	if err != nil {
		fmt.Println("Could not send Read Line", line)
		return err
	}
	return nil
}

// WriteLineEmitter will publish any Written lines to Comm
func (gnats *NatsConn) WriteLineEmitter(line string) error {
	bwl, err := gnats.marshalCommMessage(line)
	if err != nil {
		return err
	}
	err = gnats.NC.Publish(CS.WriteLine, bwl)
	if err != nil {
		fmt.Println("Could not send Written line", line)
		return err
	}
	return nil
}

func (gnats *NatsConn) marshalCommMessage(line string) ([]byte, error) {
	cm := new(CS.CommMessage)
	cm.Message = line
	bcm, err := proto.Marshal(cm)
	if err != nil {
		fmt.Println("Could not marshal message", line)
		return []byte(nil), err
	}
	return bcm, nil
}
