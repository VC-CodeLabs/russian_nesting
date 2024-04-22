package librn

import (
	"JeffR/lib"
	"cmp"
	"slices"
)

// compare two envelopes to support sorting by size
func envCmp(a Envelope, b Envelope) int {
	diff := cmp.Compare(EnvWidth(a), EnvWidth(b))
	if diff == 0 {
		diff = cmp.Compare(EnvHeight(a), EnvHeight(b))
	}
	return diff
}

// sort the collection and return it
func EnvSort(envelopes Envelopes) Envelopes {
	// TODO how expensive is this vs in-place sort?
	slices.SortFunc(envelopes, envCmp)
	return envelopes
}

// strip out the duplicates (must be sorted!!!)
//
// this significantly improves performance for
// e.g. max # envelopes all matching size
func EnvCompact(envelopes Envelopes) Envelopes {
	return slices.Compact(envelopes)
}

// get the maximum # of nested envelopes in the specified collection;
//
// the input *MUST* be sorted and ideally compacted (duplicates removed) as well
func EnvFilter(envelopes Envelopes) Envelopes {
	lib.Assert(envelopes != nil && len(envelopes) > 0, "empty []Envelope collection, nothing to filter")

	//
	// the algorithm is fairly simple,
	// once we've sorted and compacted the list;
	// using each point in the list as a starting point,
	// march thru the rest of the list,
	// finding the next (smallest) envelope
	// we can nest inside of,
	// pulling the nested envelopes into a new collection
	//
	// at the end of each run,
	// check to see if the nested collection
	// is larger than the prior max,
	// and if so, make it the new max
	//

	// track the max collection so far
	var maxNestedEnvelopes Envelopes

	// how many inputs are we dealing with
	countOfEnvelopes := len(envelopes)

	// walk thru the list of inputs, using each as
	for i := range envelopes {

		if i > 0 && len(maxNestedEnvelopes) > countOfEnvelopes-i {
			// can't possibly find a bigger one, don't bother looking
			break
		}

		// track the set of nested envelopes
		filteredEnvelopes := make(Envelopes, 0)
		// track the last envelope so we can see if it will it in the current one
		var lastEnv = Envelope{-1, -1}
		for j := i; j < countOfEnvelopes; j++ {
			env := envelopes[j]
			if j > i {
				// not the first item, will the last envelope fit in this one?
				if EnvWidth(env) > EnvWidth(lastEnv) && EnvHeight(env) > EnvHeight(lastEnv) {
					// last envelope would fit inside the current one-
					// add to the list of nested envelopes
					filteredEnvelopes = append(filteredEnvelopes, lastEnv)
					// ...and the current envelope becomes the new one
					// we're trying to find a fit for;
					// if we hit the end of input without
					// finding a fit for this one,
					// it gets tacked on below
					lastEnv = env
				}
			} else {
				// first item in this run,
				// always becomes the new initial
				// prior item for this run
				lastEnv = env
			}

			// if we're at the end of this run, tack on the last containing envelope;
			// this will either be the first envelope (because they all match)
			// or the most recent envelope that could contain at least one that came before
			if j == countOfEnvelopes-1 {
				filteredEnvelopes = append(filteredEnvelopes, lastEnv)
			}
		}

		// first pass or new larger nested collection?
		if i == 0 || len(filteredEnvelopes) > len(maxNestedEnvelopes) {
			// remember the new largest nested collection
			maxNestedEnvelopes = filteredEnvelopes
		}
	}

	// return the maximum nestable collection we found
	return maxNestedEnvelopes

}

//////////////////////////////////////////////////////////////////
