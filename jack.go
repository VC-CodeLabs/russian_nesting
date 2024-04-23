package main

import (
	"fmt"
	"sort"
	"strings"
)

var max = 1

func main() {
	var input string
	fmt.Scanln(&input)

	envelopes := parseStringToArray(input)

	sortArrayByWidthThenHeight(envelopes)

	countNestedEnvelopes(envelopes)

	sortArrayByHeightThenWidth(envelopes)

	countNestedEnvelopes(envelopes)

	fmt.Println(max)
}

func countNestedEnvelopes(envelopes [][]int) {
	count := 1
	envelope1 := envelopes[0]

	for i := 1; i < len(envelopes); i++ {
		envelope2 := envelopes[i]

		if envelopeFits(envelope1, envelope2) {
			envelope1 = envelopes[i]
			count++
			if count > max {
				max = count
			}
		}
	}
}

func parseStringToArray(input string) [][]int {
	input = input[2 : len(input)-2] // Remove outer brackets
	pairs := strings.Split(input, "],[") // Split by "],["

	result := make([][]int, len(pairs))

	for i, pair := range pairs {
		values := strings.Split(pair, ",") // Split each pair by ","
		result[i] = make([]int, len(values))

		for j, val := range values {
			result[i][j] = parseInt(val) // Parse each value to int
		}
	}

	return result
}

func sortArrayByWidthThenHeight(array [][]int) {
	sort.Slice(array, func(i, j int) bool {
		// Compare width values first
		if array[i][0] != array[j][0] {
			return array[i][0] < array[j][0]
		}
		// If width values are equal, compare height values
		return array[i][1] < array[j][1]
	})
}

func sortArrayByHeightThenWidth(array [][]int) {
	sort.Slice(array, func(i, j int) bool {
		// Compare height values first
		if array[i][1] != array[j][1] {
			return array[i][1] < array[j][1]
		}
		// If height values are equal, compare width values
		return array[i][0] < array[j][0]
	})
}

func envelopeFits(e1, e2 []int) bool {
	return e1[0] < e2[0] && e1[1] < e2[1]
}

func parseInt(s string) int {
	var res int
	fmt.Sscanf(s, "%d", &res)
	return res
}
