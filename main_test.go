package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

// Test commands: run these in dir "the-gpl"
//  go test -v ./...
//  go test ./chapter1/channels/channels.go ./chapter1/channels/channels_test.go -v
//  go test ./chapter1/lissajous/lissajous.go ./chapter1/lissajous/lissajous_test.go -v
//  go test ./chapter1/structs/structs.go ./chapter1/structs/structs_test.go -v
//
// Individual tests run in specific directory:
//  cd ./channels
//  go test -run TestFetch -v
//  cd ./lissajous
//  go test -run TestLissajous -v
