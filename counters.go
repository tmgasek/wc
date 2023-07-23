package main

import (
	"unicode"
	"unicode/utf8"
)

func countBytes(data []byte) int {
	return len(data)
}

func countNewlines(data []byte) int {
	count := 0
	for _, b := range data {
		if b == '\n' {
			count++
		}
	}

	return count
}

func countWords(data []byte) int {
	count := 0
	inWord := false
	for _, b := range data {
		if unicode.IsSpace(rune(b)) {
			if inWord {
				count++
				inWord = false
			}
		} else {
			inWord = true
		}

	}

	if inWord {
		count++
	}

	return count
}

func countChars(data []byte) int {
	return utf8.RuneCount(data)
}
