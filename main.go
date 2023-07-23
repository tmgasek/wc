package main

import (
	"fmt"
	"io"
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
	pflag.BoolVarP(&bytesFlag, "chars", "c", false, "count bytes")
	pflag.BoolVarP(&charsFlag, "bytes", "m", false, "count chars")
	pflag.Parse()

	// if no flags, default to -lwc
	if !linesFlag && !wordsFlag && !charsFlag && !bytesFlag {
		linesFlag = true
		wordsFlag = true
		bytesFlag = true
	}

	args := pflag.Args()

	var data []byte
	var err error
	var file string

	if len(args) == 0 {
		fmt.Println("Usage: [OPTIONS] FILE")
		fmt.Println("Count lines, words, characters, and bytes.")
		fmt.Println("If FILE is - read from standard input.")
		fmt.Println("Options:")
		pflag.PrintDefaults()
		os.Exit(1)
	}

	// if no file or file is -, read from stdin
	if args[0] == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		// read in file
		file = args[0]
		data, err = os.ReadFile(file)
	}

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
