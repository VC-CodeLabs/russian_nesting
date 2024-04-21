package librn

import (
	"JeffR/lib"
)

const DIM_MIN = 1

// var DIM_MAX = int(math.Pow(10, 5))
// var DIM_MAX = int(math.Pow10(5))
const DIM_MAX = int(1e5)

const ENV_MIN = 1

// var ENV_MAX = DIM_MAX
const ENV_MAX = DIM_MAX

type EnvStruct struct {
	Width  int
	Height int
}

type EnvelopesS []EnvStruct

type EnvArray [2]int

// indexes into EnvArray for width & height
const (
	WIDTH  = 0
	HEIGHT = 1
)

type EnvelopesA [][2]int

type Envelope EnvStruct
type Envelopes []Envelope

func EnvWidth(env Envelope) int {
	return env.Width
}

func EnvHeight(env Envelope) int {
	return env.Height
}

// local proxies to minimize diffs during isolation work

func envWidth(env Envelope) int {
	return EnvWidth(env)
}

func envHeight(env Envelope) int {
	return EnvHeight(env)
}

//////////////////////////////////////////////////////////////////
// local proxies for fx moved to lib
//

func assert(cond bool, msg string) {
	lib.Assert(cond, msg)
}

//////////////////////////////////////////////////////////////////
// tests
//

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
