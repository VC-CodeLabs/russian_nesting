package main

import (
	"JeffR/lib"
	. "JeffR/libsln"
	"flag"
	"fmt"
	"slices"
	"time"
)

var TEST_RUNS int = 0

func main() {

	testRunPtr := flag.Int("t", TEST_RUNS, "specifies the # times to repeat tests")
	flag.Parse()
	if testRunPtr != nil {
		TEST_RUNS = *testRunPtr
	}
	if TEST_RUNS > 0 {
		testEnvelopeOps()
	}

	if false {
		testEnvMapping()
	}
	testArray()
	testSimpleMap()
	testInterface()
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

func start() time.Time {
	return lib.Start()
}

func finish(started time.Time) time.Duration {
	return lib.Finish(started)
}

func ternary[B bool, V int | int64 | float32 | float64 | string](cond bool, ifTrue V, ifFalse V) V {
	return lib.Ternary(cond, ifTrue, ifFalse)
}

func assert(cond bool, msg string) {
	lib.Assert(cond, msg)
}

func sanitizeDuration(d time.Duration) string {
	return lib.SanitizeDuration(d)
}

func nanosToDuration[N int | int64 | float64](nanos N) time.Duration {
	return lib.NanosToDuration(nanos)
}

func nanosAvgToDuration[N int | int64 | float64](totalNanos N, counter N) time.Duration {
	return lib.NanosAvgToDuration(totalNanos, counter)
}

//////////////////////////////////////////////////////////////////
// tests
//

func testInterface() {

	{
		ts := start()
		envs := EnvArrayWithAppend{}

		// assert(envs.foo == false, "foo!")
		envs.InitData()
		// assert(envs.foo == true, "foo!!!")

		for i := 0; i < ENV_MAX; i++ {
			envs.PutDataItem(ENV_MAX-i, ENV_MAX-i)
		}

		envs.CloseData()

		nestedEnvelopes := envs.GetNestedEnvelopes()

		td := finish(ts)

		fmt.Printf("Found %d nested envelopes in %s with append interface\n", len(nestedEnvelopes), td)
	}

	{
		ts := start()
		envs := EnvArrayPreAlloc{}

		// assert(envs.foo == false, "foo!")
		envs.InitData()
		// assert(envs.foo == true, "foo!!!")

		for i := 0; i < ENV_MAX; i++ {
			envs.PutDataItem(ENV_MAX-i, ENV_MAX-i)
		}

		envs.CloseData()

		nestedEnvelopes := envs.GetNestedEnvelopes()

		td := finish(ts)

		fmt.Printf("Found %d nested envelopes in %s with prealloc'd interface\n", len(nestedEnvelopes), td)
	}

	{
		ts := start()
		envs := EnvArrayPreAlloc{}

		// assert(envs.foo == false, "foo!")
		envs.InitData()
		// assert(envs.foo == true, "foo!!!")

		for i := 0; i < ENV_MAX; i++ {
			envs.PutDataItem(1, 1)
		}

		envs.CloseData()

		nestedEnvelopes := envs.GetNestedEnvelopes()

		td := finish(ts)

		fmt.Printf("Found %d nested envelopes in %s with prealloc'd interface\n", len(nestedEnvelopes), td)
	}

	{
		ts := start()
		envs := EnvArrayPreAlloc{}

		// assert(envs.foo == false, "foo!")
		envs.InitData()
		// assert(envs.foo == true, "foo!!!")

		for i := 0; i < ENV_MAX; i++ {
			envs.PutDataItem(1, ENV_MAX-i)
		}

		envs.CloseData()

		nestedEnvelopes := envs.GetNestedEnvelopes()

		td := finish(ts)

		fmt.Printf("Found %d nested envelopes in %s with prealloc'd interface\n", len(nestedEnvelopes), td)
	}

	{
		ts := start()
		envs := EnvArrayPreAlloc{}

		// assert(envs.foo == false, "foo!")
		envs.InitData()
		// assert(envs.foo == true, "foo!!!")

		for i := 0; i < ENV_MAX; i++ {
			envs.PutDataItem(ENV_MAX-i, 1)
		}

		envs.CloseData()

		nestedEnvelopes := envs.GetNestedEnvelopes()

		td := finish(ts)

		fmt.Printf("Found %d nested envelopes in %s with prealloc'd interface\n", len(nestedEnvelopes), td)
	}

}

func testArray() {
	ts := start()

	envelopes := make(Envelopes, ENV_MAX)
	// envSeen := make(EnvMapByStruct, ENV_MAX)

	for i := 0; i < ENV_MAX; i++ {
		// env := Envelope{1, 1}
		env := Envelope{ENV_MAX - i, ENV_MAX - i}
		// envelopes = append(envelopes, Envelope{ENV_MAX - i, ENV_MAX - i})
		// if !envSeen[env] {
		// envelopes = append(envelopes, env)
		envelopes[i] = env
		// envSeen[env] = true
		// }
	}

	// envelopes = envelopes[ENV_MAX-10:]

	sortedEnvelopes := envSort(envelopes)

	// assert(envelopes[0].Width == 1 && envelopes[0].Height == 1, "bad sort")

	filteredEnvelopes := envFilter(sortedEnvelopes)

	// assert(len(filteredEnvelopes) == len(envelopes), "bad filter")

	td := finish(ts)

	fmt.Printf("%d Envelopes in %s\n", len(filteredEnvelopes), sanitizeDuration(td))

}

func testSimpleMap() {

	ts := start()
	envMapByStruct := make(EnvMapByStruct, 0)

	for i := 0; i < ENV_MAX; i++ {
		envPut(&envMapByStruct, ENV_MAX-i, ENV_MAX-i)
		// envPut(&envMapByStruct, 1, 1)
	}

	envelopes := envKeys(envMapByStruct)
	sortedEnvelopes := envSort(envelopes)

	assert(envelopes[0].Width == 1 && envelopes[0].Height == 1, "bad sort")

	filteredEnvelopes := envFilter(sortedEnvelopes)

	assert(len(filteredEnvelopes) == len(envelopes), "bad filter")

	td := finish(ts)

	fmt.Printf("%d Envelopes via map in %s\n", len(filteredEnvelopes), sanitizeDuration(td))

}

type EnvMapByStruct map[Envelope]bool

func envPut(envMapByStruct *EnvMapByStruct, width int, height int) {

	(*envMapByStruct)[Envelope{Width: width, Height: height}] = true

}

func envKeys(envMapByStruct EnvMapByStruct) Envelopes {
	envelopes := make([]Envelope, len(envMapByStruct))
	i := 0
	for keys := range envMapByStruct {
		envelopes[i].Width = keys.Width
		envelopes[i].Height = keys.Height
		i++
	}
	return envelopes
}

/*****
func envCmp(a Envelope, b Envelope) int {
	diff := cmp.Compare(envWidth(a), envWidth(b))
	if diff == 0 {
		diff = cmp.Compare(envHeight(a), envHeight(b))
	}
	return diff
}

func envSortInPlace(envelopes *Envelopes) {
	slices.SortFunc(*envelopes, envCmp)
}

func envSort(envelopes Envelopes) Envelopes {
	slices.SortFunc(envelopes, envCmp)
	return envelopes
}

func envFilter(envelopes Envelopes) Envelopes {
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
*****/

// local proxies to minimize diffs during isolation work

func envSort(envelopes Envelopes) Envelopes {
	return EnvSort(envelopes)
}

func envFilter(envelopes Envelopes) Envelopes {
	return EnvFilter(envelopes)
}

// ////////////////////////////////////////////////////////////////
type EnvMap map[int]map[int]bool

// tracks original indexes for each envelope
type EnvMapWithIndex map[int]map[int][]int

func testEnvMapping() {

	envelopes := make(Envelopes, 0) // init the envelopes collections
	for i := 0; i < ENV_MAX; i++ {  // creating N max envelopes
		envelopes = append(envelopes, Envelope{i + 1, i + 1}) // append a new envelope
	}

	{
		mws := start()
		envMap := make(EnvMap, 0)
		for _, env := range envelopes {
			_, exists := envMap[envWidth(env)]
			if !exists {
				envMap[envWidth(env)] = make(map[int]bool, 0)
			}
			envMap[envWidth(env)][envHeight(env)] = true
		}
		mwd := finish(mws)
		fmt.Printf("Wrote %d envelopes to simple map in %s\n", len(envelopes), mwd)

		testEnvMapFilter(envMap)
	}

	{
		mws := start()
		envMapWI := make(EnvMapWithIndex, 0)
		for i, env := range envelopes {
			_, exists := envMapWI[envWidth(env)]
			if !exists {
				envMapWI[envWidth(env)] = make(map[int][]int, 0)
			}
			_, exists = envMapWI[envWidth(env)][envHeight(env)]
			if !exists {
				envMapWI[envWidth(env)][envHeight(env)] = make([]int, 0)
			}
			envMapWI[envWidth(env)][envHeight(env)] = append(envMapWI[envWidth(env)][envHeight(env)], i)

		}
		mwd := finish(mws)
		fmt.Printf("Wrote %d envelopes to map with index in %s\n", len(envelopes), mwd)

	}

	{
		envSorted := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			envSorted = append(envSorted, Envelope{ENV_MAX - i, ENV_MAX - i}) // append a new envelope
		}

		envCmp := func(a, b Envelope) int {
			if envWidth(a) < envWidth(b) {
				return -1
			} else if envWidth(a) > envWidth(b) {
				return 1
			} else {
				if envHeight(a) < envHeight(b) {
					return -1
				} else if envHeight(a) > envHeight(b) {
					return 1
				} else {
					return 0
				}
			}

		}

		{
			tsd := 0
			sr := 100
			for i := 0; i < sr; i++ {
				ss := start()
				slices.SortFunc(envSorted, envCmp)
				sd := finish(ss)

				assert(envWidth(envSorted[0]) == 1 && envHeight(envSorted[0]) == 1, "sorted [0]Envelopes failed")
				assert(envWidth(envSorted[ENV_MAX-1]) == DIM_MAX && envHeight(envSorted[ENV_MAX-1]) == DIM_MAX, "sorted [<last>]Envelopes failed")

				fmt.Printf("Sorted [%d] envelopes in %s\n", len(envSorted), sd)

				tsd += int(sd.Nanoseconds())
			}
			savg := nanosAvgToDuration(tsd, sr)
			fmt.Printf("*** Sorted [%d] envelopes average %s\n", len(envSorted), sanitizeDuration(savg))
		}

		{
			tsd := 0
			sr := 100
			for i := 0; i < sr; i++ {
				nss := start()
				slices.SortFunc(envSorted, envCmp)
				nsd := finish(nss)

				assert(envWidth(envSorted[0]) == 1 && envHeight(envSorted[0]) == 1, "sorted [0]Envelopes failed")
				assert(envWidth(envSorted[ENV_MAX-1]) == DIM_MAX && envHeight(envSorted[ENV_MAX-1]) == DIM_MAX, "sorted [<last>]Envelopes failed")

				fmt.Printf("Re-sorted already sorted [%d] envelopes in %s\n", len(envSorted), nsd)

				tsd += int(nsd.Nanoseconds())
			}
			savg := nanosAvgToDuration(tsd, sr)
			fmt.Printf("*** Re-sorted already sorted [%d] envelopes average %s\n", len(envSorted), sanitizeDuration(savg))
		}

	}

	{
		envUnique := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			envUnique = append(envUnique, Envelope{i + 1, i + 1}) // append a new envelope
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envUnique)
				fd := finish(fs)

				assert(len(envUnique) == len(filteredEnvelopes), "filter broken w/ all unique items")
				// assert(envUnique == filteredEnvelopes, "filter w/ all unique items list mismatch")

				fmt.Printf("Filtered %d unique items in %s\n", len(envUnique), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d unique items average %s\n", len(envUnique), sanitizeDuration(favg))

		}

		{
			envelopes := Envelopes{Envelope{50, 50}, Envelope{100, 100}}
			env2map := envelopesToMap(envelopes)
			for w, mi := range env2map {
				for h := range mi {
					fmt.Printf("e2m: %d, %d\n", w, h)
				}
			}
			filteredEnvelopes := testEnvMapFilter(env2map)
			for i := 0; i < len(filteredEnvelopes); i++ {
				fmt.Printf("Envelope => Map => Filtered %d: %d, %d\n", i, envWidth(filteredEnvelopes[i]), envHeight(filteredEnvelopes[i]))
			}
		}

		{
			envUniqueMap := make(EnvMap, 0)
			for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
				envUniqueMap = envelopeToMap(envUniqueMap, Envelope{i + 1, i + 1}) // append a new envelope
			}

			tmfd := 0
			fr := 100
			for i := 0; i < fr; i++ {

				mfs := start()
				filteredEnvelopes := testEnvMapFilter(envUniqueMap)
				mfd := finish(mfs)

				assert(len(envUnique) == len(filteredEnvelopes), "filter broken w/ all unique items by direct map")
				tmfd += int(mfd.Nanoseconds())
			}

			mfa := nanosAvgToDuration(tmfd, fr)
			fmt.Printf("*** Filtered %d unique items by direct map in %s\n", len(envUnique), sanitizeDuration(mfa))
		}
		{
			tmfd := 0
			fr := 100
			for i := 0; i < fr; i++ {
				mfs := start()

				filteredEnvelopes := testEnvMapFilter(envelopesToMap(envUnique))
				mfd := finish(mfs)

				assert(len(envUnique) == len(filteredEnvelopes), "filter broken w/ all unique items by map")
				// assert(envUnique == filteredEnvelopes, "filter w/ all unique items list mismatch")

				fmt.Printf("Filtered %d unique items by map in %s\n", len(envUnique), sanitizeDuration(mfd))

				tmfd += int(mfd.Nanoseconds())
			}

			mfa := nanosAvgToDuration(tmfd, fr)
			fmt.Printf("*** Filtered %d unique items by map in %s\n", len(envUnique), sanitizeDuration(mfa))

		}

	}

	{
		envSame := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			envSame = append(envSame, Envelope{1, 1}) // append a new envelope
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envSame)
				fd := finish(fs)

				assert(1 == len(filteredEnvelopes), "filter broken w/ all same items")
				assert(filteredEnvelopes[0] == envSame[0], "filtered all same items != first item")

				fmt.Printf("Filtered %d same items in %s\n", len(envSame), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d same items average %s\n", len(envSame), sanitizeDuration(favg))

		}

		{
			tmfd := 0
			fr := 100
			for i := 0; i < fr; i++ {
				mfs := start()

				filteredEnvelopes := testEnvMapFilter(envelopesToMap(envSame))
				mfd := finish(mfs)

				assert(1 == len(filteredEnvelopes), "filter broken w/ all same items")
				assert(filteredEnvelopes[0] == envSame[0], "filtered all same items != first item")

				fmt.Printf("Filtered %d same items by map in %s\n", len(envSame), sanitizeDuration(mfd))

				tmfd += int(mfd.Nanoseconds())
			}

			mfa := nanosAvgToDuration(tmfd, fr)
			fmt.Printf("*** Filtered %d same items by map in %s\n", len(envSame), sanitizeDuration(mfa))

		}

	}

	{
		envSameWidth := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			envSameWidth = append(envSameWidth, Envelope{1, i + 1}) // append a new envelope
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envSameWidth)
				fd := finish(fs)

				assert(1 == len(filteredEnvelopes), "filter broken w/ all same width items")
				assert(filteredEnvelopes[0] == envSameWidth[0], "filtered all same width items != first item")

				fmt.Printf("Filtered %d same width items in %s\n", len(envSameWidth), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d same width items average %s\n", len(envSameWidth), sanitizeDuration(favg))

		}

	}

	{
		envSameHeight := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			envSameHeight = append(envSameHeight, Envelope{i + 1, 1}) // append a new envelope
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envSameHeight)
				fd := finish(fs)

				assert(1 == len(filteredEnvelopes), "filter broken w/ all same height items")
				assert(filteredEnvelopes[0] == envSameHeight[0], "filtered all same height items != first item")

				fmt.Printf("Filtered %d same height items in %s\n", len(envSameHeight), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d same height items average %s\n", len(envSameHeight), sanitizeDuration(favg))

		}

	}

	{
		envFirstLastUnique := make(Envelopes, 0)
		for i := 0; i < ENV_MAX-1; i++ { // creating N max envelopes
			envFirstLastUnique = append(envFirstLastUnique, Envelope{1, 1}) // append a new envelope
		}
		envFirstLastUnique = append(envFirstLastUnique, Envelope{DIM_MAX, DIM_MAX}) // append a new envelope

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envFirstLastUnique)
				fd := finish(fs)

				assert(2 == len(filteredEnvelopes), "filter broken w/ first+last unique items")
				assert(filteredEnvelopes[0] == envFirstLastUnique[0], "filtered first+last unique items != first item")
				assert(filteredEnvelopes[len(filteredEnvelopes)-1] == envFirstLastUnique[len(envFirstLastUnique)-1],
					"filtered first+last unique items != last item")

				fmt.Printf("Filtered %d first+last unique items in %s\n", len(envFirstLastUnique), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d first+last unique items average %s\n", len(envFirstLastUnique), sanitizeDuration(favg))

		}

	}

	{
		envFirstTwoUnique := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			if i > 0 {
				envFirstTwoUnique = append(envFirstTwoUnique, Envelope{DIM_MAX, DIM_MAX}) // append a new envelope

			} else {
				envFirstTwoUnique = append(envFirstTwoUnique, Envelope{1, 1}) // append a new envelope
			}
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envFirstTwoUnique)
				fd := finish(fs)

				assert(2 == len(filteredEnvelopes), "filter broken w/ first two unique items")
				assert(filteredEnvelopes[0] == envFirstTwoUnique[0], "filtered first two unique items != first item")
				assert(filteredEnvelopes[1] == envFirstTwoUnique[1], "filtered first two unique items != second item")

				fmt.Printf("Filtered %d first two unique items in %s\n", len(envFirstTwoUnique), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d first two unique items average %s\n", len(envFirstTwoUnique), sanitizeDuration(favg))

		}

	}

	{
		envFirstThreeUnique := make(Envelopes, 0)
		for i := 0; i < ENV_MAX; i++ { // creating N max envelopes
			if i > 1 {
				envFirstThreeUnique = append(envFirstThreeUnique, Envelope{DIM_MAX, DIM_MAX}) // append a new envelope

			} else if i > 0 {

				envFirstThreeUnique = append(envFirstThreeUnique, Envelope{2, 2}) // append a new envelope

			} else {
				envFirstThreeUnique = append(envFirstThreeUnique, Envelope{1, 1}) // append a new envelope
			}
		}

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envFirstThreeUnique)
				fd := finish(fs)

				assert(3 == len(filteredEnvelopes), "filter broken w/ first three unique items")
				assert(filteredEnvelopes[0] == envFirstThreeUnique[0], "filtered first three unique items != first item")
				assert(filteredEnvelopes[1] == envFirstThreeUnique[1], "filtered first three unique items != second item")
				assert(filteredEnvelopes[2] == envFirstThreeUnique[2], "filtered first three unique items != third item")

				fmt.Printf("Filtered %d first three unique items in %s\n", len(envFirstThreeUnique), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d first three unique items average %s\n", len(envFirstThreeUnique), sanitizeDuration(favg))

		}

	}

	{
		envEdgeTwoUnique := make(Envelopes, 0)
		for i := 0; i < ENV_MAX-2; i++ { // creating N max envelopes
			if i > 0 {
				envEdgeTwoUnique = append(envEdgeTwoUnique, Envelope{2, 2}) // append a new envelope

			} else {
				envEdgeTwoUnique = append(envEdgeTwoUnique, Envelope{1, 1}) // append a new envelope
			}
		}
		envEdgeTwoUnique = append(envEdgeTwoUnique, Envelope{3, 3}) // append a new envelope
		envEdgeTwoUnique = append(envEdgeTwoUnique, Envelope{4, 4}) // append a new envelope

		{
			tfd := 0
			fr := 100

			for i := 0; i < fr; i++ {
				fs := start()
				filteredEnvelopes := testEnvAFilter(envEdgeTwoUnique)
				fd := finish(fs)

				assert(4 == len(filteredEnvelopes), "filter broken w/ edgex2 unique items")
				assert(filteredEnvelopes[0] == envEdgeTwoUnique[0], "filtered edgex2 unique items != first item")
				assert(filteredEnvelopes[1] == envEdgeTwoUnique[1], "filtered edgex2 unique items != second item")
				assert(filteredEnvelopes[2] == envEdgeTwoUnique[len(envEdgeTwoUnique)-2], "filtered edgex2 unique items != third item")
				assert(filteredEnvelopes[3] == envEdgeTwoUnique[len(envEdgeTwoUnique)-1], "filtered edgex2 unique items != fourth item")

				fmt.Printf("Filtered %d edgex2 items in %s\n", len(envEdgeTwoUnique), sanitizeDuration(fd))

				tfd += int(fd.Nanoseconds())
			}

			favg := nanosAvgToDuration(tfd, fr)
			fmt.Printf("*** Filtered %d edgesx2 items average %s\n", len(envEdgeTwoUnique), sanitizeDuration(favg))

		}

	}

}

