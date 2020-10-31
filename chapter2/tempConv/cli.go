package tempConv

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve/shell"
)

// Section to setup CLI

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// cmd allows to refer call send this module the CLI argument
var cmd CLI
var c *float64 // flag for Celsius
var f *float64 // flag for Fahrenheit
var k *float64 // flag for Kelvin

// InitCli for the "temp" command
//   eg: the-gpl temp -c=12 # Converts 12°C to °F
func InitCli() {
	cmd.set = flag.NewFlagSet("temp", flag.ContinueOnError)
	c = cmd.set.Float64("c", float64(FreezingPointC), "°Celsius")
	f = cmd.set.Float64("f", float64(FreezingPointC), "°Fahrenheit")
	k = cmd.set.Float64("k", float64(FreezingPointC), "°Kelvin")
	shell.Add("temp", cmd)
}

// ExecCmd run temp conversion command initiated from CLI
func (t CLI) ExecCmd(args []string) {
	err := t.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: TempConv Parse Error %s\n", err.Error())
		return
	}
	fmt.Printf("\t%s is %s = %s\n", Celsius(*c), Celsius(*c).ToF(), Celsius(*c).ToK())
	fmt.Printf("\t%s is %s = %s\n", Fahrenheit(*f), Fahrenheit(*f).ToC(), Fahrenheit(*f).ToK())
	fmt.Printf("\t%s is %s = %s\n", Kelvin(*k), Kelvin(*k).ToC(), Kelvin(*k).ToF())
}

// DisplayHelp prints help on command line for temperature module
func (t CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl temp. Coverts c to f and visa versa")
	t.set.PrintDefaults()
}
