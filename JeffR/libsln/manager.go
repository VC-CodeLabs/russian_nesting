package librn

type IRussianNestingManager struct {
	russianNestingImpl EnvArrayPreAlloc
}

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
