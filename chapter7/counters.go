// Package chapter7, Interfaces, defines bytes, words and line counter writer interfaces that can be used as part of fmt methods.
package chapter7

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
)

// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.
//  You will find bufio.ScanWords useful.

// ByteCounter writer interface counts number of bytes written to it
type ByteCounter int

// WordCounter writer interface counts number of words written to it
type WordCounter int

// LineCounter writer interface counts number of lines written to it
type LineCounter int

// Write to a ByteCounter interface
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// Write to a WordCounter interface
func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

// Write to a LineCounter interface
func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

// Exercise 7.2 given an io.Writer, returns a new Writer
//   that wraps the original, and a pointer to an int64 variable that at any moment contains
//   the number of bytes written to the newWriter.

// CountWriter intercepts a io.Writer and counts numbers of characters written to it
type CountWriter struct {
	count int64
	w     io.Writer
}

// Write define the counting Write interface interface
func (c *CountWriter) Write(p []byte) (int, error) {
	c.count += int64(len(p))
	return c.w.Write(p) // Write to original writer
}

// CountingWriter intercept a writer and count bytes written to it
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{count: 0, w: w}
	return &cw, &cw.count
}

// Exercise: Create a broadcast writer list so we can call
//	ByteCounter, WordCounter and LineCounter on one list

// BroadcastWriters sequentially writes to writers in linked list
type BroadcastWriters struct {
	writers *list.List
}

// Write calls the Writers in sequence
func (c *BroadcastWriters) Write(p []byte) (int, error) {
	var data int
	var err error
	for w := c.writers.Front(); w != nil; w = w.Next() {
		if writer, ok := w.Value.(io.Writer); ok {
			if writer != nil { // ensure type and value are non-nil
				data, err = writer.Write(p)
			}
		}
	}
	return data, err // return last data
}

// NewBroadcastWriters creates a broadcast list of Writers
func NewBroadcastWriters(writers ...io.Writer) *BroadcastWriters {
	cw := &BroadcastWriters{writers: list.New()}
	for _, w := range writers {
		cw.writers.PushBack(w)
	}
	return cw
}

// CountCharsWordsLines reads from a reader and returns count of
//
//	characters words and lines
func CountCharsWordsLines(r io.Reader) (int, int, int) {
	var cw ByteCounter
	var ww WordCounter
	var lw LineCounter

	radio := NewBroadcastWriters(&cw, &ww, &lw)
	_, err := io.Copy(radio, r)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return 0, 0, 0
	}
	return int(cw), int(ww), int(lw)
}
