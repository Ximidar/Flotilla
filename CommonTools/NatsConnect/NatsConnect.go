/*
* @Author: Ximidar
* @Date:   2019-02-15 16:06:10
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-22 15:30:06
 */

// Package NatsConnect will be used for standardizing connecting to NATS
package NatsConnect

import (
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	// DockerNATS is used by background tasks
	DockerNATS = "nats://0.0.0.0:4222"
	//DockerNATS = "nats:4222"

	// LocalNATS is used by programs that are not within docker
	LocalNATS = "nats://0.0.0.0:4222"
)

// DefaultConn will Create a NatsConnect object and not modify past the server URL and
// a name for the connection. This function will return a nats.Conn and an error
func DefaultConn(natsAddress string, name string) (*nats.Conn, error) {
	attemptConn, err := NewNatsConnect()
	if err != nil {
		return nil, err
	}
	attemptConn.Options.Url = natsAddress
	attemptConn.Options.Name = strings.ToLower(strings.Replace(name, ".", "", -1))

	fmt.Println("Attempting to connect to NATS on ", attemptConn.Options.Url, " with name ", attemptConn.Options.Name)
	nc, err := attemptConn.Connect()
	if err != nil {
		return nil, err
	}
	return nc, err
}

// NatsConnect will attempt to connect to a nats server and retry
type NatsConnect struct {
	NC      *nats.Conn
	Options nats.Options
}

// GetAvailableServers will list out the available servers
func (nc *NatsConnect) GetAvailableServers() {
	servers := nc.NC.Servers()
	for s := range servers {
		fmt.Println("Server Available at: ", s)
	}
}

// NewNatsConnect will create a NatsConnect Object
func NewNatsConnect() (*NatsConnect, error) {
	natsConn := new(NatsConnect)
	natsConn.Options = nats.GetDefaultOptions()
	natsConn.setOptions()
	natsConn.NC = nil

	return natsConn, nil

}

func (nc *NatsConnect) setOptions() {
	// Options
	nc.Options.AllowReconnect = true
	nc.Options.MaxReconnect = 10
	nc.Options.ReconnectWait = 5 * time.Second
	nc.Options.Timeout = 15 * time.Second

	// Callbacks
	nc.Options.ReconnectedCB = nc.ReconnectHandler
	nc.Options.ClosedCB = nc.ClosedHandler
	nc.Options.DisconnectedCB = nc.DisconnectHandler
}

// Connect will attempt to Connect to a Nats Server. Once the Connection has been made
// It will be passed back
func (nc *NatsConnect) Connect() (*nats.Conn, error) {

	// attempt a Connection to Nats
	var err error
	nc.NC, err = nc.Options.Connect()
	if err != nil {
		if nc.recoverable(err) {
			return nc.RetryNatsConn()
		}
		return nil, err
	}

	return nc.NC, nil
}

// RetryNatsConn will attempt to retry the connection five times
func (nc *NatsConnect) RetryNatsConn() (*nats.Conn, error) {
	var err error
	for i := 0; i < 5; i++ {
		// Wait 6 seconds before trying connection again
		<-time.After(6 * time.Second)
		nc.NC, err = nc.Options.Connect()
		if err == nil {
			return nc.NC, nil
		} // else try again
	}

	return nil, err
}

func (nc *NatsConnect) recoverable(err error) bool {
	switch err {
	case nats.ErrTimeout:
		return true
	case nats.ErrNoServers:
		return true
	case nil:
		return true
	default:
		return false
	}
}

// Connection Handlers

// DisconnectHandler will be fired off when the nats.Conn closes
func (nc *NatsConnect) DisconnectHandler(conn *nats.Conn) {
	fmt.Println("NATS Disconnected due to:", conn.LastError())

}

// ClosedHandler will be fired off when the nats.Conn is closed
func (nc *NatsConnect) ClosedHandler(conn *nats.Conn) {
	fmt.Println("NATS Closed due to:", conn.LastError())

}

// ReconnectHandler will be fired off when the nats.Conn is reconnected
func (nc *NatsConnect) ReconnectHandler(conn *nats.Conn) {
	fmt.Println("NATS Reconnected")
}
