// automatically generated by stateify.

package veth

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (v *vethPacket) StateTypeName() string {
	return "pkg/tcpip/link/veth.vethPacket"
}

func (v *vethPacket) StateFields() []string {
	return []string{
		"e",
		"protocol",
		"pkt",
	}
}

func (v *vethPacket) beforeSave() {}

// +checklocksignore
func (v *vethPacket) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.e)
	stateSinkObject.Save(1, &v.protocol)
	stateSinkObject.Save(2, &v.pkt)
}

func (v *vethPacket) afterLoad(context.Context) {}

// +checklocksignore
func (v *vethPacket) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.e)
	stateSourceObject.Load(1, &v.protocol)
	stateSourceObject.Load(2, &v.pkt)
}

func (e *Endpoint) StateTypeName() string {
	return "pkg/tcpip/link/veth.Endpoint"
}

func (e *Endpoint) StateFields() []string {
	return []string{
		"pair",
		"mtu",
		"backlogQueue",
		"dispatcher",
		"stack",
		"idx",
		"linkAddr",
	}
}

func (e *Endpoint) beforeSave() {}

// +checklocksignore
func (e *Endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.pair)
	stateSinkObject.Save(1, &e.mtu)
	stateSinkObject.Save(2, &e.backlogQueue)
	stateSinkObject.Save(3, &e.dispatcher)
	stateSinkObject.Save(4, &e.stack)
	stateSinkObject.Save(5, &e.idx)
	stateSinkObject.Save(6, &e.linkAddr)
}

func (e *Endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *Endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.pair)
	stateSourceObject.Load(1, &e.mtu)
	stateSourceObject.Load(2, &e.backlogQueue)
	stateSourceObject.Load(3, &e.dispatcher)
	stateSourceObject.Load(4, &e.stack)
	stateSourceObject.Load(5, &e.idx)
	stateSourceObject.Load(6, &e.linkAddr)
}

func init() {
	state.Register((*vethPacket)(nil))
	state.Register((*Endpoint)(nil))
}
