// automatically generated by stateify.

package mm

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (a *aioManager) StateTypeName() string {
	return "pkg/sentry/mm.aioManager"
}

func (a *aioManager) StateFields() []string {
	return []string{
		"contexts",
	}
}

func (a *aioManager) beforeSave() {}

func (a *aioManager) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.contexts)
}

func (a *aioManager) afterLoad() {}

func (a *aioManager) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.contexts)
}

func (i *ioResult) StateTypeName() string {
	return "pkg/sentry/mm.ioResult"
}

func (i *ioResult) StateFields() []string {
	return []string{
		"data",
		"ioEntry",
	}
}

func (i *ioResult) beforeSave() {}

func (i *ioResult) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.data)
	stateSinkObject.Save(1, &i.ioEntry)
}

func (i *ioResult) afterLoad() {}

func (i *ioResult) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.data)
	stateSourceObject.Load(1, &i.ioEntry)
}

func (a *AIOContext) StateTypeName() string {
	return "pkg/sentry/mm.AIOContext"
}

func (a *AIOContext) StateFields() []string {
	return []string{
		"results",
		"maxOutstanding",
		"outstanding",
	}
}

func (a *AIOContext) beforeSave() {}

func (a *AIOContext) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	if !state.IsZeroValue(&a.dead) {
		state.Failf("dead is %#v, expected zero", &a.dead)
	}
	stateSinkObject.Save(0, &a.results)
	stateSinkObject.Save(1, &a.maxOutstanding)
	stateSinkObject.Save(2, &a.outstanding)
}

func (a *AIOContext) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.results)
	stateSourceObject.Load(1, &a.maxOutstanding)
	stateSourceObject.Load(2, &a.outstanding)
	stateSourceObject.AfterLoad(a.afterLoad)
}

func (a *aioMappable) StateTypeName() string {
	return "pkg/sentry/mm.aioMappable"
}

func (a *aioMappable) StateFields() []string {
	return []string{
		"aioMappableRefs",
		"mfp",
		"fr",
	}
}

func (a *aioMappable) beforeSave() {}

func (a *aioMappable) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.aioMappableRefs)
	stateSinkObject.Save(1, &a.mfp)
	stateSinkObject.Save(2, &a.fr)
}

func (a *aioMappable) afterLoad() {}

func (a *aioMappable) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.aioMappableRefs)
	stateSourceObject.Load(1, &a.mfp)
	stateSourceObject.Load(2, &a.fr)
}

func (a *aioMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.aioMappableRefs"
}

func (a *aioMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (a *aioMappableRefs) beforeSave() {}

func (a *aioMappableRefs) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.refCount)
}

func (a *aioMappableRefs) afterLoad() {}

func (a *aioMappableRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.refCount)
}

func (f *fileRefcountSet) StateTypeName() string {
	return "pkg/sentry/mm.fileRefcountSet"
}

func (f *fileRefcountSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (f *fileRefcountSet) beforeSave() {}

func (f *fileRefcountSet) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	var rootValue *fileRefcountSegmentDataSlices = f.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (f *fileRefcountSet) afterLoad() {}

func (f *fileRefcountSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*fileRefcountSegmentDataSlices), func(y interface{}) { f.loadRoot(y.(*fileRefcountSegmentDataSlices)) })
}

func (f *fileRefcountnode) StateTypeName() string {
	return "pkg/sentry/mm.fileRefcountnode"
}

func (f *fileRefcountnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (f *fileRefcountnode) beforeSave() {}

func (f *fileRefcountnode) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.nrSegments)
	stateSinkObject.Save(1, &f.parent)
	stateSinkObject.Save(2, &f.parentIndex)
	stateSinkObject.Save(3, &f.hasChildren)
	stateSinkObject.Save(4, &f.maxGap)
	stateSinkObject.Save(5, &f.keys)
	stateSinkObject.Save(6, &f.values)
	stateSinkObject.Save(7, &f.children)
}

func (f *fileRefcountnode) afterLoad() {}

func (f *fileRefcountnode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.nrSegments)
	stateSourceObject.Load(1, &f.parent)
	stateSourceObject.Load(2, &f.parentIndex)
	stateSourceObject.Load(3, &f.hasChildren)
	stateSourceObject.Load(4, &f.maxGap)
	stateSourceObject.Load(5, &f.keys)
	stateSourceObject.Load(6, &f.values)
	stateSourceObject.Load(7, &f.children)
}

func (f *fileRefcountSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/mm.fileRefcountSegmentDataSlices"
}

func (f *fileRefcountSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (f *fileRefcountSegmentDataSlices) beforeSave() {}

func (f *fileRefcountSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Start)
	stateSinkObject.Save(1, &f.End)
	stateSinkObject.Save(2, &f.Values)
}

