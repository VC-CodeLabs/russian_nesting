package librn

type EnvStruct struct {
	Width  int
	Height int
}

type Envelope EnvStruct
type Envelopes []Envelope

func EnvWidth(env Envelope) int {
	return env.Width
}

func EnvHeight(env Envelope) int {
	return env.Height
}

//////////////////////////////////////////////////////////////////
