package counters

import (
	"strings"
	"testing"
)

func TestGetCharCount(t *testing.T) {
	data := []byte("hello world")
	count := CountBytes(data)

	if count != 11 {
		t.Errorf("getBytes() = %v, want %v", count, 11)
	}
}

func TestGetNewlineCount(t *testing.T) {
	data := []byte("hello\nworld")
	count := CountNewlines(data)

	if count != 1 {
		t.Errorf("getNewlineCount() = %v, want %v", count, 1)
	}

	data = []byte("hello\nworld\n")
	count = CountNewlines(data)

	if count != 2 {
		t.Errorf("getNewlineCount() = %v, want %v", count, 2)
	}
}

func TestGetWords(t *testing.T) {
	data := []byte("hello world")
	count := CountWords(data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}

	data = []byte("hello world\n")
	count = CountWords(data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}

	data = []byte("hello world\n\n")
	count = CountWords(data)

	if count != 2 {
		t.Errorf("getWords() = %v, want %v", count, 2)
	}
}

func TestGetChars(t *testing.T) {
	data := []byte("hello world")
	count := CountChars(data)

	if count != 11 {
		t.Errorf("getChars() = %v, want %v", count, 11)
	}

	data = []byte("hello world\n")
	count = CountChars(data)

	if count != 12 {
		t.Errorf("getChars() = %v, want %v", count, 12)
	}

}

func TestEmptyInput(t *testing.T) {
	data := []byte("")
	if count := CountBytes(data); count != 0 {
		t.Errorf("CountBytes() = %v, want 0", count)
	}
	if count := CountNewlines(data); count != 0 {
		t.Errorf("CountNewlines() = %v, want 0", count)
	}
	if count := CountWords(data); count != 0 {
		t.Errorf("CountWords() = %v, want 0", count)
	}
	if count := CountChars(data); count != 0 {
		t.Errorf("CountChars() = %v, want 0", count)
	}
}

func TestMultipleSpaces(t *testing.T) {
	data := []byte("hello   world")
	if count := CountWords(data); count != 2 {
		t.Errorf("CountWords() = %v, want 2", count)
	}
}

func TestNonAsciiChars(t *testing.T) {
	data := []byte("héllo wörld")
	if count := CountChars(data); count != 11 {
		t.Errorf("CountChars() = %v, want 11", count)
	}
	if count := CountBytes(data); count != 13 {
		t.Errorf("CountBytes() = %v, want 13", count)
	}
}

func TestLargeData(t *testing.T) {
	data := []byte(strings.Repeat("hello world ", 1000))
	if count := CountWords(data); count != 2000 {
		t.Errorf("CountWords() = %v, want 2000", count)
	}
}

func TestBinaryData(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if count := CountBytes(data); count != 10 {
		t.Errorf("CountBytes() = %v, want 10", count)
	}
}

func TestPunctuation(t *testing.T) {
	data := []byte("Hello, world! How are you?")
	if count := CountWords(data); count != 5 {
		t.Errorf("CountWords() = %v, want 5", count)
	}
}
