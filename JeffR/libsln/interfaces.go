package librn

//
// defines an interface for the solution algorithm,
// allowing for different impls to be swapped in and tested
//
// NOTE GetMaxNestedCount() will generally be impl'd as
// len(GetMaxNestedEnvelopes()), so if you want both,
// get the length yourself else you're *probably*
// finding the max nested collection twice!
//
type IRussianNestingAlgo interface {
	// prepare data storage
	InitData()
	// add a new envelope to data storage
	PutDataItem(int, int)
	// finalize data storage
	CloseData()
	// get the max nested envelopes collection
	GetMaxNestedEnvelopes() Envelopes
	// get the count of max nested envelopes
	GetMaxNestedCount() int
}

// define an implementation of the solution
// that uses a simple array
type EnvArrayWithAppend struct {
	envelopes Envelopes
}

// prepare the data storage
func (x *EnvArrayWithAppend) InitData() {
	// not really necessary the first time, but allows for reuse
	(*x).envelopes = make(Envelopes, 0)
}

// appends a new envelope
func (x *EnvArrayWithAppend) PutDataItem(w int, h int) {
	(*x).envelopes = append(x.envelopes, Envelope{w, h})
}

// finalize the dateset
func (x EnvArrayWithAppend) CloseData() {
	// this method is intentionally empty
}

// get the max collection by # of nested envelopes;
// if there are multiple such collections in the input,
// the first one (smallest starting envelope dims) is returned
func (x EnvArrayWithAppend) GetMaxNestedEnvelopes() Envelopes {
	return EnvFilter(EnvCompact(EnvSort(x.envelopes)))
}

// get the # of max nested envelopes;
// NOTE this is impl'd as len(GetMaxNestedEnvelopes())
// so if you want both the collection and count,
// get the count locally!!!
func (x EnvArrayWithAppend) GetMaxNestedCount() int {
	return len(x.GetMaxNestedEnvelopes())
}

// define an implementation of the solution
// that preallocates the data storage
// rather than using append every time we add a new item;
// testing indicates this can make a difference with max collections
type EnvArrayPreAlloc struct {
	base      EnvArrayWithAppend
	itemCount int
}

// preallocates data storage reserving space for ENV_MAX Envelope objects
func (x *EnvArrayPreAlloc) InitData() {
	(*x).base.envelopes = make(Envelopes, ENV_MAX)
	// allow for reuse
	(*x).itemCount = 0
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

// delegates to base
func (x EnvArrayPreAlloc) GetMaxNestedEnvelopes() Envelopes {
	return x.base.GetMaxNestedEnvelopes()
}

// delegates to base
func (x EnvArrayPreAlloc) GetMaxNestedCount() int {
	return x.base.GetMaxNestedCount()
}

//////////////////////////////////////////////////////////////////
