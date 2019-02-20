/*
* @Author: Ximidar
* @Date:   2019-02-15 16:06:10
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-20 14:47:13
 */

// Package NatsConnect will be used for standardizing connecting to NATS
package NatsConnect

import (
	"time"

	nats "github.com/nats-io/go-nats"
)

// AttemptConn will attempt to connect to a Nats Server using NatsConnect
func AttemptConn(natsAddress string, options ...nats.Option) (*nats.Conn, error) {
	attemptConn := new(NatsConnect)
	nc, err := attemptConn.Connect(natsAddress, options...)
	if err != nil {
		return nil, err
	}
	return nc, err
}

// NatsConnect will attempt to connect to a nats server and retry
type NatsConnect struct {
	NC      *nats.Conn
	address string
	options []nats.Option
}

// Connect will attempt to Connect to a Nats Server. Once the Connection has been made
// It will be passed back
func (nc *NatsConnect) Connect(natsAddress string, options ...nats.Option) (*nats.Conn, error) {
	// Store Variables
	nc.address = natsAddress
	nc.options = options

	// attempt a Connection to Nats
	var err error
	nc.NC, err = nats.Connect(nc.address, nc.options...)
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
		nc.NC, err = nats.Connect(nc.address, nc.options...)
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
