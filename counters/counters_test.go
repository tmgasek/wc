package counters

import (
	"testing"
)

func Test_getCharCount(t *testing.T) {
	data := []byte("hello world")
	count := CountBytes(data)

	if count != 11 {
		t.Errorf("getBytes() = %v, want %v", count, 11)
	}
}

func Test_getNewlineCount(t *testing.T) {
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

func Test_getWords(t *testing.T) {
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

func Test_getChars(t *testing.T) {
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
