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

	envelopes := getNestedEnvelopes(os.Stdin)

	fmt.Printf("%d", len(envelopes))
}

func getNestedEnvelopes(file *os.File) Envelopes {
	rnMgr := IRussianNestingManager{}

	rnMgr.InitData()
	readEnvelopeData(file, rnMgr.PutDataItem)
	rnMgr.CloseData()

	envelopes := rnMgr.GetNestedEnvelopes()

	return envelopes
}

func readEnvelopeData(file *os.File, put func(w int, h int)) {

	if VERBOSE {
		fmt.Println("reading from stdin...")
	}

	stdin, err := io.ReadAll(file)

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

	mode := 0
	width := -1
	height := -1
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
			width, err = strconv.Atoi(strings.TrimSpace(tail[0:n]))
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
			height, err = strconv.Atoi(strings.TrimSpace(tail[0:n]))
			if err != nil {
				panic(err)
			}
			if VERBOSE {
				fmt.Printf("Storing [ %d, %d ]\n", width, height)
			}
			put(width, height)
			mode = 1
			i += n
		}

		if mode < 0 {
			break
		}

	}

}