func envelopeToMap(envMap EnvMap, env Envelope) EnvMap {

	if envMap == nil {
		envMap = make(EnvMap, 0)
	}

	_, exists := envMap[envWidth(env)]

	if !exists {
		envMap[envWidth(env)] = make(map[int]bool, 0)
	}
	envMap[envWidth(env)][envHeight(env)] = true

	return envMap

}

func envelopesToMap(envelopes []Envelope) EnvMap {
	envMap := make(EnvMap, 0)
	for _, env := range envelopes {
		_, exists := envMap[envWidth(env)]
		if !exists {
			envMap[envWidth(env)] = make(map[int]bool, 0)
		}
		envMap[envWidth(env)][envHeight(env)] = true
	}
	return envMap
}

func testEnvMapFilter(envMap EnvMap) []Envelope {

	assert(envMap != nil && len(envMap) > 0, "empty Envelope map, nothing to filter")

	filteredEnvelopes := make([]Envelope, 0)
	var lastEnv = Envelope{-1, -1}

	widths := make([]int, len(envMap))
	wi := 0
	for width := range envMap {
		widths[wi] = width
		wi++
	}
	slices.Sort(widths)

	i := 0
	for wi, width := range widths {
		widthHeights := envMap[width]
		heights := make([]int, len(widthHeights))
		hi := 0
		for height := range widthHeights {
			heights[hi] = height
			hi++
		}
		slices.Sort(heights)
		for hi, height := range heights {
			env := Envelope{width, height}

			// fmt.Printf("Checking [%d]: %d, %d\n", i, width, height)

			if i > 0 {
				if envWidth(env) > envWidth(lastEnv) && envHeight(env) > envHeight(lastEnv) {
					// last envelope would fit inside the current one
					filteredEnvelopes = append(filteredEnvelopes, lastEnv)
					lastEnv = env
				}

			} else {
				lastEnv = env
			}

			if wi == len(widths)-1 && hi == len(heights)-1 {
				filteredEnvelopes = append(filteredEnvelopes, lastEnv)
			}
			i++
		}

	}

	return filteredEnvelopes

}

