package main

import (
	. "JeffR/lib"
	. "JeffR/libsln"
	"flag"
	"fmt"
	"os"
)

func main() {

	verbosePtr := flag.Bool("v", VERBOSE, "specifies whether to emit troubleshooting output")
	flag.Parse()

	if verbosePtr != nil {
		VERBOSE = *verbosePtr
	}

	// get the max nested envelopes from stdin
	envelopes := GetNestedEnvelopes(os.Stdin)

	if VERBOSE {
		fmt.Println("Max Nested Envelopes:")
		fmt.Print("[ ")
		for i, envelope := range envelopes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("[ %d, %d ]", EnvWidth(envelope), EnvHeight(envelope))

		}
		fmt.Println(" ]")
		fmt.Printf("Max Nested Envelope Count: %d", len(envelopes))
	} else {
		// standard output is just the # of nested envelopes
		fmt.Printf("%d", len(envelopes))
	}
}

/*****
// get the max nested envelope collection from
// a set of envelopes defined in an input file
// (which may be stdin!)
func getNestedEnvelopes(file *os.File) Envelopes {
	// instantiate our manager
	rnMgr := IRussianNestingManager{}

	// prep data storage
	rnMgr.InitData()
	// feed input data into our manager
	readEnvelopeDatafile(file, rnMgr.PutDataItem)
	// finalize data storage
	rnMgr.CloseData()

	// get the max nested envelope collection
	envelopes := rnMgr.GetNestedEnvelopes()

	return envelopes
}

// read envelope data from a file and put it somewhere
func readEnvelopeDatafile(file *os.File, putfx func(w int, h int)) {

	if VERBOSE {
		fmt.Printf("reading from source %s...\n", file.Name())
	}

	// given abundant memory and the constraints placed on input,
	// we can probably get away with just reading all the contents in at once
	stdin, err := io.ReadAll(file)

	if VERBOSE {
		fmt.Println("...read from source.")
	}

	// we had an issue reading data, bail out
	if err != nil {
		panic(err)
	}

	// transform the byte array read from the file into a string
	input := string(stdin)

	readEnvelopeData(input, putfx)
}

func readEnvelopeData(input string, putfx func(w int, h int)) {

	if VERBOSE {
		if len(input) > 200 {
			log.Printf("input: %d chars\n", len(input))
		} else {
			log.Printf("input: `%s`\n", input)
		}
	}

	//
	// NOTE when processing the input, we support brackets [] or braces {}
	// supporting braces means we could copy & paste between external file and inline def in source code
	//

	// define the different modes used when reading the input
	const (
		// end of file
		EOF = -1
		// looking for start of collection aka first bracket/brace
		COLL = 0
		// looking for start of a given envelope
		ENV = 1
		// looking for width inside an envelope
		WIDTH = 2
		// looking for height inside an envelope
		HEIGHT = 3
	)

	// initially we're looking for the start of a collection
	mode := COLL

	// track the width and height so when we find the pair we can store them
	width := -1
	height := -1
	var err error
	for i := 0; i < len(input); i++ {

		// get the tail of the input starting at the current offset
		tail := input[i:]

		if VERBOSE {
			if len(tail) > 40 {
				log.Printf("mode: %d offset: %d tail: `%s`...\n", mode, i, tail[0:40])
			} else {
				log.Printf("mode: %d offset: %d tail: `%s`\n", mode, i, tail)
			}
		}
		switch mode {
		case COLL, ENV: // 0 == opening Envelopes collection, 1 == start of Envelope expected

			// find the delimiter for start of collection of envelopes or individual envelope
			n := strings.IndexAny(tail, "[{")
			if n < 0 {
				// opening delimiter not found, check for final closing delimiter
				final := strings.TrimSpace(tail)
				if !(final == "]" || final == "}") {
					// closing collection delimiter not found, malformed input
					panic(fmt.Sprintf("Unexpected tail `%s` at offset %d", final, i))
				}
				// closing collection delimiter found, shift mode to EOF
				mode = EOF
				break
			}

			// if COLL, shift mode to ENV; if ENV, shift to WIDTH
			mode += 1
			// skip to the position of the delimiter we found in the input
			i += n

		case WIDTH: // 2 == width expected

			// find the delimiter between width and height
			n := strings.IndexAny(tail, ",")
			if n < 0 {
				// delimiter not found, malformed input
				panic(fmt.Sprintf("Missing `,` delimiter in input at offset %d", i))
			}
			// convert the tail up to the delimiter to a #
			width, err = strconv.Atoi(strings.TrimSpace(tail[0:n]))
			if err != nil {
				// conversion to # failed, malformed input
				panic(err)
			}
			// shift mode to HEIGHT
			mode++
			// skip to the position of the delimiter we found in the input
			i += n

		case HEIGHT: // 3 == height expected

			// find the end-of-envelope delimiter
			n := strings.IndexAny(tail, "}]")
			if n < 0 {
				// delimiter not found, malformed input
				panic(fmt.Sprintf("Missing closing bracket/brace in input at offset %d", i))
			}

			// convert the tail up to the delimiter to a #
			height, err = strconv.Atoi(strings.TrimSpace(tail[0:n]))
			if err != nil {
				// conversion to # failed, malformed input
				panic(err)
			}

			if VERBOSE {
				fmt.Printf("Storing [ %d, %d ]\n", width, height)
			}

			// invoke the callback to process the newly found envelope
			putfx(width, height)

			// shift mode back to looking for start of next envelope;
			// note we don't bother looking for the comma between envelopes
			mode = ENV
			// skip to the position of the delimiter we found in the input
			i += n
		}

		if mode == EOF {
			// we reached the end of input
			break
		}

	}

}
*****/
