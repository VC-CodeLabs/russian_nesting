package main

//
// executes 1-N tests defined in a test case file;
// by default runs each specified test once with
// algorithm timings for each plus overall average time;
// command-line parameters support various options
// like repeating each test multiple times-
// use -h parameter for help with command-line parameters
//

import (
	"JeffR/lib"
	librn "JeffR/libsln"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// the basic definition of a test-
// the file containing the data and
// the expected result: correct max# of nested envelopes
type TestCase struct {
	filename string
	expected int
}

// a collection of test cases
type TestCases []TestCase

func main() {

	// how many times each test will be run- set with -r=#
	TEST_RUNS := 1
	// the single test file named in test cases to run vs all- set with -f=<filename>
	TEST_NAME := ""
	// the test cases file to use- set with -c=<filename>
	TEST_CASES_FILE := "testCases.txt"

	//
	// parse command-line parameters and set up testing conditions
	//

	verbosePtr := flag.Bool("v", lib.VERBOSE, "specifies whether to emit algorithm troubleshooting output")
	threadedPtr := flag.Bool("t", librn.THREADED, "specifies whether to use threading")
	repeatPtr := flag.Int("r", TEST_RUNS, "# of times to repeat each test")
	namePtr := flag.String("f", TEST_NAME, "a specific sample to test (as ref'd in test case file) vs all in test case file")
	caseFilePtr := flag.String("c", TEST_CASES_FILE, "use an alternative test case file")

	flag.Parse()

	if verbosePtr != nil {
		lib.VERBOSE = *verbosePtr
	}

	if threadedPtr != nil {
		librn.THREADED = *threadedPtr
	}

	if repeatPtr != nil {
		TEST_RUNS = *repeatPtr
	}

	if namePtr != nil {
		TEST_NAME = *namePtr
	}

	if caseFilePtr != nil {
		TEST_CASES_FILE = *caseFilePtr
	}

	//
	// test conditions resolved.
	//

	//
	// read in the test cases
	//

	testCases := make(TestCases, 0)
	file, err := os.Open(TEST_CASES_FILE)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		testSpec := strings.TrimSpace(scanner.Text())

		// ignore empty lines
		if len(testSpec) == 0 {
			continue
		}

		// comments supported in the test case file-
		// first non-whitespace char pound (or hash for the kids)
		if testSpec[0] == '#' {
			continue
		}

		// split the test case filename and expected answer
		testParts := strings.Split(testSpec, "=")

		// clean up and convert
		filename := strings.TrimSpace(testParts[0])

		// get the right-hand side of the test case
		// (aka expected answer value)
		// allowing for comments at the end of the line
		rhs := strings.TrimSpace(strings.Split(testParts[1], "#")[0])

		// get our expected answer for this test case as a number
		expected, err := strconv.Atoi(rhs)
		if err != nil {
			panic(err)
		}

		// add the test case
		testCases = append(testCases, TestCase{filename, expected})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	if len(testCases) == 0 {
		log.Fatal("No test cases found in ", TEST_CASES_FILE)
	}

	//
	// all test case definitions loaded
	//

	//
	// emit the results header with test conditions
	//

	fmt.Printf("*** Read %d test cases from %s\n", len(testCases), TEST_CASES_FILE)

	if TEST_RUNS > 1 {
		fmt.Printf("*** Repeating each test case %d times\n", TEST_RUNS)
	}

	fmt.Printf("*** Threading %s\n", lib.Ternary(librn.THREADED, "enabled", "disabled"))

	//
	// execute the tests
	//

	for _, testCase := range testCases {
		// run all test cases or the specific one
		// spec'd by -f=<filename> parameter
		// (as ref'd in test case file)
		if len(TEST_NAME) == 0 || TEST_NAME == testCase.filename {
			runTest(testCase, TEST_RUNS)
		}
	}

	// report the stats across all tests
	fmt.Printf("**** All tests completed in %15s in %10d total runs with an average run of %15s\n",
		lib.SanitizeDuration(lib.NanosToDuration(allTestRunsDuration)),
		allTestRunsCount,
		lib.SanitizeDuration(lib.NanosAvgToDuration(allTestRunsDuration, int64(allTestRunsCount))))
}

// track the duration for all test runs
var allTestRunsDuration int64 = 0

// track the total # of times we ran a test to compute average
var allTestRunsCount int64 = 0

func runTest(testCase TestCase, repeat int) {

	var runDuration int64 = 0
	for r := 0; r < repeat; r++ {

		//
		// note we only time the actual algorithm here
		//

		testStarted := lib.Start()
		testFile, err := os.Open(testCase.filename)

		if err != nil {
			log.Fatal(err)
		}

		// get the max nested envelopes collection from the algorithm impl
		envelopes := librn.GetNestedEnvelopes(testFile)
		testFile.Close()

		//
		// get the answer the russian nesting algorithm came up with-
		// the maximum # of nested envelopes from this test case
		//
		actual := len(envelopes)

		// how long did it take to read the data and get our answer
		testDuration := lib.Finish(testStarted)
		runDuration += testDuration.Nanoseconds()

		// track the stats across all runs for all tests
		allTestRunsDuration += runDuration
		allTestRunsCount++

		// describe our results
		result := fmt.Sprintf("tested %-30s expected: %6d actual: %6d in %15s",
			testCase.filename, testCase.expected, actual,
			lib.SanitizeDuration(testDuration))

		if r > 0 {
			result += fmt.Sprintf(" on run# %d", r+1)
		}

		// did the actual answer match what we expected?
		if testCase.expected != actual {
			// no, we got the wrong answer- report the error and bail out
			log.Fatal(result)
		} else {
			// yes, we got the right answer- report the result at least once
			// if r == 0 || repeat == 1 || lib.VERBOSE {
			fmt.Println(result)
			// }
		}
	}

	if repeat > 1 {
		// report the average of all runs
		fmt.Printf("****** %-30s x %10d repeat runs averaging %15s/run\n",
			testCase.filename, repeat,
			lib.SanitizeDuration(lib.NanosAvgToDuration(runDuration, int64(repeat))))
	}
}
