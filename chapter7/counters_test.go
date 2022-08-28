package chapter7

import (
	"fmt"
	"strings"
	"testing"
)

// testData is set that contains values to be tested and expectation of various lengths
var testData = []struct {
	data                 string
	length, words, lines int
}{
	{data: "Hello 2020", length: len("Hello 2020"), words: 2, lines: 1},
	{data: "The quick brown fox\n jumps over the lazy dog", length: len("The quick brown fox\n jumps over the lazy dog"), words: 9, lines: 2},
	{data: "The quick brown 狐狸\n jumps over the lazy 狗", length: len("The quick brown 狐狸\n jumps over the lazy 狗"), words: 9, lines: 2},
}

// TestByteCounter_Write tests ByteCounter writer interface
//
//	cd chapter7
//	go test -run TestByteCounter_Write -v
func TestByteCounter_Write(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, len: %d", w.data, len(w.data))
		t.Run(title, func(t *testing.T) {
			var c ByteCounter
			_, _ = fmt.Fprintf(&c, "%s", w.data)
			if int(c) != w.length {
				t.Logf("Error: Expected bytes %d, Got: %d", w.length, c)
				t.Fail()
			}
		})
	}
}

// TestWordCounter_Write tests WordCounter writer interface
//
//	cd chapter7
//	go test -run TestWordCounter_Write -v
func TestWordCounter_Write(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, words: %d", w.data, w.words)
		t.Run(title, func(t *testing.T) {
			var c WordCounter
			_, _ = fmt.Fprintf(&c, "%s", w.data)
			if int(c) != w.words {
				t.Logf("Error: Expected words %d, Got: %d", w.words, c)
				t.Fail()
			}
		})
	}
}

// TestLineCounter_Write tests LineCounter writer interface
//
//	cd chapter7
//	go test -run TestLineCounter_Write -v
func TestLineCounter_Write(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, lines: %d", w.data, w.lines)
		t.Run(title, func(t *testing.T) {
			var c LineCounter
			_, _ = fmt.Fprintf(&c, "%s", w.data)
			if int(c) != w.lines {
				t.Logf("Error: Expected lines %d, Got: %d", w.lines, c)
				t.Fail()
			}
		})
	}
}

// TestCountingWriter tests CountingByteWriter writer interface
//
//	cd chapter7
//	go test -run TestCountingWriter -v
func TestCountingWriter(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, len: %d", w.data, len(w.data))
		t.Run(title, func(t *testing.T) {
			var chars ByteCounter
			cw, nChars := CountingWriter(&chars) //

			_, _ = fmt.Fprintf(cw, "%s", w.data)
			if *nChars != int64(w.length) { // Count of CountingWriter
				t.Logf("Error: Expected bytes %d, Got: %d", w.length, *nChars)
				t.Fail()
			}
			if int(chars) != w.length { // Count of ByteCounter Writer
				t.Logf("Error: Expected bytes %d, Got: %d", w.length, chars)
				t.Fail()
			}
		})
	}
}

// TestNewBroadcastWriters tests broadcasting to groups of writer interfaces
//
//	cd chapter7
//	go test -run TestNewBroadcastWriters -v
func TestNewBroadcastWriters(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, len: %d, words: %d, lines: %d", w.data, len(w.data), w.words, w.lines)
		t.Run(title, func(t *testing.T) {
			var cw ByteCounter // character writer
			var ww WordCounter // word writer
			var lw LineCounter // line writer

			radio := NewBroadcastWriters(&cw, &ww, &lw)
			_, _ = fmt.Fprintf(radio, "%s", w.data)
			if int(cw) != w.length {
				t.Logf("Error: Expected words %d, Got: %d", w.length, cw)
				t.Fail()
			}
			if int(ww) != w.words {
				t.Logf("Error: Expected words %d, Got: %d", w.words, ww)
				t.Fail()
			}
			if int(lw) != w.lines {
				t.Logf("Error: Expected lines %d, Got: %d", w.lines, lw)
				t.Fail()
			}
		})
	}
}

// TestCountCharsWordsLines tests chaining of writer interface
//
//	cd chapter7
//	go test -run TestCountCharsWordsLines -v
func TestCountCharsWordsLines(t *testing.T) {
	for _, w := range testData {
		title := fmt.Sprintf("Bytes: %s, len: %d, words: %d, lines: %d", w.data, len(w.data), w.words, w.lines)
		t.Run(title, func(t *testing.T) {
			// Get the character count, word count and line count
			cc, wc, lc := CountCharsWordsLines(strings.NewReader(w.data))
			if cc != w.length {
				t.Logf("Error: Expected words %d, Got: %d", w.length, cc)
				t.Fail()
			}
			if wc != w.words {
				t.Logf("Error: Expected words %d, Got: %d", w.words, wc)
				t.Fail()
			}
			if lc != w.lines {
				t.Logf("Error: Expected lines %d, Got: %d", w.lines, lc)
				t.Fail()
			}
		})
	}
}
