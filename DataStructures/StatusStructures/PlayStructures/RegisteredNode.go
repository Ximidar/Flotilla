/*
* @Author: Ximidar
* @Date:   2019-01-16 15:24:52
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-20 14:48:43
 */

package PlayStructures

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Ximidar/Flotilla/DataStructures"
	"github.com/nats-io/nats.go"
)

// Node just keeps track of the node name and status
type Node struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// NewNodeFromMSG will create a Node object from a nats MSG
func NewNodeFromMSG(msg *nats.Msg) (*Node, error) {
	node := new(Node)
	err := json.Unmarshal(msg.Data, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// RegisteredNode will register a node that needs to update it's status with
// the PlayStatus
type RegisteredNode struct {
	Node
	StatusObserver *StatusObserver
	NC             *nats.Conn
}

// NewRegisteredNode will create a RegisteredNode struct
func NewRegisteredNode(name string, nc *nats.Conn) (*RegisteredNode, error) {
	rn := new(RegisteredNode)

	rn.NC = nc
	rn.Name = name
	rn.Status = IDLE

	// Make Status Observer
	var err error
	rn.StatusObserver, err = NewStatusObserver()
	if err != nil {
		return nil, err
	}

	// Connect Status Observer to nats
	_, err = rn.NC.Subscribe(PublishStatus, rn.StatusObserver.UpdateStatusFromNats)
	if err != nil {
		return nil, err
	}

	// Add in the callback for any unregistered callbacks
	rn.StatusObserver.AddUknownStatusFunc(rn.UknownStatusCallback)

	// Register Node
	err = rn.registerWithPlayStatus()
	if err != nil {
		return nil, err
	}

	return rn, nil
}

func (rn *RegisteredNode) marshalNode() ([]byte, error) {
	marshaled, err := json.Marshal(&rn.Node)
	if err != nil {
		return nil, err
	}
	return marshaled, nil
}

func (rn *RegisteredNode) registerWithPlayStatus() error {
	// Create request package
	marshed, err := rn.marshalNode()
	if err != nil {
		return err
	}

	// Subscribe HeartBeat to the name of the Node
	fmt.Println("Registering Name:", rn.Name)
	_, err = rn.NC.Subscribe(rn.Name, rn.HeartBeat)

	if err != nil {
		fmt.Printf("Could not subscribe %v for requests: %v", rn.Name, err.Error())
		return err
	}

	// Request register
	fmt.Println("Registering node using subject", RegisterNode)
	fmt.Println("Sendng them", string(marshed))
	reply, err := rn.NC.Request(RegisterNode, marshed, 5*time.Second)
	if err != nil {
		fmt.Println("ERROR Could not register node over nats", err.Error())
		return err
	}
	fmt.Println(string(reply.Data))
	uReply, err := DataStructures.NewReplyStringFromMSG(reply)
	if err != nil {
		fmt.Println("ERROR could not construct reply from message", err.Error())
		return err
	}

	if !uReply.Success {
		return fmt.Errorf("Could not register node because reply was negative: %v", err.Error())
	}

	return nil
}

// UpdateNode will update PlayStatus with the node's current status Change
func (rn *RegisteredNode) UpdateNode(status string) error {

	err := CheckSubj(status)
	if err != nil {
		return err
	}

	rn.Node.Status = status

	// Get JSON for status
	byteNode, err := rn.marshalNode()
	if err != nil {
		return err
	}

	// Request a statusChange
	reply, err := rn.NC.Request(RogerUp, byteNode, 5*time.Second)
	if err != nil {
		fmt.Println("Cannot Request", RogerUp, "Because:", err)
	}

	uReply, err := DataStructures.NewReplyStringFromMSG(reply)

	if err != nil || !uReply.Success {
		return fmt.Errorf("Could not update node: %v", err.Error())
	}

	return nil

}

// HeartBeat will ping back a response to PlayStatus when it asks.
// This will stop responding when the program turns off
func (rn *RegisteredNode) HeartBeat(msg *nats.Msg) {
	if msg.Reply != "" {
		err := rn.NC.Publish(msg.Reply, []byte("ok!"))
		if err != nil {
			fmt.Println("HEART BEAT ERROR:", err)
		}
	}

}

// UknownStatusCallback is a callback for updating the node when the statusObserver does not
// have a registered callback for the supplied action. This callback is used so that the entire
// system will not hang if every node doesn't handle every status call
func (rn *RegisteredNode) UknownStatusCallback(action string) {
	err := rn.UpdateNode(action)
	if err != nil {
		fmt.Println("UKNOWN STATUS ERROR:", err)
	}
}
