package librn

import (
	"cmp"
	"slices"
)

func envCmp(a Envelope, b Envelope) int {
	diff := cmp.Compare(envWidth(a), envWidth(b))
	if diff == 0 {
		diff = cmp.Compare(envHeight(a), envHeight(b))
	}
	return diff
}

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
