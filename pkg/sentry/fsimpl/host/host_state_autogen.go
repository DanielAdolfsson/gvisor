// automatically generated by stateify.

package host

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (r *ConnectedEndpointRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.ConnectedEndpointRefs"
}

func (r *ConnectedEndpointRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *ConnectedEndpointRefs) beforeSave() {}

func (r *ConnectedEndpointRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

func (r *ConnectedEndpointRefs) afterLoad() {}

func (r *ConnectedEndpointRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
}

func (f *filesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.filesystemType"
}

func (f *filesystemType) StateFields() []string {
	return []string{}
}

func (f *filesystemType) beforeSave() {}

func (f *filesystemType) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *filesystemType) afterLoad() {}

func (f *filesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
	}
}

func (fs *filesystem) beforeSave() {}

func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
}

func (fs *filesystem) afterLoad() {}

func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"InodeNoStatFS",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"InodeTemporary",
		"locks",
		"inodeRefs",
		"hostFD",
		"ino",
		"isTTY",
		"seekable",
		"wouldBlock",
		"queue",
		"canMap",
		"mappings",
		"pf",
	}
}

func (i *inode) beforeSave() {}

func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeNoStatFS)
	stateSinkObject.Save(1, &i.InodeNotDirectory)
	stateSinkObject.Save(2, &i.InodeNotSymlink)
	stateSinkObject.Save(3, &i.InodeTemporary)
	stateSinkObject.Save(4, &i.locks)
	stateSinkObject.Save(5, &i.inodeRefs)
	stateSinkObject.Save(6, &i.hostFD)
	stateSinkObject.Save(7, &i.ino)
	stateSinkObject.Save(8, &i.isTTY)
	stateSinkObject.Save(9, &i.seekable)
	stateSinkObject.Save(10, &i.wouldBlock)
	stateSinkObject.Save(11, &i.queue)
	stateSinkObject.Save(12, &i.canMap)
	stateSinkObject.Save(13, &i.mappings)
	stateSinkObject.Save(14, &i.pf)
}

func (i *inode) afterLoad() {}

func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeNoStatFS)
	stateSourceObject.Load(1, &i.InodeNotDirectory)
	stateSourceObject.Load(2, &i.InodeNotSymlink)
	stateSourceObject.Load(3, &i.InodeTemporary)
	stateSourceObject.Load(4, &i.locks)
	stateSourceObject.Load(5, &i.inodeRefs)
	stateSourceObject.Load(6, &i.hostFD)
	stateSourceObject.Load(7, &i.ino)
	stateSourceObject.Load(8, &i.isTTY)
	stateSourceObject.Load(9, &i.seekable)
	stateSourceObject.Load(10, &i.wouldBlock)
	stateSourceObject.Load(11, &i.queue)
	stateSourceObject.Load(12, &i.canMap)
	stateSourceObject.Load(13, &i.mappings)
	stateSourceObject.Load(14, &i.pf)
}

func (f *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.fileDescription"
}

func (f *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"LockFD",
		"inode",
		"offset",
	}
}

func (f *fileDescription) beforeSave() {}

func (f *fileDescription) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.vfsfd)
	stateSinkObject.Save(1, &f.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &f.LockFD)
	stateSinkObject.Save(3, &f.inode)
	stateSinkObject.Save(4, &f.offset)
}

func (f *fileDescription) afterLoad() {}

func (f *fileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.vfsfd)
	stateSourceObject.Load(1, &f.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &f.LockFD)
	stateSourceObject.Load(3, &f.inode)
	stateSourceObject.Load(4, &f.offset)
}

func (r *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.inodeRefs"
}

func (r *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *inodeRefs) beforeSave() {}

func (r *inodeRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

func (r *inodeRefs) afterLoad() {}

func (r *inodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
}

func (i *inodePlatformFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.inodePlatformFile"
}

func (i *inodePlatformFile) StateFields() []string {
	return []string{
		"inode",
		"fdRefs",
		"fileMapper",
	}
}

func (i *inodePlatformFile) beforeSave() {}

func (i *inodePlatformFile) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.inode)
	stateSinkObject.Save(1, &i.fdRefs)
	stateSinkObject.Save(2, &i.fileMapper)
}

func (i *inodePlatformFile) afterLoad() {}

func (i *inodePlatformFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.inode)
	stateSourceObject.Load(1, &i.fdRefs)
	stateSourceObject.Load(2, &i.fileMapper)
}

func (c *ConnectedEndpoint) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.ConnectedEndpoint"
}

func (c *ConnectedEndpoint) StateFields() []string {
	return []string{
		"ConnectedEndpointRefs",
		"fd",
		"addr",
		"stype",
	}
}

func (c *ConnectedEndpoint) beforeSave() {}

func (c *ConnectedEndpoint) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.ConnectedEndpointRefs)
	stateSinkObject.Save(1, &c.fd)
	stateSinkObject.Save(2, &c.addr)
	stateSinkObject.Save(3, &c.stype)
}

func (c *ConnectedEndpoint) afterLoad() {}

func (c *ConnectedEndpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.ConnectedEndpointRefs)
	stateSourceObject.Load(1, &c.fd)
	stateSourceObject.Load(2, &c.addr)
	stateSourceObject.Load(3, &c.stype)
}

func (t *TTYFileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/host.TTYFileDescription"
}

func (t *TTYFileDescription) StateFields() []string {
	return []string{
		"fileDescription",
		"session",
		"fgProcessGroup",
		"termios",
	}
}

func (t *TTYFileDescription) beforeSave() {}

func (t *TTYFileDescription) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.fileDescription)
	stateSinkObject.Save(1, &t.session)
	stateSinkObject.Save(2, &t.fgProcessGroup)
	stateSinkObject.Save(3, &t.termios)
}

func (t *TTYFileDescription) afterLoad() {}

func (t *TTYFileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.fileDescription)
	stateSourceObject.Load(1, &t.session)
	stateSourceObject.Load(2, &t.fgProcessGroup)
	stateSourceObject.Load(3, &t.termios)
}

func init() {
	state.Register((*ConnectedEndpointRefs)(nil))
	state.Register((*filesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inode)(nil))
	state.Register((*fileDescription)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*inodePlatformFile)(nil))
	state.Register((*ConnectedEndpoint)(nil))
	state.Register((*TTYFileDescription)(nil))
}
