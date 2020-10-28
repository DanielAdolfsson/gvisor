// automatically generated by stateify.

package kernfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (l *dentryList) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.dentryList"
}

func (l *dentryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *dentryList) beforeSave() {}

func (l *dentryList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *dentryList) afterLoad() {}

func (l *dentryList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *dentryEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.dentryEntry"
}

func (e *dentryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *dentryEntry) beforeSave() {}

func (e *dentryEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *dentryEntry) afterLoad() {}

func (e *dentryEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (f *DynamicBytesFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.DynamicBytesFile"
}

func (f *DynamicBytesFile) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoStatFS",
		"InodeNoopRefCount",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"locks",
		"data",
	}
}

func (f *DynamicBytesFile) beforeSave() {}

func (f *DynamicBytesFile) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.InodeAttrs)
	stateSinkObject.Save(1, &f.InodeNoStatFS)
	stateSinkObject.Save(2, &f.InodeNoopRefCount)
	stateSinkObject.Save(3, &f.InodeNotDirectory)
	stateSinkObject.Save(4, &f.InodeNotSymlink)
	stateSinkObject.Save(5, &f.locks)
	stateSinkObject.Save(6, &f.data)
}

func (f *DynamicBytesFile) afterLoad() {}

func (f *DynamicBytesFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.InodeAttrs)
	stateSourceObject.Load(1, &f.InodeNoStatFS)
	stateSourceObject.Load(2, &f.InodeNoopRefCount)
	stateSourceObject.Load(3, &f.InodeNotDirectory)
	stateSourceObject.Load(4, &f.InodeNotSymlink)
	stateSourceObject.Load(5, &f.locks)
	stateSourceObject.Load(6, &f.data)
}

func (fd *DynamicBytesFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.DynamicBytesFD"
}

func (fd *DynamicBytesFD) StateFields() []string {
	return []string{
		"FileDescriptionDefaultImpl",
		"DynamicBytesFileDescriptionImpl",
		"LockFD",
		"vfsfd",
		"inode",
	}
}

func (fd *DynamicBytesFD) beforeSave() {}

func (fd *DynamicBytesFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(1, &fd.DynamicBytesFileDescriptionImpl)
	stateSinkObject.Save(2, &fd.LockFD)
	stateSinkObject.Save(3, &fd.vfsfd)
	stateSinkObject.Save(4, &fd.inode)
}

func (fd *DynamicBytesFD) afterLoad() {}

func (fd *DynamicBytesFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(1, &fd.DynamicBytesFileDescriptionImpl)
	stateSourceObject.Load(2, &fd.LockFD)
	stateSourceObject.Load(3, &fd.vfsfd)
	stateSourceObject.Load(4, &fd.inode)
}

func (s *SeekEndConfig) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.SeekEndConfig"
}

func (s *SeekEndConfig) StateFields() []string {
	return nil
}

func (g *GenericDirectoryFDOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.GenericDirectoryFDOptions"
}

func (g *GenericDirectoryFDOptions) StateFields() []string {
	return []string{
		"SeekEnd",
	}
}

func (g *GenericDirectoryFDOptions) beforeSave() {}

func (g *GenericDirectoryFDOptions) StateSave(stateSinkObject state.Sink) {
	g.beforeSave()
	stateSinkObject.Save(0, &g.SeekEnd)
}

func (g *GenericDirectoryFDOptions) afterLoad() {}

func (g *GenericDirectoryFDOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &g.SeekEnd)
}

func (fd *GenericDirectoryFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.GenericDirectoryFD"
}

func (fd *GenericDirectoryFD) StateFields() []string {
	return []string{
		"FileDescriptionDefaultImpl",
		"DirectoryFileDescriptionDefaultImpl",
		"LockFD",
		"seekEnd",
		"vfsfd",
		"children",
		"off",
	}
}

func (fd *GenericDirectoryFD) beforeSave() {}

func (fd *GenericDirectoryFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.LockFD)
	stateSinkObject.Save(3, &fd.seekEnd)
	stateSinkObject.Save(4, &fd.vfsfd)
	stateSinkObject.Save(5, &fd.children)
	stateSinkObject.Save(6, &fd.off)
}

func (fd *GenericDirectoryFD) afterLoad() {}

