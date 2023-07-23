package counters

import (
	"unicode"
	"unicode/utf8"
)

func CountBytes(data []byte) int {
	return len(data)
}

func CountNewlines(data []byte) int {
	count := 0
	for _, b := range data {
		if b == '\n' {
			count++
		}
	}

	return count
}

func CountWords(data []byte) int {
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

func CountChars(data []byte) int {
	return utf8.RuneCount(data)
}