func testEnvAFilter(envelopes []Envelope) []Envelope {

	assert(envelopes != nil && len(envelopes) > 0, "empty []Envelope collection, nothing to filter")

	filteredEnvelopes := make([]Envelope, 0)
	var lastEnv = Envelope{-1, -1}
	// var nextEnv = Envelope{DIM_MAX + 1, DIM_MAX + 1}
	for i, env := range envelopes {
		if i > 0 {
			if envWidth(env) > envWidth(lastEnv) && envHeight(env) > envHeight(lastEnv) {
				// last envelope would fit inside the current one
				filteredEnvelopes = append(filteredEnvelopes, lastEnv)
				/*
					if envWidth(env) < envWidth(nextEnv) && envHeight(env) < envHeight(nextEnv) {
						// wait for the tightest fit

					}
				*/
				/*
					if i == len(envelopes)-1 {
						filteredEnvelopes = append(filteredEnvelopes, env)
					}
				*/
				lastEnv = env
			}
		} else {
			lastEnv = env
		}
		if i == len(envelopes)-1 {
			filteredEnvelopes = append(filteredEnvelopes, lastEnv)
		}

	}

	if len(filteredEnvelopes) == 0 {
		filteredEnvelopes = append(filteredEnvelopes, lastEnv)
	}

	return filteredEnvelopes
}

