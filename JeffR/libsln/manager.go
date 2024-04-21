package librn

// defines the high-level interface used in the main code;
// abstracting at this level means I can easily swap in
// a different impl for testing, while theoretically
// having little or no significant impact on performance
type IRussianNestingManager struct {
	russianNestingImpl EnvArrayPreAlloc
	// vs e.g. EnvArrayWithAppend
}

//
// implement the RussianNesting interface for our manager;
// note this just delegates to the underlying impl,
// which must implement the same interface!
//

func (rnMgr *IRussianNestingManager) InitData() {
	rnMgr.russianNestingImpl.InitData()
}

func (rnMgr *IRussianNestingManager) PutDataItem(w int, h int) {
	rnMgr.russianNestingImpl.PutDataItem(w, h)
}

func (rnMgr *IRussianNestingManager) CloseData() {
	rnMgr.russianNestingImpl.CloseData()
}

func (rnMgr IRussianNestingManager) GetNestedEnvelopes() Envelopes {
	return rnMgr.russianNestingImpl.GetNestedEnvelopes()
}

func (rnMgr IRussianNestingManager) GetNestedCount() int {
	return rnMgr.russianNestingImpl.GetNestedCount()
}
