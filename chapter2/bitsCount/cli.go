package bitsCount

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve/shell"
)

// Command line help func

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// bitCountCmd allows to refer call send this module the CLI argument
var bitCountCmd CLI
var bits *uint64 // flag for -n=#bits count

// InitCli for the "bits" command
//   eg: the-gpl bits -n=0x1234BAD0 # counts number of 1 bits in n
func InitCli() {
	bitCountCmd.set = flag.NewFlagSet("bits", flag.ContinueOnError)
	bits = bitCountCmd.set.Uint64("n", 0xBAD0FACEC0FFEE, "A 64-bit int")
	shell.Add("bits", bitCountCmd)
}

// ExecCmd run bit count from CLI
func (b CLI) ExecCmd(args []string) {
	err := b.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecBitsCountCmd: Bit count Parse Error %s\n", err.Error())
		return
	}
	fmt.Printf("\tThere are %d one bits in 0X%016X\n", BitCountEachOne(*bits), *bits)
}

// DisplayHelp prints help on command line for bits module
func (b CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl bits. Number of 1 bits in n")
	b.set.PrintDefaults()
}