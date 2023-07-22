package main

import (
	"testing"
)

func Test_getCharCount(t *testing.T) {
	data := []byte("hello world")
	count := countBytes(&data)

	if count != 11 {
		t.Errorf("getBytes() = %v, want %v", count, 11)
	}
}

func Test_getNewlineCount(t *testing.T) {
	data := []byte("hello\nworld")
	count := countNewlines(&data)

	if count != 1 {
		t.Errorf("getNewlineCount() = %v, want %v", count, 1)
	}

	data = []byte("hello\nworld\n")
	count = countNewlines(&data)

	if count != 2 {
		t.Errorf("getNewlineCount() = %v, want %v", count, 2)
	}
}

func Test_getWords(t *testing.T) {
	data := []byte("hello world")
	count := countWords(&data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}

	data = []byte("hello world\n")
	count = countWords(&data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}

	data = []byte("hello world\n\n")
	count = countWords(&data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}
}
