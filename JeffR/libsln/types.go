package librn

// defines the basic 2d envelope structure;
// testing indicates this is generally faster than [2]int
type EnvStruct struct {
	Width  int
	Height int
}

//
// a little abstraction helped during development when considering different impls
//

// the fundamental envelope type
type Envelope EnvStruct

// a collection of envelope objects
type Envelopes []Envelope

// helpers support abstraction of envelope type

// gets the width of the specified envelope
func EnvWidth(env Envelope) int {
	return env.Width
}

// gets the height of the specified envelope
func EnvHeight(env Envelope) int {
	return env.Height
}

//////////////////////////////////////////////////////////////////
