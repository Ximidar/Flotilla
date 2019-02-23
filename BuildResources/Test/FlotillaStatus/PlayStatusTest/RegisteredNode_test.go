/*
* @Author: Ximidar
* @Date:   2019-01-29 14:27:42
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-22 13:08:12
 */

package PlayStatusTest

import (
	"fmt"
	"testing"
	"time"

	"github.com/nats-io/gnatsd/test"
	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	"github.com/ximidar/Flotilla/DataStructures/StatusStructures/PlayStructures"
	"github.com/ximidar/Flotilla/FlotillaStatus/PlayStatus"
)

// Before running this test package you will need a nats server up

func TestNatsUp(t *testing.T) {
	server := test.RunDefaultServer()
	defer server.Shutdown()
	_, err := nats.Connect(nats.DefaultURL)
	CommonTestTools.CheckErr(t, "TestNatsUp Nats is not up", err)

}

func TestMakeRegisteredNode(t *testing.T) {
	server := test.RunDefaultServer()
	defer server.Shutdown()
	// Create Nats Conn to use
	NC, err := nats.Connect(nats.DefaultURL)
	CommonTestTools.CheckErr(t, "TestMakeRegisteredNode", err)
	defer NC.Close()

	// Make the nodes for registering
	node, err := MakeRNNatsNode(NC)
	CommonTestTools.CheckErr(t, "TestMakeRegisteredNode", err)
	// Ping Nodes so we don't get an error for initializing the node
	node.PingNodes()

	// Create a registered node
	RN, err := PlayStructures.NewRegisteredNode("TestNode", NC)
	CommonTestTools.CheckErr(t, "TestMakeRegisteredNode Making RN", err)
	fmt.Println(RN.Status)
}

func TestUpdateNodeStatus(t *testing.T) {
	server := test.RunDefaultServer()
	defer server.Shutdown()
	// Create Nats Conn to use
	NC, err := nats.Connect(nats.DefaultURL)
	CommonTestTools.CheckErr(t, "TestMakeRegisteredNode", err)
	defer NC.Close()

	// Make the nodes for registering
	node, err := MakeRNNatsNode(NC)
	CommonTestTools.CheckErr(t, "TestMakeRegisteredNode", err)
	// Ping Nodes so we don't get an error for initializing the node variable
	node.PingNodes()

	// Create a nodes on goroutines

	testNodeName := "TestNode"
	exitchan := make(chan bool, 20)
	exit := func() {
		exitchan <- true
	}
	defer exit()

	fmt.Println("Creating Nodes")
	for x := 0; x < 5; x++ {
		go CreateNodeGoroutine(fmt.Sprintf("%v%v", testNodeName, x), NC, exitchan)
	}
	fmt.Println("Finished creating nodes")
	<-time.After(500 * time.Millisecond)
	node.Control.RollCall.PrintNodeStatus()

	err = RequestStatusChange(PlayStructures.PLAY, NC)
	CommonTestTools.CheckErr(t, "Updating to PlayStatus", err)

	<-time.After(2 * time.Second)
	if node.Control.RollCall.AllNodesEqual(PlayStructures.PLAY) {
		fmt.Println("All nodes are equal!")
	} else {
		node.Control.RollCall.PrintNodeStatus()
		t.Fatal("All nodes are not equal")
	}

}

// Helpers

func MakeRNNatsNode(NC *nats.Conn) (*PlayStatus.PlayStatusNats, error) {
	psn, err := PlayStatus.NewPlayStatusNats(NC)
	return psn, err
}

func RequestStatusChange(status string, NC *nats.Conn) error {
	streamAction, err := PlayStructures.NewPlayAction(status)
	if err != nil {
		return err
	}

	err = streamAction.Send(NC)
	if err != nil {
		return err
	}
	return nil
}

func CreateNodeGoroutine(Name string, NC *nats.Conn, exit chan bool) {
	node, err := PlayStructures.NewRegisteredNode(Name, NC)
	if err != nil {
		fmt.Println("ERROR!", err)
	}
	fmt.Println(node.Status)
	select {
	case <-exit:
		fmt.Println("Exiting", Name)
	}
}
