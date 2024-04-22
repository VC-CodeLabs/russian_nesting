package librn

import . "JeffR/libsln"

type EnvelopesS []EnvStruct

type EnvArray [2]int

// indexes into EnvArray for width & height
const (
	WIDTH  = 0
	HEIGHT = 1
)

type EnvelopesA [][2]int
