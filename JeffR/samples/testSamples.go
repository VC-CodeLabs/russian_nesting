package main

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

type TestCase struct {
	filename string
	expected int
}

type TestCases []TestCase

var TEST_RUNS int = 1
var TEST_NAME string = ""

func main() {

	verbosePtr := flag.Bool("v", lib.VERBOSE, "specifies whether to emit troubleshooting output")
	threadedPtr := flag.Bool("t", librn.THREADED, "specifies whether to use threading")
	repeatPtr := flag.Int("r", TEST_RUNS, "# of times to repeat each test")
	namePtr := flag.String("f", TEST_NAME, "specific test to run")
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

	testCases := make(TestCases, 0)
	file, err := os.Open("answerKey.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		testSpec := scanner.Text()
		if strings.TrimSpace(testSpec)[0] == '#' {
			continue
		}
		testParts := strings.Split(testSpec, "=")
		expected, err := strconv.Atoi(strings.TrimSpace(testParts[1]))
		if err != nil {
			panic(err)
		}
		testCases = append(testCases, TestCase{testParts[0], expected})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Read %d test cases from answerKey.txt\n", len(testCases))

	if TEST_RUNS > 1 {
		fmt.Printf("Repeating each test case %d times\n", TEST_RUNS)
	}

	fmt.Printf("Threaded: %t\n", librn.THREADED)

	for _, testCase := range testCases {
		if len(TEST_NAME) == 0 || TEST_NAME == testCase.filename {
			runTest(testCase)
		}
	}

	fmt.Printf("All tests completed in %s in %d total runs with an average run of %s\n",
		lib.NanosToDuration(allRunsDuration), allRunsCount,
		lib.NanosAvgToDuration(allRunsDuration, int64(allRunsCount)))
}

var allRunsDuration int64 = 0
var allRunsCount = 0

func runTest(testCase TestCase) {

	var runDuration int64 = 0
	for r := 0; r < TEST_RUNS; r++ {
		ts := lib.Start()
		testFile, err := os.Open(testCase.filename)

		if err != nil {
			log.Fatal(err)
		}

		envelopes := librn.GetNestedEnvelopes(testFile)

		testFile.Close()
		td := lib.Finish(ts)
		runDuration += td.Nanoseconds()
		allRunsDuration += runDuration
		allRunsCount++

		actual := len(envelopes)

		result := fmt.Sprintf("testing %s expected %d actual %d in %s",
			testCase.filename, testCase.expected, actual, td)

		if testCase.expected != actual {
			log.Fatal(result)
		} else {
			fmt.Println(result)
		}
	}

	if TEST_RUNS > 1 {
		fmt.Printf("Completed %d runs of %s averaging %s/run\n",
			TEST_RUNS, testCase.filename, lib.NanosAvgToDuration(runDuration, int64(TEST_RUNS)))
	}
}
