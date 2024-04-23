package librn

// defines the high-level interface used in the main code;
// abstracting at this level means I can easily swap in
// a different impl for testing, while theoretically
// having little or no significant impact on performance
type RussianNestingManager struct {
	russianNestingAlgoImpl EnvArrayPreAlloc
	// vs e.g. EnvArrayWithAppend
}

//
// implement the IRussianNestingAlgo interface for our manager;
// note this just delegates to the underlying impl,
// which must implement the same interface!
//

func (rnMgr *RussianNestingManager) InitData() {
	rnMgr.russianNestingAlgoImpl.InitData()
}

func (rnMgr *RussianNestingManager) PutDataItem(w int, h int) {
	rnMgr.russianNestingAlgoImpl.PutDataItem(w, h)
}

func (rnMgr *RussianNestingManager) CloseData() {
	rnMgr.russianNestingAlgoImpl.CloseData()
}

func (rnMgr RussianNestingManager) GetMaxNestedEnvelopes() Envelopes {
	return rnMgr.russianNestingAlgoImpl.GetMaxNestedEnvelopes()
}

func (rnMgr RussianNestingManager) GetMaxNestedCount() int {
	return rnMgr.russianNestingAlgoImpl.GetMaxNestedCount()
}
