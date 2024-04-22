package main

import (
	librn "JeffR/libsln"
	"bufio"
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

func main() {

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

	for _, testCase := range testCases {
		runTest(testCase)
	}
}

func runTest(testCase TestCase) {
	testFile, err := os.Open(testCase.filename)

	if err != nil {
		log.Fatal(err)
	}

	envelopes := librn.GetNestedEnvelopes(testFile)

	actual := len(envelopes)

	result := fmt.Sprintf("testing %s expected %d actual %d", testCase.filename, testCase.expected, actual)

	if testCase.expected != actual {
		log.Fatal(result)
	} else {
		fmt.Println(result)
	}
}