func (fd *GenericDirectoryFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(1, &fd.DirectoryFileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.LockFD)
	stateSourceObject.Load(3, &fd.seekEnd)
	stateSourceObject.Load(4, &fd.vfsfd)
	stateSourceObject.Load(5, &fd.children)
	stateSourceObject.Load(6, &fd.off)
}

func (i *InodeNoopRefCount) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeNoopRefCount"
}

func (i *InodeNoopRefCount) StateFields() []string {
	return []string{
		"InodeTemporary",
	}
}

func (i *InodeNoopRefCount) beforeSave() {}

func (i *InodeNoopRefCount) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeTemporary)
}

func (i *InodeNoopRefCount) afterLoad() {}

func (i *InodeNoopRefCount) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeTemporary)
}

func (i *InodeDirectoryNoNewChildren) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeDirectoryNoNewChildren"
}

func (i *InodeDirectoryNoNewChildren) StateFields() []string {
	return []string{}
}

func (i *InodeDirectoryNoNewChildren) beforeSave() {}

func (i *InodeDirectoryNoNewChildren) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *InodeDirectoryNoNewChildren) afterLoad() {}

func (i *InodeDirectoryNoNewChildren) StateLoad(stateSourceObject state.Source) {
}

func (i *InodeNotDirectory) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeNotDirectory"
}

func (i *InodeNotDirectory) StateFields() []string {
	return []string{
		"InodeAlwaysValid",
	}
}

func (i *InodeNotDirectory) beforeSave() {}

func (i *InodeNotDirectory) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeAlwaysValid)
}

func (i *InodeNotDirectory) afterLoad() {}

func (i *InodeNotDirectory) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeAlwaysValid)
}

func (i *InodeNotSymlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeNotSymlink"
}

func (i *InodeNotSymlink) StateFields() []string {
	return []string{}
}

func (i *InodeNotSymlink) beforeSave() {}

func (i *InodeNotSymlink) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *InodeNotSymlink) afterLoad() {}

func (i *InodeNotSymlink) StateLoad(stateSourceObject state.Source) {
}

func (a *InodeAttrs) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeAttrs"
}

func (a *InodeAttrs) StateFields() []string {
	return []string{
		"devMajor",
		"devMinor",
		"ino",
		"mode",
		"uid",
		"gid",
		"nlink",
		"blockSize",
		"atime",
		"mtime",
		"ctime",
	}
}

func (a *InodeAttrs) beforeSave() {}

func (a *InodeAttrs) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.devMajor)
	stateSinkObject.Save(1, &a.devMinor)
	stateSinkObject.Save(2, &a.ino)
	stateSinkObject.Save(3, &a.mode)
	stateSinkObject.Save(4, &a.uid)
	stateSinkObject.Save(5, &a.gid)
	stateSinkObject.Save(6, &a.nlink)
	stateSinkObject.Save(7, &a.blockSize)
	stateSinkObject.Save(8, &a.atime)
	stateSinkObject.Save(9, &a.mtime)
	stateSinkObject.Save(10, &a.ctime)
}

func (a *InodeAttrs) afterLoad() {}

func (a *InodeAttrs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.devMajor)
	stateSourceObject.Load(1, &a.devMinor)
	stateSourceObject.Load(2, &a.ino)
	stateSourceObject.Load(3, &a.mode)
	stateSourceObject.Load(4, &a.uid)
	stateSourceObject.Load(5, &a.gid)
	stateSourceObject.Load(6, &a.nlink)
	stateSourceObject.Load(7, &a.blockSize)
	stateSourceObject.Load(8, &a.atime)
	stateSourceObject.Load(9, &a.mtime)
	stateSourceObject.Load(10, &a.ctime)
}

func (s *slot) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.slot"
}

func (s *slot) StateFields() []string {
	return []string{
		"name",
		"inode",
		"static",
		"slotEntry",
	}
}

func (s *slot) beforeSave() {}

func (s *slot) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.name)
	stateSinkObject.Save(1, &s.inode)
	stateSinkObject.Save(2, &s.static)
	stateSinkObject.Save(3, &s.slotEntry)
}

func (s *slot) afterLoad() {}

func (s *slot) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.name)
	stateSourceObject.Load(1, &s.inode)
	stateSourceObject.Load(2, &s.static)
	stateSourceObject.Load(3, &s.slotEntry)
}

func (o *OrderedChildrenOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.OrderedChildrenOptions"
}

func (o *OrderedChildrenOptions) StateFields() []string {
	return []string{
		"Writable",
	}
}

