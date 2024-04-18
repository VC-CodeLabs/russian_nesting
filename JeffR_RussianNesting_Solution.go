package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var TEST_RUNS int = 10

func main() {
	testRunPtr := flag.Int("t", TEST_RUNS, "specifies the # times to repeat tests")
	flag.Parse()
	if testRunPtr != nil {
		TEST_RUNS = *testRunPtr
	}
	testEnvelopeOps()
}

const DIM_MIN = 1

// var DIM_MAX = int(math.Pow(10, 5))
// var DIM_MAX = int(math.Pow10(5))
const DIM_MAX = int(1e5)

const ENV_MIN = 1

// var ENV_MAX = DIM_MAX
const ENV_MAX = DIM_MAX

type EnvStruct struct {
	width  int
	height int
}

type EnvelopesS []EnvStruct

type EnvArray [2]int

// indexes into EnvArray for width & height
const (
	WIDTH  = 0
	HEIGHT = 1
)

type EnvelopesA [][2]int

// returns time.Now()
//
// use along with finish() to time some code:
//
//		started := start()
//	 	// ...some code...
//	 	elapsed := finish(started)
func start() time.Time {
	return time.Now()
}

// returns duration since prior start() or time.Time;
// use along with start() to time some code
func finish(started time.Time) time.Duration {
	return time.Since(started)
}

// provides an approximation of good ol' Java ternary operator:
//
//	<condBool> ? <valueIfTrue> : <valueIfFalse>
//
// the equivalent of
//
//	if <condBool> {
//		return <valueIfTrue>
//	} else {
//		return <valueIfFalse>
//	}
//
// (in fact, this is the impl in golang)
// *BUT* is nestable inside another expression e.g.
//
//	fmt.Printf("%s", ternary(flag,"foo","bar"))
//
// vs
//
//	     if flag {
//				fmt.Print("foo")
//			} else {
//				fmt.Print("bar")
//			}
//
// with one *MAJOR* difference: the <valueIfTrue> and <valueIfFalse>
// are both eval'd regardless of the <condBool> outcome
// SOOO don't use anything that has side-effects for 2nd or 3rd args!!!
//
// # Java:
//
//	int y = 1;
//	int z = 1;
//	int a = y < 10 ? y++ : z++;
//	System.out.println( "y=" + y + " z=" + z );
//	// console output: y=2 z=1
//
// # Golang:
//
//	y := 1
//	z := 1
//	a := ternary( y < 10, y++, z++ )
//	fmt.Printf( "y=%d z=%d\n", y, z)
//	// console output: y=2 z=2
func ternary[B bool, V int | int64 | float32 | float64 | string](cond bool, ifTrue V, ifFalse V) V {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// the rough equivalent of java assert,
// minus the nifty auto-stringification of check expression (1st arg)
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// sanitizes stringified duration so the microsecond indicator
// can pass thru console pipes like find cmd;
// otherwise, it prints as "#┬╡s" instead of "#µs"
func sanitizeDuration(d time.Duration) string {
	s := fmt.Sprintf("%s", d)
	s = strings.Replace(s, "µs", "us", 1)
	return s
}

// converts nanos# to equivalent duration
func nanosToDuration[N int | int64 | float64](nanos N) time.Duration {
	return time.Duration(float64(nanos) * float64(time.Nanosecond))
}

// computes the average time from total / counter, converts to duration
func nanosAvgToDuration[N int | int64 | float64](totalNanos N, counter N) time.Duration {
	return nanosToDuration(float64(totalNanos) / float64(counter))
}

//////////////////////////////////////////////////////////////////
// tests
//

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
			envIWidth := envIth.width   // width for this envelope instance
			envIHeight := envIth.height // height for this envelope instance
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
		assert(envA[i][WIDTH] == envS[i].width, fmt.Sprintf("mismatched env[%d] width", i))
		// equivalent height values
		assert(envA[i][HEIGHT] == envS[i].height, fmt.Sprintf("mismatched env[%d] height", i))
	}
}
