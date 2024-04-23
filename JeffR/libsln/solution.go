package librn

//
// the solution impl, minus most all the layers of interfaces, details about where the data came from, &c
//
// I considered using a map (to eliminate duplicates),
// but testing suggests maps in general are slower than arrays / slices
// not to mention how much more complicated it felt and looked doing so
//

import (
	"cmp"
	"slices"
	"sync"
)

// compare two envelopes to support sorting by size (width then height)
func CompareEnvelopes(a Envelope, b Envelope) int {
	diff := cmp.Compare(EnvWidth(a), EnvWidth(b))
	if diff == 0 {
		diff = cmp.Compare(EnvHeight(a), EnvHeight(b))
	}
	return diff
}

// sort the collection by width then height and return it
func SortEnvelopes(envelopes Envelopes) Envelopes {
	// TODO how expensive is this vs in-place sort?
	slices.SortFunc(envelopes, CompareEnvelopes)
	return envelopes
}

// strip out the duplicates (input must be sorted!!!)
//
// this significantly improves performance for
// e.g. max # envelopes all matching size
func CompactEnvelopes(envelopes Envelopes) Envelopes {
	return slices.Compact(envelopes)
}

// support threading, where each pass (w/ different starting point) is its' own thread
var (
	mu                 sync.Mutex
	maxNestings        int
	maxNestedEnvelopes Envelopes
	wg                 sync.WaitGroup
)

// threading is optional
var THREADED bool = true

// get the maximum # of nested envelopes in the specified collection;
//
// the input *MUST* be sorted by width then height and ideally compacted (duplicates removed) as well
func FindMaxNestedEnvelopes(envelopes Envelopes) Envelopes {

	// fmt.Printf("Threaded: %t\n", THREADED)

	// lib.Assert(envelopes != nil && len(envelopes) > 0, "empty []Envelope collection, nothing to filter")

	if len(envelopes) <= 1 { // len(nil) == 0
		return envelopes
	}

	maxNestings = 0
	maxNestedEnvelopes = make(Envelopes, 0)

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

	// how many inputs are we dealing with (fighting for every µs)
	countOfEnvelopes := len(envelopes)

	// walk thru the list of inputs, using each position as new starting point
	for startingOffset := range envelopes {

		// threading really uglifies this :/

		if THREADED {
			mu.Lock()
		}

		// the following check may not make much difference for most test cases,
		// especially if threading is enabled, but definitely does when most or all can nest
		cantBeABiggerNestingLeft := startingOffset > 0 && maxNestings > countOfEnvelopes-startingOffset

		if THREADED {
			mu.Unlock()
		}

		if cantBeABiggerNestingLeft {
			// can't possibly find a bigger one, don't bother looking
			break
		}

		if THREADED {
			// starting a new thread, add it to the waitgroup
			wg.Add(1)
			// start the thread for this section starting at i'th offset
			go getNestedEnvelopesInTail(envelopes, countOfEnvelopes, startingOffset)
		} else {
			getNestedEnvelopesInTail(envelopes, countOfEnvelopes, startingOffset)
		}

	}

	if THREADED {
		// wait for all the threads to complete
		wg.Wait()
	}

	// return the maximum nestable collection we found
	return maxNestedEnvelopes

}

// get the nested envelopes for the tail of the input starting at startingOffset'th item;
// input collection *MUST* be sorted by width then height!!!
func getNestedEnvelopesInTail(envelopes Envelopes, countOfEnvelopes int, startingOffset int) {

	// track the set of nested envelopes
	nestedEnvelopes := make(Envelopes, 0)
	// track the first/last containing envelope so we can see if it will fit inside a subsequent envelope
	var lastContainingEnv = Envelope{-1, -1}
	// walk the tail: starting at the offset thru the remainder of the collection
	for ndx := startingOffset; ndx < countOfEnvelopes; ndx++ {
		// get the envelope at this index
		currEnvelope := envelopes[ndx]
		if ndx > startingOffset {
			// not the first item, will the last containing envelope fit in this one?
			if EnvWidth(currEnvelope) > EnvWidth(lastContainingEnv) && EnvHeight(currEnvelope) > EnvHeight(lastContainingEnv) {
				// last containing envelope would fit inside the current one-
				// add to the list of nested envelopes
				nestedEnvelopes = append(nestedEnvelopes, lastContainingEnv)
				// ...and the current envelope
				// becomes the new container
				// we're trying to find a fit for;
				// if we hit the end of input without
				// finding a fit for this one,
				// it gets tacked on below
				lastContainingEnv = currEnvelope
			}
		} else {
			// first item in this run,
			// always becomes the new initial containing (empty) envelope
			// we're looking for the next envelope this fits inside of, if any
			lastContainingEnv = currEnvelope
		}

		// if we're at the end of this run, tack on the last containing envelope;
		// this will either be the first envelope (because they all match)
		// or the most recent envelope that could contain at least one that came before
		// (not necessarily the last envelope in the input collection)
		if ndx == countOfEnvelopes-1 {
			nestedEnvelopes = append(nestedEnvelopes, lastContainingEnv)
		}
	}

	if THREADED {
		mu.Lock()
	}

	// first pass (if not threading) or new larger nested collection?
	if (!THREADED && startingOffset == 0) || len(nestedEnvelopes) > maxNestings {
		// remember the new largest nested collection (so far)
		maxNestedEnvelopes = nestedEnvelopes
		maxNestings = len(maxNestedEnvelopes)
	}

	if THREADED {
		mu.Unlock()
		wg.Done()
	}

}

//////////////////////////////////////////////////////////////////
