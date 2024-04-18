package main

import (
	"fmt"
	"strings"
	"time"
)

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

func main() {
	testEnvelopeOps()
}

func start() time.Time {
	return time.Now()
}

func finish(started time.Time) time.Duration {
	return time.Since(started)
}

// provides an approximation of good ol' Java ternary operator:
// <condBool> ? <valueIfTrue> : <valueIfFalse>
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

func nanosToDuration(nanos int64) time.Duration {
	return time.Duration(float64(nanos) * float64(time.Nanosecond))
}

//////////////////////////////////////////////////////////////////
// tests
//

const TEST_RUNS = 100

func testEnvelopeOps() {

	testEnvelopeStructArrayInterOps()

	testEnvelopesArrayOps()

	testEnvelopesStructOps()

}

func testEnvelopesArrayOps() {

	var (
		writeTotalTime         int64 = 0 // total time spent writing across all test runs
		writeTotalCountNonZero int64 = 0 // # of write test runs w/ measurable duration > 0
		readTotalTime          int64 = 0 // total time spend reading across all test runs
		readTotalCountNonZero  int64 = 0 // # of read test runs w/ measurable duration > 0
	)

	for testRunIndex := 0; testRunIndex < TEST_RUNS; testRunIndex++ {
		fmt.Println("Writing test [][w#,h#] array...")
		writeTestStarted := start()
		envs := make(EnvelopesA, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvArray{i + 1, i + 1})

		}
		writeTestDuration := finish(writeTestStarted)
		fmt.Printf("Array with %d envelopes [w#,h#] made in %15s\n", len(envs), sanitizeDuration(writeTestDuration))
		writeTestNanos := writeTestDuration.Nanoseconds()
		writeTotalTime += writeTestNanos
		if writeTestNanos > 0 {
			writeTotalCountNonZero++
		}

		fmt.Println("Reading test [][w#,h#] array...")
		readTestStarted := start()
		for i := 0; i < len(envs); i++ {
			envIth := envs[i] // envelope at index i
			envIWidth := envIth[WIDTH]
			envIHeight := envIth[HEIGHT]
			// be sure the read doesn't get optimized out!!!
			if envIWidth > DIM_MAX || envIHeight > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}
		readTestDuration := finish(readTestStarted)
		fmt.Printf("Array with %d envelopes [w#,h#] read in %15s\n", len(envs), sanitizeDuration(readTestDuration))
		readTestNanos := readTestDuration.Nanoseconds()
		readTotalTime += readTestNanos
		if readTestNanos > 0 {
			readTotalCountNonZero++
		}
	}
	writeTestsDuration := nanosToDuration(writeTotalTime)
	writeTestAverage := nanosToDuration(writeTotalTime / TEST_RUNS)
	writeTestAvgNonZero := nanosToDuration(writeTotalTime / max(writeTotalCountNonZero, 1))
	fmt.Printf("*** Average time to make [%d][w#,h#] %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX,
		sanitizeDuration(writeTestsDuration), TEST_RUNS, sanitizeDuration(writeTestAverage),
		sanitizeDuration(writeTestAvgNonZero), writeTotalCountNonZero)

	readTestsDuration := nanosToDuration(readTotalTime)
	readTestAverage := nanosToDuration(readTotalTime / TEST_RUNS)
	readTestAvgNonZero := nanosToDuration(readTotalTime / max(readTotalCountNonZero, 1))
	fmt.Printf("*** Average time to read [%d][w#,h#] %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX,
		sanitizeDuration(readTestsDuration), TEST_RUNS, sanitizeDuration(readTestAverage),
		sanitizeDuration(readTestAvgNonZero), readTotalCountNonZero)
}

func testEnvelopesStructOps() {

	var (
		writeTotalTime         int64 = 0 // total time spent writing across all test runs
		writeTotalCountNonZero int64 = 0 // # of write test runs w/ measurable duration > 0
		readTotalTime          int64 = 0 // total time spend reading across all test runs
		readTotalCountNonZero  int64 = 0 // # of read test runs w/ measurable duration > 0
	)

	for testRunIndex := 0; testRunIndex < TEST_RUNS; testRunIndex++ {
		fmt.Println("Writing test []{w#,h#} array...")
		writeTestStarted := start()
		envs := make(EnvelopesS, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvStruct{i + 1, i + 1})

		}
		writeTestDuration := finish(writeTestStarted)
		fmt.Printf("Array with %d envelopes {w#,h#} made in %15s\n", len(envs), sanitizeDuration(writeTestDuration))
		writeTestNanos := writeTestDuration.Nanoseconds()
		writeTotalTime += writeTestNanos
		if writeTestNanos > 0 {
			writeTotalCountNonZero++
		}

		fmt.Println("Reading test []{w#,h#} array...")
		readTestStarted := start()
		for i := 0; i < len(envs); i++ {
			envIth := envs[i] // envelope at index i
			envIWidth := envIth.width
			envIHeight := envIth.height
			// be sure the read doesn't get optimized out!!!
			if envIWidth > DIM_MAX || envIHeight > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}
		readTestDuration := finish(readTestStarted)
		fmt.Printf("Array with %d envelopes {w#,h#} read in %15s\n", len(envs), sanitizeDuration(readTestDuration))
		readTestNanos := readTestDuration.Nanoseconds()
		readTotalTime += readTestNanos
		if readTestNanos > 0 {
			readTotalCountNonZero++
		}

	}

	writeTestsDuration := nanosToDuration(writeTotalTime)
	writeTestAverage := nanosToDuration(writeTotalTime / TEST_RUNS)
	writeTestAvgNonZero := nanosToDuration(writeTotalTime / max(writeTotalCountNonZero, 1))
	fmt.Printf("*** Average time to make [%d]{w#,h#} %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(writeTestsDuration), TEST_RUNS, sanitizeDuration(writeTestAverage), sanitizeDuration(writeTestAvgNonZero), writeTotalCountNonZero)

	readTestsDuration := nanosToDuration(readTotalTime)
	readTestAverage := nanosToDuration(readTotalTime / TEST_RUNS)
	readTestAvgNonZero := nanosToDuration(readTotalTime / max(readTotalCountNonZero, 1))
	fmt.Printf("*** Average time to read [%d]{w#,h#} %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(readTestsDuration), TEST_RUNS, sanitizeDuration(readTestAverage), sanitizeDuration(readTestAvgNonZero), readTotalCountNonZero)
}

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