func (o *OrderedChildrenOptions) beforeSave() {}

func (o *OrderedChildrenOptions) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.Writable)
}

func (o *OrderedChildrenOptions) afterLoad() {}

func (o *OrderedChildrenOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.Writable)
}

func (o *OrderedChildren) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.OrderedChildren"
}

func (o *OrderedChildren) StateFields() []string {
	return []string{
		"writable",
		"order",
		"set",
	}
}

func (o *OrderedChildren) beforeSave() {}

func (o *OrderedChildren) StateSave(stateSinkObject state.Sink) {
	o.beforeSave()
	stateSinkObject.Save(0, &o.writable)
	stateSinkObject.Save(1, &o.order)
	stateSinkObject.Save(2, &o.set)
}

func (o *OrderedChildren) afterLoad() {}

func (o *OrderedChildren) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &o.writable)
	stateSourceObject.Load(1, &o.order)
	stateSourceObject.Load(2, &o.set)
}

func (r *renameAcrossDifferentImplementationsError) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.renameAcrossDifferentImplementationsError"
}

func (r *renameAcrossDifferentImplementationsError) StateFields() []string {
	return []string{}
}

func (r *renameAcrossDifferentImplementationsError) beforeSave() {}

func (r *renameAcrossDifferentImplementationsError) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
}

func (r *renameAcrossDifferentImplementationsError) afterLoad() {}

func (r *renameAcrossDifferentImplementationsError) StateLoad(stateSourceObject state.Source) {
}

func (i *InodeSymlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeSymlink"
}

func (i *InodeSymlink) StateFields() []string {
	return []string{
		"InodeNotDirectory",
	}
}

func (i *InodeSymlink) beforeSave() {}

func (i *InodeSymlink) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeNotDirectory)
}

func (i *InodeSymlink) afterLoad() {}

func (i *InodeSymlink) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeNotDirectory)
}

func (s *StaticDirectory) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.StaticDirectory"
}

func (s *StaticDirectory) StateFields() []string {
	return []string{
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeDirectoryNoNewChildren",
		"InodeNoStatFS",
		"InodeNotSymlink",
		"InodeTemporary",
		"OrderedChildren",
		"StaticDirectoryRefs",
		"locks",
		"fdOpts",
	}
}

func (s *StaticDirectory) beforeSave() {}

func (s *StaticDirectory) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.InodeAlwaysValid)
	stateSinkObject.Save(1, &s.InodeAttrs)
	stateSinkObject.Save(2, &s.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(3, &s.InodeNoStatFS)
	stateSinkObject.Save(4, &s.InodeNotSymlink)
	stateSinkObject.Save(5, &s.InodeTemporary)
	stateSinkObject.Save(6, &s.OrderedChildren)
	stateSinkObject.Save(7, &s.StaticDirectoryRefs)
	stateSinkObject.Save(8, &s.locks)
	stateSinkObject.Save(9, &s.fdOpts)
}

func (s *StaticDirectory) afterLoad() {}

func (s *StaticDirectory) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.InodeAlwaysValid)
	stateSourceObject.Load(1, &s.InodeAttrs)
	stateSourceObject.Load(2, &s.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(3, &s.InodeNoStatFS)
	stateSourceObject.Load(4, &s.InodeNotSymlink)
	stateSourceObject.Load(5, &s.InodeTemporary)
	stateSourceObject.Load(6, &s.OrderedChildren)
	stateSourceObject.Load(7, &s.StaticDirectoryRefs)
	stateSourceObject.Load(8, &s.locks)
	stateSourceObject.Load(9, &s.fdOpts)
}

func (i *InodeAlwaysValid) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeAlwaysValid"
}

func (i *InodeAlwaysValid) StateFields() []string {
	return []string{}
}

func (i *InodeAlwaysValid) beforeSave() {}

func (i *InodeAlwaysValid) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *InodeAlwaysValid) afterLoad() {}

func (i *InodeAlwaysValid) StateLoad(stateSourceObject state.Source) {
}

func (i *InodeTemporary) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeTemporary"
}

func (i *InodeTemporary) StateFields() []string {
	return []string{}
}

func (i *InodeTemporary) beforeSave() {}

func (i *InodeTemporary) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *InodeTemporary) afterLoad() {}

func (i *InodeTemporary) StateLoad(stateSourceObject state.Source) {
}

