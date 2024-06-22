// automatically generated by stateify.

//go:build linux && linux && linux && linux && linux && linux
// +build linux,linux,linux,linux,linux,linux

package sharedmem

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (q *QueueConfig) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem.QueueConfig"
}

func (q *QueueConfig) StateFields() []string {
	return []string{
		"DataFD",
		"EventFD",
		"TxPipeFD",
		"RxPipeFD",
		"SharedDataFD",
	}
}

func (q *QueueConfig) beforeSave() {}

// +checklocksignore
func (q *QueueConfig) StateSave(stateSinkObject state.Sink) {
	q.beforeSave()
	stateSinkObject.Save(0, &q.DataFD)
	stateSinkObject.Save(1, &q.EventFD)
	stateSinkObject.Save(2, &q.TxPipeFD)
	stateSinkObject.Save(3, &q.RxPipeFD)
	stateSinkObject.Save(4, &q.SharedDataFD)
}

func (q *QueueConfig) afterLoad(context.Context) {}

// +checklocksignore
func (q *QueueConfig) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &q.DataFD)
	stateSourceObject.Load(1, &q.EventFD)
	stateSourceObject.Load(2, &q.TxPipeFD)
	stateSourceObject.Load(3, &q.RxPipeFD)
	stateSourceObject.Load(4, &q.SharedDataFD)
}

func (o *Options) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem.Options"
}

func (o *Options) StateFields() []string {
	return []string{
		"MTU",
		"BufferSize",
		"LinkAddress",
		"TX",
		"RX",
		"PeerFD",
		"OnClosed",
		"TXChecksumOffload",
		"RXChecksumOffload",
		"VirtioNetHeaderRequired",
		"GSOMaxSize",
	}
}

func (o *Options) beforeSave() {}

// +checklocksignore
func (o *Options) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.MTU)
	stateSinkObject.Save(1, &o.BufferSize)
	stateSinkObject.Save(2, &o.LinkAddress)
	stateSinkObject.Save(3, &o.TX)
	stateSinkObject.Save(4, &o.RX)
	stateSinkObject.Save(5, &o.PeerFD)
	stateSinkObject.Save(6, &o.OnClosed)
	stateSinkObject.Save(7, &o.TXChecksumOffload)
	stateSinkObject.Save(8, &o.RXChecksumOffload)
	stateSinkObject.Save(9, &o.VirtioNetHeaderRequired)
	stateSinkObject.Save(10, &o.GSOMaxSize)
}

func (o *Options) afterLoad(context.Context) {}

// +checklocksignore
func (o *Options) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.MTU)
	stateSourceObject.Load(1, &o.BufferSize)
	stateSourceObject.Load(2, &o.LinkAddress)
	stateSourceObject.Load(3, &o.TX)
	stateSourceObject.Load(4, &o.RX)
	stateSourceObject.Load(5, &o.PeerFD)
	stateSourceObject.Load(6, &o.OnClosed)
	stateSourceObject.Load(7, &o.TXChecksumOffload)
	stateSourceObject.Load(8, &o.RXChecksumOffload)
	stateSourceObject.Load(9, &o.VirtioNetHeaderRequired)
	stateSourceObject.Load(10, &o.GSOMaxSize)
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/link/sharedmem.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"mtu",
		"bufferSize",
		"peerFD",
		"caps",
		"hdrSize",
		"gsoMaxSize",
		"virtioNetHeaderRequired",
		"rx",
		"stopRequested",
		"completed",
		"tx",
		"workerStarted",
		"addr",
	}
}

func (e *endpoint) beforeSave() {}

// +checklocksignore
func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.mtu)
	stateSinkObject.Save(1, &e.bufferSize)
	stateSinkObject.Save(2, &e.peerFD)
	stateSinkObject.Save(3, &e.caps)
	stateSinkObject.Save(4, &e.hdrSize)
	stateSinkObject.Save(5, &e.gsoMaxSize)
	stateSinkObject.Save(6, &e.virtioNetHeaderRequired)
	stateSinkObject.Save(7, &e.rx)
	stateSinkObject.Save(8, &e.stopRequested)
	stateSinkObject.Save(9, &e.completed)
	stateSinkObject.Save(10, &e.tx)
	stateSinkObject.Save(11, &e.workerStarted)
	stateSinkObject.Save(12, &e.addr)
}

func (e *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.mtu)
	stateSourceObject.Load(1, &e.bufferSize)
	stateSourceObject.Load(2, &e.peerFD)
	stateSourceObject.Load(3, &e.caps)
	stateSourceObject.Load(4, &e.hdrSize)
	stateSourceObject.Load(5, &e.gsoMaxSize)
	stateSourceObject.Load(6, &e.virtioNetHeaderRequired)
	stateSourceObject.Load(7, &e.rx)
	stateSourceObject.Load(8, &e.stopRequested)
	stateSourceObject.Load(9, &e.completed)
	stateSourceObject.Load(10, &e.tx)
	stateSourceObject.Load(11, &e.workerStarted)
	stateSourceObject.Load(12, &e.addr)
}

func init() {
	state.Register((*QueueConfig)(nil))
	state.Register((*Options)(nil))
	state.Register((*endpoint)(nil))
}