func (f *fileRefcountSegmentDataSlices) afterLoad() {}

func (f *fileRefcountSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Start)
	stateSourceObject.Load(1, &f.End)
	stateSourceObject.Load(2, &f.Values)
}

func (i *ioList) StateTypeName() string {
	return "pkg/sentry/mm.ioList"
}

func (i *ioList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (i *ioList) beforeSave() {}

func (i *ioList) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.head)
	stateSinkObject.Save(1, &i.tail)
}

func (i *ioList) afterLoad() {}

func (i *ioList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.head)
	stateSourceObject.Load(1, &i.tail)
}

func (i *ioEntry) StateTypeName() string {
	return "pkg/sentry/mm.ioEntry"
}

func (i *ioEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (i *ioEntry) beforeSave() {}

func (i *ioEntry) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.next)
	stateSinkObject.Save(1, &i.prev)
}

func (i *ioEntry) afterLoad() {}

func (i *ioEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.next)
	stateSourceObject.Load(1, &i.prev)
}

func (m *MemoryManager) StateTypeName() string {
	return "pkg/sentry/mm.MemoryManager"
}

func (m *MemoryManager) StateFields() []string {
	return []string{
		"p",
		"mfp",
		"layout",
		"privateRefs",
		"users",
		"vmas",
		"brk",
		"usageAS",
		"lockedAS",
		"dataAS",
		"defMLockMode",
		"pmas",
		"curRSS",
		"maxRSS",
		"argv",
		"envv",
		"auxv",
		"executable",
		"dumpability",
		"aioManager",
		"sleepForActivation",
		"vdsoSigReturnAddr",
		"membarrierPrivateEnabled",
		"membarrierRSeqEnabled",
	}
}

func (m *MemoryManager) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	if !state.IsZeroValue(&m.active) {
		state.Failf("active is %#v, expected zero", &m.active)
	}
	if !state.IsZeroValue(&m.captureInvalidations) {
		state.Failf("captureInvalidations is %#v, expected zero", &m.captureInvalidations)
	}
	stateSinkObject.Save(0, &m.p)
	stateSinkObject.Save(1, &m.mfp)
	stateSinkObject.Save(2, &m.layout)
	stateSinkObject.Save(3, &m.privateRefs)
	stateSinkObject.Save(4, &m.users)
	stateSinkObject.Save(5, &m.vmas)
	stateSinkObject.Save(6, &m.brk)
	stateSinkObject.Save(7, &m.usageAS)
	stateSinkObject.Save(8, &m.lockedAS)
	stateSinkObject.Save(9, &m.dataAS)
	stateSinkObject.Save(10, &m.defMLockMode)
	stateSinkObject.Save(11, &m.pmas)
	stateSinkObject.Save(12, &m.curRSS)
	stateSinkObject.Save(13, &m.maxRSS)
	stateSinkObject.Save(14, &m.argv)
	stateSinkObject.Save(15, &m.envv)
	stateSinkObject.Save(16, &m.auxv)
	stateSinkObject.Save(17, &m.executable)
	stateSinkObject.Save(18, &m.dumpability)
	stateSinkObject.Save(19, &m.aioManager)
	stateSinkObject.Save(20, &m.sleepForActivation)
	stateSinkObject.Save(21, &m.vdsoSigReturnAddr)
	stateSinkObject.Save(22, &m.membarrierPrivateEnabled)
	stateSinkObject.Save(23, &m.membarrierRSeqEnabled)
}

func (m *MemoryManager) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.p)
	stateSourceObject.Load(1, &m.mfp)
	stateSourceObject.Load(2, &m.layout)
	stateSourceObject.Load(3, &m.privateRefs)
	stateSourceObject.Load(4, &m.users)
	stateSourceObject.Load(5, &m.vmas)
	stateSourceObject.Load(6, &m.brk)
	stateSourceObject.Load(7, &m.usageAS)
	stateSourceObject.Load(8, &m.lockedAS)
	stateSourceObject.Load(9, &m.dataAS)
	stateSourceObject.Load(10, &m.defMLockMode)
	stateSourceObject.Load(11, &m.pmas)
	stateSourceObject.Load(12, &m.curRSS)
	stateSourceObject.Load(13, &m.maxRSS)
	stateSourceObject.Load(14, &m.argv)
	stateSourceObject.Load(15, &m.envv)
	stateSourceObject.Load(16, &m.auxv)
	stateSourceObject.Load(17, &m.executable)
	stateSourceObject.Load(18, &m.dumpability)
	stateSourceObject.Load(19, &m.aioManager)
	stateSourceObject.Load(20, &m.sleepForActivation)
	stateSourceObject.Load(21, &m.vdsoSigReturnAddr)
	stateSourceObject.Load(22, &m.membarrierPrivateEnabled)
	stateSourceObject.Load(23, &m.membarrierRSeqEnabled)
	stateSourceObject.AfterLoad(m.afterLoad)
}

