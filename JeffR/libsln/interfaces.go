package librn

//
// defines an interface for the solution,
// allowing for different impls to be swapped in and tested
//
type RussianNesting interface {
	// prepare data storage
	InitData()
	// add a new envelope to data storage
	PutDataItem(int, int)
	// finalize data storage
	CloseData()
	// get the max nested envelopes collection
	GetNestedEnvelopes() Envelopes
	// get the count of max nested envelopes
	GetNestedCount() int
}

// define an implementation of the solution
// that uses a simple array
type EnvArrayWithAppend struct {
	envelopes Envelopes
}

func (x EnvArrayWithAppend) InitData() {
}

// appends a new envelope
func (x *EnvArrayWithAppend) PutDataItem(w int, h int) {
	(*x).envelopes = append(x.envelopes, Envelope{w, h})
}

func (x EnvArrayWithAppend) CloseData() {
	// this method is intentionally empty
}

func (x EnvArrayWithAppend) GetNestedEnvelopes() Envelopes {
	return EnvFilter(EnvCompact(EnvSort(x.envelopes)))
}

func (x EnvArrayWithAppend) GetNestedCount() int {
	return len(x.GetNestedEnvelopes())
}

// define an implementation of the solution
// that preallocates the data storage
// rather than using append every time we add a new item
type EnvArrayPreAlloc struct {
	base      EnvArrayWithAppend
	itemCount int
}

// preallocates data storage reserving space for ENV_MAX Envelope objects
func (x *EnvArrayPreAlloc) InitData() {
	(*x).base.envelopes = make(Envelopes, ENV_MAX)
	// (*x).itemCount = 0
}

// adds new envelope to next available storage slot
func (x *EnvArrayPreAlloc) PutDataItem(w int, h int) {
	(*x).base.envelopes[x.itemCount] = Envelope{w, h}
	(*x).itemCount++
}

// truncates the prealloc'd storage to what was actually written
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
