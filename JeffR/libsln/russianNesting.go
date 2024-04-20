package libsln

import (
	"JeffR/lib"
	"cmp"
	"slices"
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

func envCmp(a Envelope, b Envelope) int {
	diff := cmp.Compare(envWidth(a), envWidth(b))
	if diff == 0 {
		diff = cmp.Compare(envHeight(a), envHeight(b))
	}
	return diff
}

/*
func envSortInPlace(envelopes *Envelopes) {
	slices.SortFunc(*envelopes, envCmp)
}
*/

func EnvSort(envelopes Envelopes) Envelopes {
	slices.SortFunc(envelopes, envCmp)
	return envelopes
}

func EnvFilter(envelopes Envelopes) Envelopes {
	assert(envelopes != nil && len(envelopes) > 0, "empty []Envelope collection, nothing to filter")

	filteredEnvelopes := make(Envelopes, 0)
	var lastEnv = Envelope{-1, -1}
	// var nextEnv = Envelope{DIM_MAX + 1, DIM_MAX + 1}
	for i, env := range envelopes {
		if i > 0 {
			if envWidth(env) > envWidth(lastEnv) && envHeight(env) > envHeight(lastEnv) {
				// last envelope would fit inside the current one
				filteredEnvelopes = append(filteredEnvelopes, lastEnv)
				lastEnv = env
			}
		} else {
			lastEnv = env
		}
		if i == len(envelopes)-1 {
			filteredEnvelopes = append(filteredEnvelopes, lastEnv)
		}

	}

	return filteredEnvelopes

}

//////////////////////////////////////////////////////////////////
