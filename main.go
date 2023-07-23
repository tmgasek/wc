package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

// handle flags -lwc
// -c, --bytes            print the byte counts
// -m, --chars            print the character counts
// -l, --lines            print the newline counts
//
//	--files0-from=F    read input from the files specified by
//	                     NUL-terminated names in file F;
//	                     If F is - then read names from standard input
//
// -L, --max-line-length  print the maximum display width
// -w, --words            print the word counts
// by default, it does newline, word, byte
// order: newline, word, char, byte, max line length

type Operation func(data []byte) int

var operations = map[string]Operation{
	"l": countNewlines,
	"w": countWords,
	"c": countBytes,
	"m": countChars,
}

func main() {
	var linesFlag bool
	var wordsFlag bool
	var charsFlag bool
	var bytesFlag bool
	pflag.BoolVarP(&linesFlag, "lines", "l", false, "count lines")
	pflag.BoolVarP(&wordsFlag, "words", "w", false, "count words")
	pflag.BoolVarP(&charsFlag, "chars", "c", false, "count bytes")
	pflag.BoolVarP(&bytesFlag, "bytes", "m", false, "count chars")
	pflag.Parse()

	args := pflag.Args()

	if (len(args) == 0) || (len(args) > 2) {
		log.Fatal("Usage: wc [-clmw] [file]")
	}

	if len(args) != 1 {
		log.Fatal("expected one file")

	}

	file := args[0]

	// read in file
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	flags := map[string]bool{
		"l": linesFlag,
		"w": wordsFlag,
		"c": bytesFlag,
		"m": charsFlag,
	}

	var output string

	for flag, operation := range operations {
		if flags[flag] {
			output += fmt.Sprintf("%v \t", operation(data))
		}
	}

	fmt.Printf("%v\n", output+file)
}
