package api

import (
	"fmt"

	CS "github.com/Ximidar/Flotilla/DataStructures/CommStructures"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

func (fw *FlotillaWeb) setupCommRelay() {
	_, err := fw.Nats.Subscribe(CS.ReadLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.ReadLine, err)
	}
	_, err = fw.Nats.Subscribe(CS.WriteLine, fw.CommRelay)
	if err != nil {
		fmt.Println("Could not subscribe to ", CS.WriteLine, err)
	}
}

// CommRelay will receive COMM messages from NATS
func (fw *FlotillaWeb) CommRelay(msg *nats.Msg) {
	cm := new(CS.CommMessage)
	err := proto.Unmarshal(msg.Data, cm)
	if err != nil {
		fmt.Println("Could not deconstruct proto message for commrelay")
	}
	fw.wsWrite <- []byte(cm.Message)
}
