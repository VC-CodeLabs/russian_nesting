package librn

type RussianNesting interface {
	InitData()
	PutDataItem(int, int)
	CloseData()
	GetNestedEnvelopes() Envelopes
	GetNestedCount() int
}

type EnvArrayWithAppend struct {
	envelopes Envelopes
}

func (x EnvArrayWithAppend) InitData() {
}

func (x *EnvArrayWithAppend) PutDataItem(w int, h int) {
	(*x).envelopes = append(x.envelopes, Envelope{w, h})
}

func (x EnvArrayWithAppend) CloseData() {

}

func (x EnvArrayWithAppend) GetNestedEnvelopes() Envelopes {
	return EnvFilter(EnvSort(x.envelopes))
}

func (x EnvArrayWithAppend) GetNestedCount() int {
	return len(x.GetNestedEnvelopes())
}

type EnvArrayPreAlloc struct {
	base      EnvArrayWithAppend
	itemCount int
}

func (x *EnvArrayPreAlloc) InitData() {
	(*x).base.envelopes = make(Envelopes, ENV_MAX)
	// (*x).itemCount = 0
}

func (x *EnvArrayPreAlloc) PutDataItem(w int, h int) {
	(*x).base.envelopes[x.itemCount] = Envelope{w, h}
	(*x).itemCount++
}

func (x *EnvArrayPreAlloc) CloseData() {
	(*x).base.envelopes = x.base.envelopes[:x.itemCount]
}

func (x EnvArrayPreAlloc) GetNestedEnvelopes() Envelopes {
	return x.base.GetNestedEnvelopes()
}

func (x EnvArrayPreAlloc) GetNestedCount() int {
	return x.base.GetNestedCount()
}

//////////////////////////////////////////////////////////////////