func (v *vma) StateTypeName() string {
	return "pkg/sentry/mm.vma"
}

func (v *vma) StateFields() []string {
	return []string{
		"mappable",
		"off",
		"realPerms",
		"dontfork",
		"mlockMode",
		"numaPolicy",
		"numaNodemask",
		"id",
		"hint",
	}
}

func (v *vma) beforeSave() {}

func (v *vma) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	var realPermsValue int = v.saveRealPerms()
	stateSinkObject.SaveValue(2, realPermsValue)
	stateSinkObject.Save(0, &v.mappable)
	stateSinkObject.Save(1, &v.off)
	stateSinkObject.Save(3, &v.dontfork)
	stateSinkObject.Save(4, &v.mlockMode)
	stateSinkObject.Save(5, &v.numaPolicy)
	stateSinkObject.Save(6, &v.numaNodemask)
	stateSinkObject.Save(7, &v.id)
	stateSinkObject.Save(8, &v.hint)
}

func (v *vma) afterLoad() {}

func (v *vma) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.mappable)
	stateSourceObject.Load(1, &v.off)
	stateSourceObject.Load(3, &v.dontfork)
	stateSourceObject.Load(4, &v.mlockMode)
	stateSourceObject.Load(5, &v.numaPolicy)
	stateSourceObject.Load(6, &v.numaNodemask)
	stateSourceObject.Load(7, &v.id)
	stateSourceObject.Load(8, &v.hint)
	stateSourceObject.LoadValue(2, new(int), func(y interface{}) { v.loadRealPerms(y.(int)) })
}

func (p *pma) StateTypeName() string {
	return "pkg/sentry/mm.pma"
}

func (p *pma) StateFields() []string {
	return []string{
		"off",
		"translatePerms",
		"effectivePerms",
		"maxPerms",
		"needCOW",
		"private",
	}
}

func (p *pma) beforeSave() {}

func (p *pma) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.off)
	stateSinkObject.Save(1, &p.translatePerms)
	stateSinkObject.Save(2, &p.effectivePerms)
	stateSinkObject.Save(3, &p.maxPerms)
	stateSinkObject.Save(4, &p.needCOW)
	stateSinkObject.Save(5, &p.private)
}

func (p *pma) afterLoad() {}

func (p *pma) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.off)
	stateSourceObject.Load(1, &p.translatePerms)
	stateSourceObject.Load(2, &p.effectivePerms)
	stateSourceObject.Load(3, &p.maxPerms)
	stateSourceObject.Load(4, &p.needCOW)
	stateSourceObject.Load(5, &p.private)
}

func (p *privateRefs) StateTypeName() string {
	return "pkg/sentry/mm.privateRefs"
}

func (p *privateRefs) StateFields() []string {
	return []string{
		"refs",
	}
}

func (p *privateRefs) beforeSave() {}

func (p *privateRefs) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.refs)
}

func (p *privateRefs) afterLoad() {}

func (p *privateRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.refs)
}

func (p *pmaSet) StateTypeName() string {
	return "pkg/sentry/mm.pmaSet"
}

func (p *pmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (p *pmaSet) beforeSave() {}

func (p *pmaSet) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var rootValue *pmaSegmentDataSlices = p.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (p *pmaSet) afterLoad() {}

func (p *pmaSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*pmaSegmentDataSlices), func(y interface{}) { p.loadRoot(y.(*pmaSegmentDataSlices)) })
}

func (p *pmanode) StateTypeName() string {
	return "pkg/sentry/mm.pmanode"
}

func (p *pmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (p *pmanode) beforeSave() {}

func (p *pmanode) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.nrSegments)
	stateSinkObject.Save(1, &p.parent)
	stateSinkObject.Save(2, &p.parentIndex)
	stateSinkObject.Save(3, &p.hasChildren)
	stateSinkObject.Save(4, &p.maxGap)
	stateSinkObject.Save(5, &p.keys)
	stateSinkObject.Save(6, &p.values)
	stateSinkObject.Save(7, &p.children)
}

func (p *pmanode) afterLoad() {}

func (p *pmanode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.nrSegments)
	stateSourceObject.Load(1, &p.parent)
	stateSourceObject.Load(2, &p.parentIndex)
	stateSourceObject.Load(3, &p.hasChildren)
	stateSourceObject.Load(4, &p.maxGap)
	stateSourceObject.Load(5, &p.keys)
	stateSourceObject.Load(6, &p.values)
	stateSourceObject.Load(7, &p.children)
}

