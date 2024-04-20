package main

import (
	. "JeffR/libsln"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var VERBOSE bool = false

func main() {

	verbosePtr := flag.Bool("v", VERBOSE, "specifies whether to emit troubleshooting output")
	flag.Parse()
	if verbosePtr != nil {
		VERBOSE = *verbosePtr
	}

	if VERBOSE {
		fmt.Println("reading from stdin...")
	}

	stdin, err := io.ReadAll(os.Stdin)

	if VERBOSE {
		fmt.Println("...read from stdin.")
	}

	if err != nil {
		panic(err)
	}

	input := string(stdin)

	if VERBOSE {
		log.Printf("input: `%s`\n", input)
	}

	envelopes := getNestedEnvelopes(input)

	fmt.Printf("%d", len(envelopes))
}

func getNestedEnvelopes(input string) Envelopes {
	rnMgr := IRussianNestingManager{}

	processData(&rnMgr, input)

	envelopes := rnMgr.GetNestedEnvelopes()

	return envelopes
}

func processData(rnMgr *IRussianNestingManager, input string) {
	rnMgr.InitData()

	mode := 0
	width := -1
	height := -1
	var err error
	for i := 0; i < len(input); i++ {
		tail := input[i:]
		if VERBOSE {
			log.Printf("mode: %d tail: %s\n", mode, tail)
		}
		switch mode {
		case 0, 1: // 0 == opening Envelopes collection, 1 == start of Envelope expected
			n := strings.IndexAny(tail, "[{")
			if n < 0 {
				final := strings.TrimSpace(tail)
				if !(final == "]" || final == "}") {
					panic(fmt.Sprintf("Unexpected tail `%s` at offset %d", final, i))
				}
				mode = -1
				break
			}
			mode += 1
			i += n
		case 2: // 2 == width expected
			n := strings.IndexAny(tail, ",")
			if n < 0 {
				panic(fmt.Sprintf("Missing `,` delimiter in input at offset %d", i))
			}
			width, err = strconv.Atoi(tail[0:n])
			if err != nil {
				panic(err)
			}
			mode++
			i += n
		case 3: // 3 == height expected
			n := strings.IndexAny(tail, "}]")
			if n < 0 {
				panic(fmt.Sprintf("Missing closing bracket/brace in input at offset %d", i))
			}
			height, err = strconv.Atoi(tail[0:n])
			if err != nil {
				panic(err)
			}
			if VERBOSE {
				fmt.Printf("Storing [ %d, %d ]\n", width, height)
			}
			rnMgr.PutDataItem(width, height)
			mode = 1
			i += n
		}

		if mode < 0 {
			break
		}

	}

	rnMgr.CloseData()

}
