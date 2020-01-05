/*
* @Author: Ximidar
* @Date:   2019-01-04 15:07:31
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-29 15:33:23
 */

package PlayStatus

import (
	"fmt"

	nats "github.com/nats-io/go-nats"
	"github.com/Ximidar/Flotilla/DataStructures"
	"github.com/Ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
)

// PlayStatusNats will act as the Nats connection for setting the play status.
// Usually this is handled by a single file, but the status is proving to be
// a difficult thing to keep track of / update
type PlayStatusNats struct {
	NC      *nats.Conn
	Control *Control
}

// NewPlayStatusNats will construct a new PlayStatusNats object
func NewPlayStatusNats(nc *nats.Conn) (*PlayStatusNats, error) {
	psn := new(PlayStatusNats)
	psn.NC = nc

	var err error
	psn.Control, err = NewControl(psn.PublishStatus)
	if err != nil {
		return nil, err
	}

	err = psn.setupSubscribers()
	if err != nil {
		return nil, err
	}

	return psn, nil
}

func (psn *PlayStatusNats) setupSubscribers() error {
	// PlayStatus
	_, err := psn.NC.Subscribe(PlayStructures.ControlAction, psn.RequestPlayStatus)
	_, err = psn.NC.Subscribe(PlayStructures.GetStatus, psn.GetStatus)
	_, err = psn.NC.Subscribe(PlayStructures.SetError, psn.SetError)

	// RollCall
	_, err = psn.NC.Subscribe(PlayStructures.RegisterNode, psn.RegisterNode)
	_, err = psn.NC.Subscribe(PlayStructures.RogerUp, psn.UpdateNode)

	return err
}

////////////////
// PlayStatus //
////////////////

// RequestPlayStatus will set the Play/Pause/Resume/Cancel flags and update the Nats server
// with the current status
func (psn *PlayStatusNats) RequestPlayStatus(msg *nats.Msg) {
	fmt.Println("Got request to Change Status", string(msg.Data))
	err := psn.Control.ProposeAction(string(msg.Data))
	if err != nil {
		fmt.Println("Could not set status correctly", err)
		reply := DataStructures.ConstructNegativeReplyString(err.Error())
		DataStructures.PackageAndSendReplyString(reply, psn.NC, msg.Reply)
		return
	}
	fmt.Println("Set Status Correctly")
	goodreply := DataStructures.ReplyString{Success: true, Message: "set"}
	fmt.Println(goodreply.Message, goodreply.Success)
	DataStructures.PackageAndSendReplyString(goodreply, psn.NC, msg.Reply)
}

// PublishStatus will publish the current Status to nats
func (psn *PlayStatusNats) PublishStatus(status []byte) {
	fmt.Println("Publishing New Status:", string(status))
	psn.NC.Publish(PlayStructures.PublishStatus, status)
}

// GetStatus will send back a json object with the current status
func (psn *PlayStatusNats) GetStatus(msg *nats.Msg) {
	psn.NC.Publish(msg.Reply, psn.Control.PlayStatus.GetCurrentJSON())
}

// SetError will accept the error message from Nats and send it to the play status object
func (psn *PlayStatusNats) SetError(msg *nats.Msg) {
	psn.Control.PlayStatus.SetError(string(msg.Data))
}

//////////////
// RollCall //
//////////////

// RegisterNode will register a node with RollCall
func (psn *PlayStatusNats) RegisterNode(msg *nats.Msg) {
	fmt.Println("Registering new node!")
	// Get the node to register
	node, err := PlayStructures.NewNodeFromMSG(msg)
	if err != nil {
		neg := DataStructures.ConstructNegativeReplyString(err.Error())
		DataStructures.PackageAndSendReplyString(neg, psn.NC, msg.Reply)
	}

	// Check node heartbeat
	_, err = psn.NC.Request(node.Name, []byte(nil), nats.DefaultTimeout)
	if err != nil {
		neg := DataStructures.ConstructNegativeReplyString("Could not get heartbeat")
		DataStructures.PackageAndSendReplyString(neg, psn.NC, msg.Reply)
	}

	// Register Node
	psn.Control.RollCall.RegisterNode(node.Name)

	// reply with success
	reply := DataStructures.ReplyString{}
	reply.Success = true
	reply.Message = fmt.Sprintf("%v has been registered", node.Name)

	DataStructures.PackageAndSendReplyString(reply, psn.NC, msg.Reply)
}

// UpdateNode will update the specified node's status
func (psn *PlayStatusNats) UpdateNode(msg *nats.Msg) {

	// Get the node to update
	node, err := PlayStructures.NewNodeFromMSG(msg)
	if err != nil {
		neg := DataStructures.ConstructNegativeReplyString(err.Error())
		DataStructures.PackageAndSendReplyString(neg, psn.NC, msg.Reply)
	}

	// Try to update node
	err = psn.Control.RollCall.RogerUp(node.Name, node.Status)
	if err != nil {
		neg := DataStructures.ConstructNegativeReplyString(err.Error())
		DataStructures.PackageAndSendReplyString(neg, psn.NC, msg.Reply)
	}

	// Reply Success
	reply := DataStructures.ReplyString{}
	reply.Success = true
	reply.Message = fmt.Sprintf("%v has been updated to status: %v", node.Name, node.Status)

	DataStructures.PackageAndSendReplyString(reply, psn.NC, msg.Reply)

}

// PingNodes will ping all registered nodes and update the node list if any don't respond back
func (psn *PlayStatusNats) PingNodes() {

	for key := range psn.Control.RollCall.Roll {
		_, err := psn.NC.Request(key, []byte(nil), nats.DefaultTimeout)
		if err != nil {
			if err == nats.ErrTimeout {
				defer psn.Control.RollCall.DeleteNode(key)
			} else {
				fmt.Println("Error with node:", key, err.Error())
			}

		}
	}

}