// const TEST_RUNS = 100

func testEnvelopeOps() {

	allTestsStarted := start() // get the time we started testing

	testEnvelopeStructArrayInterOps()

	// NOTE order doesn't seem to matter here-
	// testing envelope-as-structs before or after envelope-as-arrays,
	// the struct version seems to be consistently faster
	// on my machine, at least :/

	testEnvelopesArrayOps()

	testEnvelopesStructOps()

	allTestsDuration := finish(allTestsStarted) // get total time to run all tests

	fmt.Printf("*** All tests completed in %s\n", allTestsDuration)

}

// generates a max-sized set of envelopes-as-array, then reads it all back to gauge performance
// compare to testEnvelopesStructOps()- a mirror, except for the envelope type
func testEnvelopesArrayOps() {

	var (
		writeTotalTime         int64 = 0 // total time spent writing across all test runs
		writeTotalCountNonZero int64 = 0 // # of write test runs w/ measurable duration > 0
		readTotalTime          int64 = 0 // total time spend reading across all test runs
		readTotalCountNonZero  int64 = 0 // # of read test runs w/ measurable duration > 0
	)

	// repeat the test N times (-t=# on the command-line, defaults to 10)
	for testRunIndex := 0; testRunIndex < TEST_RUNS; testRunIndex++ {

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// write tests: creating a collection of envelopes
		// since we have to read from stdin, this may be important
		fmt.Println("Writing test [][w#,h#] array...")
		writeTestStarted := start()      // start our timer for this write test run
		envelopes := make(EnvelopesA, 0) // init the envelopes collections
		for i := 0; i < ENV_MAX; i++ {   // creating N max envelopes
			envelopes = append(envelopes, EnvArray{i + 1, i + 1}) // append a new envelope
		}
		writeTestDuration := finish(writeTestStarted) // get the time taken for this write test run

		// report the results for this write test
		fmt.Printf("Array with %d envelopes [w#,h#] made in %15s\n", len(envelopes), sanitizeDuration(writeTestDuration))

		// track the total time & non-zero duration counter for all write tests
		writeTestNanos := writeTestDuration.Nanoseconds()
		writeTotalTime += writeTestNanos
		if writeTestNanos > 0 {
			writeTotalCountNonZero++
		}

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// read tests: reading a collection of envelopes
		fmt.Println("Reading test [][w#,h#] array...")
		readTestStarted := start()            // start our timer for this read test run
		for i := 0; i < len(envelopes); i++ { // walk thru the collection
			envIth := envelopes[i]       // envelope instance at index i
			envIWidth := envIth[WIDTH]   // width for this envelope instance
			envIHeight := envIth[HEIGHT] // height for this envelope instance
			// be sure the read doesn't get optimized out!!!
			if envIWidth > DIM_MAX || envIHeight > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}
		readTestDuration := finish(readTestStarted) // get the time taken for this read test run

		// report the results for this read test
		fmt.Printf("Array with %d envelopes [w#,h#] read in %15s\n", len(envelopes), sanitizeDuration(readTestDuration))

		// track the total time & non-zero duration counter for all read tests
		readTestNanos := readTestDuration.Nanoseconds()
		readTotalTime += readTestNanos
		if readTestNanos > 0 {
			readTotalCountNonZero++
		}
	}
	// report the results for all write tests
	writeTestsDuration := nanosToDuration(writeTotalTime)
	writeTestAverage := nanosAvgToDuration(writeTotalTime, int64(TEST_RUNS))
	writeTestAvgNonZero := nanosAvgToDuration(writeTotalTime, max(writeTotalCountNonZero, 1))
	fmt.Printf("*** Average time to make [%d][w#,h#] %15s / %6d = %15s (!0 %15s x %6d)\n",
		ENV_MAX,
		sanitizeDuration(writeTestsDuration), TEST_RUNS, sanitizeDuration(writeTestAverage),
		sanitizeDuration(writeTestAvgNonZero), writeTotalCountNonZero)

	// report the results for all read tests
	readTestsDuration := nanosToDuration(readTotalTime)
	readTestAverage := nanosAvgToDuration(readTotalTime, int64(TEST_RUNS))
	readTestAvgNonZero := nanosAvgToDuration(readTotalTime, max(readTotalCountNonZero, 1))
	fmt.Printf("*** Average time to read [%d][w#,h#] %15s / %6d = %15s (!0 %15s x %6d)\n",
		ENV_MAX,
		sanitizeDuration(readTestsDuration), TEST_RUNS, sanitizeDuration(readTestAverage),
		sanitizeDuration(readTestAvgNonZero), readTotalCountNonZero)
}