func (p *pmaSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/mm.pmaSegmentDataSlices"
}

func (p *pmaSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (p *pmaSegmentDataSlices) beforeSave() {}

func (p *pmaSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Start)
	stateSinkObject.Save(1, &p.End)
	stateSinkObject.Save(2, &p.Values)
}

func (p *pmaSegmentDataSlices) afterLoad() {}

func (p *pmaSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Start)
	stateSourceObject.Load(1, &p.End)
	stateSourceObject.Load(2, &p.Values)
}

func (s *SpecialMappable) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappable"
}

func (s *SpecialMappable) StateFields() []string {
	return []string{
		"SpecialMappableRefs",
		"mfp",
		"fr",
		"name",
	}
}

func (s *SpecialMappable) beforeSave() {}

func (s *SpecialMappable) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.SpecialMappableRefs)
	stateSinkObject.Save(1, &s.mfp)
	stateSinkObject.Save(2, &s.fr)
	stateSinkObject.Save(3, &s.name)
}

func (s *SpecialMappable) afterLoad() {}

func (s *SpecialMappable) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.SpecialMappableRefs)
	stateSourceObject.Load(1, &s.mfp)
	stateSourceObject.Load(2, &s.fr)
	stateSourceObject.Load(3, &s.name)
}

func (s *SpecialMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappableRefs"
}

func (s *SpecialMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (s *SpecialMappableRefs) beforeSave() {}

func (s *SpecialMappableRefs) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.refCount)
}

func (s *SpecialMappableRefs) afterLoad() {}

func (s *SpecialMappableRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.refCount)
}

func (v *vmaSet) StateTypeName() string {
	return "pkg/sentry/mm.vmaSet"
}

func (v *vmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (v *vmaSet) beforeSave() {}

func (v *vmaSet) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	var rootValue *vmaSegmentDataSlices = v.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (v *vmaSet) afterLoad() {}

func (v *vmaSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*vmaSegmentDataSlices), func(y interface{}) { v.loadRoot(y.(*vmaSegmentDataSlices)) })
}

func (v *vmanode) StateTypeName() string {
	return "pkg/sentry/mm.vmanode"
}

func (v *vmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (v *vmanode) beforeSave() {}

func (v *vmanode) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.nrSegments)
	stateSinkObject.Save(1, &v.parent)
	stateSinkObject.Save(2, &v.parentIndex)
	stateSinkObject.Save(3, &v.hasChildren)
	stateSinkObject.Save(4, &v.maxGap)
	stateSinkObject.Save(5, &v.keys)
	stateSinkObject.Save(6, &v.values)
	stateSinkObject.Save(7, &v.children)
}

func (v *vmanode) afterLoad() {}

func (v *vmanode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.nrSegments)
	stateSourceObject.Load(1, &v.parent)
	stateSourceObject.Load(2, &v.parentIndex)
	stateSourceObject.Load(3, &v.hasChildren)
	stateSourceObject.Load(4, &v.maxGap)
	stateSourceObject.Load(5, &v.keys)
	stateSourceObject.Load(6, &v.values)
	stateSourceObject.Load(7, &v.children)
}

func (v *vmaSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/mm.vmaSegmentDataSlices"
}

func (v *vmaSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (v *vmaSegmentDataSlices) beforeSave() {}

func (v *vmaSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.Start)
	stateSinkObject.Save(1, &v.End)
	stateSinkObject.Save(2, &v.Values)
}

func (v *vmaSegmentDataSlices) afterLoad() {}

func (v *vmaSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.Start)
	stateSourceObject.Load(1, &v.End)
	stateSourceObject.Load(2, &v.Values)
}

func init() {
	state.Register((*aioManager)(nil))
	state.Register((*ioResult)(nil))
	state.Register((*AIOContext)(nil))
	state.Register((*aioMappable)(nil))
	state.Register((*aioMappableRefs)(nil))
	state.Register((*fileRefcountSet)(nil))
	state.Register((*fileRefcountnode)(nil))
	state.Register((*fileRefcountSegmentDataSlices)(nil))
	state.Register((*ioList)(nil))
	state.Register((*ioEntry)(nil))
	state.Register((*MemoryManager)(nil))
	state.Register((*vma)(nil))
	state.Register((*pma)(nil))
	state.Register((*privateRefs)(nil))
	state.Register((*pmaSet)(nil))
	state.Register((*pmanode)(nil))
	state.Register((*pmaSegmentDataSlices)(nil))
	state.Register((*SpecialMappable)(nil))
	state.Register((*SpecialMappableRefs)(nil))
	state.Register((*vmaSet)(nil))
	state.Register((*vmanode)(nil))
	state.Register((*vmaSegmentDataSlices)(nil))
}
