/*
* @Author: Ximidar
* @Date:   2019-02-15 14:47:34
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-20 14:41:59
 */

package version

import (
	"encoding/json"
	"strings"

	nats "github.com/nats-io/go-nats"
)

// Nats will return a version JSON package of the version when queried over NATS
type Nats struct {
	Version      string `json:"version"`
	CompiledBy   string `json:"author"`
	CompiledDate string `json:"compiledDate"`
	CommitHash   string `json:"commitHash"`

	nc          *nats.Conn
	requestName string
}

// NewNats will return a Nats object
func NewNats(nc *nats.Conn, Name string) (*Nats, error) {

	nv := new(Nats)

	// Check for suffix having a period. Subjects apparently cannot have 2 periods in them
	if strings.HasSuffix(Name, ".") {
		nv.requestName = Name + "Version"
	} else {
		nv.requestName = Name + ".Version"
	}

	nv.nc = nc

	// Subscribe to name events
	_, err := nv.nc.Subscribe(nv.requestName, nv.returnVersion)

	return nv, err

}

// NewNatsFromMsg will return a Nats object
func NewNatsFromMsg(msg *nats.Msg) (*Nats, error) {
	nv := new(Nats)
	err := json.Unmarshal(msg.Data, nv)

	if err != nil {
		return nil, err
	}
	return nv, nil
}

// ReturnVersion will return a JSON string to the reply
func (nv *Nats) returnVersion(msg *nats.Msg) {
	if nv.nc == nil {
		return
	}
	nv.Version = Version
	nv.CompiledBy = CompiledBy
	nv.CompiledDate = CompiledDate
	nv.CommitHash = CommitHash

	mVersion, _ := json.Marshal(nv)
	nv.nc.Publish(msg.Reply, mVersion)
}