func (i *InodeNoStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.InodeNoStatFS"
}

func (i *InodeNoStatFS) StateFields() []string {
	return []string{}
}

func (i *InodeNoStatFS) beforeSave() {}

func (i *InodeNoStatFS) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *InodeNoStatFS) afterLoad() {}

func (i *InodeNoStatFS) StateLoad(stateSourceObject state.Source) {
}

func (fs *Filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.Filesystem"
}

func (fs *Filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"droppedDentries",
		"nextInoMinusOne",
		"cachedDentries",
		"cachedDentriesLen",
		"MaxCachedDentries",
	}
}

func (fs *Filesystem) beforeSave() {}

func (fs *Filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.droppedDentries)
	stateSinkObject.Save(2, &fs.nextInoMinusOne)
	stateSinkObject.Save(3, &fs.cachedDentries)
	stateSinkObject.Save(4, &fs.cachedDentriesLen)
	stateSinkObject.Save(5, &fs.MaxCachedDentries)
}

func (fs *Filesystem) afterLoad() {}

func (fs *Filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.droppedDentries)
	stateSourceObject.Load(2, &fs.nextInoMinusOne)
	stateSourceObject.Load(3, &fs.cachedDentries)
	stateSourceObject.Load(4, &fs.cachedDentriesLen)
	stateSourceObject.Load(5, &fs.MaxCachedDentries)
}

func (d *Dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.Dentry"
}

func (d *Dentry) StateFields() []string {
	return []string{
		"vfsd",
		"refs",
		"fs",
		"flags",
		"parent",
		"name",
		"cached",
		"dentryEntry",
		"children",
		"inode",
	}
}

func (d *Dentry) beforeSave() {}

func (d *Dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.refs)
	stateSinkObject.Save(2, &d.fs)
	stateSinkObject.Save(3, &d.flags)
	stateSinkObject.Save(4, &d.parent)
	stateSinkObject.Save(5, &d.name)
	stateSinkObject.Save(6, &d.cached)
	stateSinkObject.Save(7, &d.dentryEntry)
	stateSinkObject.Save(8, &d.children)
	stateSinkObject.Save(9, &d.inode)
}

func (d *Dentry) afterLoad() {}

func (d *Dentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.refs)
	stateSourceObject.Load(2, &d.fs)
	stateSourceObject.Load(3, &d.flags)
	stateSourceObject.Load(4, &d.parent)
	stateSourceObject.Load(5, &d.name)
	stateSourceObject.Load(6, &d.cached)
	stateSourceObject.Load(7, &d.dentryEntry)
	stateSourceObject.Load(8, &d.children)
	stateSourceObject.Load(9, &d.inode)
}

func (i *inodePlatformFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.inodePlatformFile"
}

func (i *inodePlatformFile) StateFields() []string {
	return []string{
		"hostFD",
		"fdRefs",
		"fileMapper",
	}
}

func (i *inodePlatformFile) beforeSave() {}

func (i *inodePlatformFile) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.hostFD)
	stateSinkObject.Save(1, &i.fdRefs)
	stateSinkObject.Save(2, &i.fileMapper)
}

func (i *inodePlatformFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.hostFD)
	stateSourceObject.Load(1, &i.fdRefs)
	stateSourceObject.Load(2, &i.fileMapper)
	stateSourceObject.AfterLoad(i.afterLoad)
}

func (i *CachedMappable) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.CachedMappable"
}

func (i *CachedMappable) StateFields() []string {
	return []string{
		"mappings",
		"pf",
	}
}

func (i *CachedMappable) beforeSave() {}

func (i *CachedMappable) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.mappings)
	stateSinkObject.Save(1, &i.pf)
}

func (i *CachedMappable) afterLoad() {}

func (i *CachedMappable) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.mappings)
	stateSourceObject.Load(1, &i.pf)
}

func (l *slotList) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.slotList"
}

func (l *slotList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *slotList) beforeSave() {}

func (l *slotList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *slotList) afterLoad() {}

func (l *slotList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *slotEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.slotEntry"
}

func (e *slotEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *slotEntry) beforeSave() {}

func (e *slotEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *slotEntry) afterLoad() {}

func (e *slotEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *StaticDirectoryRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.StaticDirectoryRefs"
}

func (r *StaticDirectoryRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *StaticDirectoryRefs) beforeSave() {}

