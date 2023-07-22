package main

import (
	"flag"
	"fmt"
	"os"
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

type Operation func(data *[]byte) int

var operations = map[string]Operation{
	"l": countNewlines,
	"w": countWords,
	"c": countBytes,
	"m": countChars,
}

func main() {
	linesFlag := flag.Bool("l", false, "print the newline counts")
	wordsFlag := flag.Bool("w", false, "print the word counts")
	bytesFlag := flag.Bool("c", false, "print the byte counts")
	charsFlag := flag.Bool("m", false, "print the character counts")

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		panic("expected one file")
	}

	file := args[0]

	// read in file
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	flags := map[string]bool{
		"l": *linesFlag,
		"w": *wordsFlag,
		"c": *bytesFlag,
		"m": *charsFlag,
	}

	var output string

	for flag, operation := range operations {
		if flags[flag] {
			output += fmt.Sprintf("%v %s \t", operation(&data), flag)
		}
	}

	fmt.Printf("%v\n", output)
}