// generates a max-sized set of envelopes-as-struct, then reads it all back to gauge performance
// compare to testEnvelopesArrayOps()- a mirror, except for the envelope type
func testEnvelopesStructOps() {

	var (
		writeTotalTime         int64 = 0 // total time spent writing across all test runs
		writeTotalCountNonZero int64 = 0 // # of write test runs w/ measurable duration > 0
		readTotalTime          int64 = 0 // total time spend reading across all test runs
		readTotalCountNonZero  int64 = 0 // # of read test runs w/ measurable duration > 0
	)

	// repeat the test N times (-t=# on the command-line, defaults to 10)
	for testRunIndex := 0; testRunIndex < TEST_RUNS; testRunIndex++ {

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// write tests: creating a collection of envelopes
		// since we have to read from stdin, this may be important
		fmt.Println("Writing test []{w#,h#} array...")
		writeTestStarted := start()      // start our timer for this write test run
		envelopes := make(EnvelopesS, 0) // init the envelopes collections
		for i := 0; i < ENV_MAX; i++ {   // creating N max envelopes
			envelopes = append(envelopes, EnvStruct{i + 1, i + 1}) // append a new envelope
		}
		writeTestDuration := finish(writeTestStarted) // get the time taken for this write test run

		// report the results for this write test
		fmt.Printf("Array with %d envelopes {w#,h#} made in %15s\n", len(envelopes), sanitizeDuration(writeTestDuration))

		// track the total time & non-zero duration counter for all write tests
		writeTestNanos := writeTestDuration.Nanoseconds()
		writeTotalTime += writeTestNanos
		if writeTestNanos > 0 {
			writeTotalCountNonZero++
		}

		/////////////////////////////////////////////////////////////////////////////////////////////////////
		// read tests: reading a collection of envelopes
		fmt.Println("Reading test []{w#,h#} array...")
		readTestStarted := start()            // start our timer for this read test run
		for i := 0; i < len(envelopes); i++ { // walk thru the collection
			envIth := envelopes[i]      // envelope instance at index i
			envIWidth := envIth.Width   // width for this envelope instance
			envIHeight := envIth.Height // height for this envelope instance
			// be sure the read doesn't get optimized out!!!
			if envIWidth > DIM_MAX || envIHeight > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}

		readTestDuration := finish(readTestStarted) // get the time taken for this read test run

		// report the results for this read test
		fmt.Printf("Array with %d envelopes {w#,h#} read in %15s\n", len(envelopes), sanitizeDuration(readTestDuration))

		// track the total time & non-zero duration counter for all read tests
		readTestNanos := readTestDuration.Nanoseconds()
		readTotalTime += readTestNanos
		if readTestNanos > 0 {
			readTotalCountNonZero++
		}

	}

	// report the results for all write tests
	writeTestsDuration := nanosToDuration(writeTotalTime)
	writeTestAverage := nanosAvgToDuration(writeTotalTime, int64(TEST_RUNS))
	writeTestAvgNonZero := nanosAvgToDuration(writeTotalTime, max(writeTotalCountNonZero, 1))
	fmt.Printf("*** Average time to make [%d]{w#,h#} %15s / %6d = %15s (!0 %15s x %6d)\n",
		ENV_MAX, sanitizeDuration(writeTestsDuration), TEST_RUNS, sanitizeDuration(writeTestAverage), sanitizeDuration(writeTestAvgNonZero), writeTotalCountNonZero)

	// report the results for all read tests
	readTestsDuration := nanosToDuration(readTotalTime)
	readTestAverage := nanosAvgToDuration(readTotalTime, int64(TEST_RUNS))
	readTestAvgNonZero := nanosAvgToDuration(readTotalTime, max(readTotalCountNonZero, 1))
	fmt.Printf("*** Average time to read [%d]{w#,h#} %15s / %6d = %15s (!0 %15s x %6d)\n",
		ENV_MAX, sanitizeDuration(readTestsDuration), TEST_RUNS, sanitizeDuration(readTestAverage), sanitizeDuration(readTestAvgNonZero), readTotalCountNonZero)
}

// basic validation to "prove" we can use the same inline constant defs for envelope-as-array vs envelope-as-struct,
// and the results are equivalent w/r/t the state of the envelope(s)
func testEnvelopeStructArrayInterOps() {

	// verifies we can init either struct or array with the same expression

	var envA = EnvelopesA{{1, 2}, {3, 4}} // where each envelope is an array
	var envS = EnvelopesS{{1, 2}, {3, 4}} // where each envelope is a struct

	// verify the results are equivalent :)

	// array and struct are same length
	assert(len(envA) == len(envS), "mismatched array vs struct lengths")

	for i := 0; i < len(envA); i++ {
		// equivalent width values
		assert(envA[i][WIDTH] == envS[i].Width, fmt.Sprintf("mismatched env[%d] width", i))
		// equivalent height values
		assert(envA[i][HEIGHT] == envS[i].Height, fmt.Sprintf("mismatched env[%d] height", i))
	}
}
