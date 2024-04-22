package main

import (
	"JeffR/lib"
	librn "JeffR/libsln" // package "librn"
	"flag"
	"fmt"
	"os"
)

func main() {

	// -v for troubleshooting only, defaults to disabled
	verbosePtr := flag.Bool("v", lib.VERBOSE, "specifies whether to emit troubleshooting output")
	flag.Parse()

	if verbosePtr != nil {
		lib.VERBOSE = *verbosePtr
	}

	// get the max nested envelopes from stdin
	envelopes := librn.GetNestedEnvelopes(os.Stdin)

	if lib.VERBOSE {
		fmt.Println("Max Nested Envelopes:")
		fmt.Print("[ ")
		for i, envelope := range envelopes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("[ %d, %d ]", librn.EnvWidth(envelope), librn.EnvHeight(envelope))

		}
		fmt.Println(" ]")
		fmt.Printf("Max Nested Envelope Count: %d", len(envelopes))
	} else {
		// standard output is just the # of nested envelopes
		fmt.Printf("%d", len(envelopes))
	}
}
