// Package gplCLI provides interface for displaying help and executing
//  command line. A module wishing to participate in CLI need to conform
//	to CmdHandlers interface and its ExecCmd method will be called when command
//	matches.
package gplCLI

import (
	"fmt"
	"os"
	"sort"
)

// CmdHandlers interface to invoke command or display help.
type CmdHandlers interface{
	ExecCmd([]string)
	DisplayHelp()
}

// flagHandlerMap is map of first "cmd" and it Set handlers
//  E.G. The "lissajous" interface{} entry associated CLI for module lissajous
var flagHandlerMap = map[string]CmdHandlers{}

// Add registers interface associated with command for a module so it can ve invoked
func Add(cmd string, handlers CmdHandlers) {
	flagHandlerMap[cmd] = handlers
}

// ExecCLICmd dispatches the command to a module registered via Add
func ExecCLICmd(args []string) {
	// Flag in Main
	if len(args) < 2 {
		printArgsHelp() // Print help
		os.Exit(1)
	}

	// Invoke command line handlers
	if v, ok := flagHandlerMap[args[1]]; ok {
		if v != nil {
			v.ExecCmd(args[2:])
		}
	}
}

// printArgsHelp Print help for all arguments in sorted order of commands.
func printArgsHelp() {
	fmt.Print("\nUsage: the-gpl ")
	fmt.Print("[")
	// Sort keys first.
	var keys []string
	for k := range flagHandlerMap { // Fetch all keys
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort keys in place
	addSpace := false // Print keys
	for _, k := range keys {
		if addSpace {
			fmt.Print(" ")
		}
		addSpace = true
		fmt.Printf("%s", k)
	}
	fmt.Print("]\n")
	// Print help in sorted order as well
	for _, k := range keys {
		if v, ok := flagHandlerMap[k]; ok {
			if v != nil {
				v.DisplayHelp()
			}
		}
	}
	fmt.Println("")
}