func (r *StaticDirectoryRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

func (r *StaticDirectoryRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (s *StaticSymlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.StaticSymlink"
}

func (s *StaticSymlink) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeSymlink",
		"InodeNoStatFS",
		"target",
	}
}

func (s *StaticSymlink) beforeSave() {}

func (s *StaticSymlink) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.InodeAttrs)
	stateSinkObject.Save(1, &s.InodeNoopRefCount)
	stateSinkObject.Save(2, &s.InodeSymlink)
	stateSinkObject.Save(3, &s.InodeNoStatFS)
	stateSinkObject.Save(4, &s.target)
}

func (s *StaticSymlink) afterLoad() {}

func (s *StaticSymlink) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.InodeAttrs)
	stateSourceObject.Load(1, &s.InodeNoopRefCount)
	stateSourceObject.Load(2, &s.InodeSymlink)
	stateSourceObject.Load(3, &s.InodeNoStatFS)
	stateSourceObject.Load(4, &s.target)
}

func (dir *syntheticDirectory) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.syntheticDirectory"
}

func (dir *syntheticDirectory) StateFields() []string {
	return []string{
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeNoStatFS",
		"InodeNotSymlink",
		"OrderedChildren",
		"syntheticDirectoryRefs",
		"locks",
	}
}

func (dir *syntheticDirectory) beforeSave() {}

func (dir *syntheticDirectory) StateSave(stateSinkObject state.Sink) {
	dir.beforeSave()
	stateSinkObject.Save(0, &dir.InodeAlwaysValid)
	stateSinkObject.Save(1, &dir.InodeAttrs)
	stateSinkObject.Save(2, &dir.InodeNoStatFS)
	stateSinkObject.Save(3, &dir.InodeNotSymlink)
	stateSinkObject.Save(4, &dir.OrderedChildren)
	stateSinkObject.Save(5, &dir.syntheticDirectoryRefs)
	stateSinkObject.Save(6, &dir.locks)
}

func (dir *syntheticDirectory) afterLoad() {}

func (dir *syntheticDirectory) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &dir.InodeAlwaysValid)
	stateSourceObject.Load(1, &dir.InodeAttrs)
	stateSourceObject.Load(2, &dir.InodeNoStatFS)
	stateSourceObject.Load(3, &dir.InodeNotSymlink)
	stateSourceObject.Load(4, &dir.OrderedChildren)
	stateSourceObject.Load(5, &dir.syntheticDirectoryRefs)
	stateSourceObject.Load(6, &dir.locks)
}

func (r *syntheticDirectoryRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/kernfs.syntheticDirectoryRefs"
}

func (r *syntheticDirectoryRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *syntheticDirectoryRefs) beforeSave() {}

func (r *syntheticDirectoryRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

func (r *syntheticDirectoryRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*dentryList)(nil))
	state.Register((*dentryEntry)(nil))
	state.Register((*DynamicBytesFile)(nil))
	state.Register((*DynamicBytesFD)(nil))
	state.Register((*SeekEndConfig)(nil))
	state.Register((*GenericDirectoryFDOptions)(nil))
	state.Register((*GenericDirectoryFD)(nil))
	state.Register((*InodeNoopRefCount)(nil))
	state.Register((*InodeDirectoryNoNewChildren)(nil))
	state.Register((*InodeNotDirectory)(nil))
	state.Register((*InodeNotSymlink)(nil))
	state.Register((*InodeAttrs)(nil))
	state.Register((*slot)(nil))
	state.Register((*OrderedChildrenOptions)(nil))
	state.Register((*OrderedChildren)(nil))
	state.Register((*renameAcrossDifferentImplementationsError)(nil))
	state.Register((*InodeSymlink)(nil))
	state.Register((*StaticDirectory)(nil))
	state.Register((*InodeAlwaysValid)(nil))
	state.Register((*InodeTemporary)(nil))
	state.Register((*InodeNoStatFS)(nil))
	state.Register((*Filesystem)(nil))
	state.Register((*Dentry)(nil))
	state.Register((*inodePlatformFile)(nil))
	state.Register((*CachedMappable)(nil))
	state.Register((*slotList)(nil))
	state.Register((*slotEntry)(nil))
	state.Register((*StaticDirectoryRefs)(nil))
	state.Register((*StaticSymlink)(nil))
	state.Register((*syntheticDirectory)(nil))
	state.Register((*syntheticDirectoryRefs)(nil))
}
