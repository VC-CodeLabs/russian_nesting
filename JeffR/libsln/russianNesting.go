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
