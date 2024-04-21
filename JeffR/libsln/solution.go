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

	var maxNestedEnvelopes Envelopes

	// var nextEnv = Envelope{DIM_MAX + 1, DIM_MAX + 1}
	countOfEnvelopes := len(envelopes)

	for i := range envelopes {

		if i > 0 && len(maxNestedEnvelopes) > countOfEnvelopes-i {
			// can't possibly find a bigger one, don't bother
			break
		}

		filteredEnvelopes := make(Envelopes, 0)
		var lastEnv = Envelope{-1, -1}
		for j := i; j < countOfEnvelopes; j++ {
			env := envelopes[j]
			if j > i {
				if envWidth(env) > envWidth(lastEnv) && envHeight(env) > envHeight(lastEnv) {
					// last envelope would fit inside the current one
					filteredEnvelopes = append(filteredEnvelopes, lastEnv)
					lastEnv = env
				}
			} else {
				lastEnv = env
			}
			if j == countOfEnvelopes-1 {
				filteredEnvelopes = append(filteredEnvelopes, lastEnv)
			}
		}

		if i == 0 || len(filteredEnvelopes) > len(maxNestedEnvelopes) {
			maxNestedEnvelopes = filteredEnvelopes
		}
	}

	return maxNestedEnvelopes

}

//////////////////////////////////////////////////////////////////
