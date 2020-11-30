package chapter7

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/chapter2/tempConv"
	"github.com/opendroid/the-gpl/serve/shell"
	"strings"
)

// CLIDegrees wrapper for *flag.FlagSet for tempConv
type CLIDegrees struct {
	set *flag.FlagSet
}

// CLICounter wrapper for *flag.FlagSet for CountCharsWordsLines
type CLICounter struct {
	set *flag.FlagSet
}

// degrees allows to refer call send this module the CLIDegrees argument
var degrees CLIDegrees
var celsius *tempConv.Celsius
var fahrenheit *tempConv.Fahrenheit
var kelvin *tempConv.Kelvin

// counter counts charcaters, words and lines in a input text string
var counter CLICounter
var text *string

// InitCli initialized CLIDegrees for temperature conversions, eg:
//   the-gpl degrees -c=-20°F -f-20°C -k=-20°K
func InitCli() {
	// Set for temperature utilities
	degrees.set = flag.NewFlagSet("degrees", flag.ContinueOnError)
	celsius = CelsiusFlag("c", -20, "degrees convert to Celsius", degrees.set)
	fahrenheit = FahrenheitFlag("f", -20, "degrees convert to Fahrenheit", degrees.set)
	kelvin = KelvinFlag("k", -20, "degrees convert to Kelvin", degrees.set)
	shell.Add("degrees", degrees)

	// Set for counter
	counter.set = flag.NewFlagSet("count", flag.ContinueOnError)
	text = counter.set.String("text", "The quick brown 狐狸\n jumps over the lazy 狗", `count characters, words and lines in -text=""`)
	shell.Add("count", counter)
}

// ExecCmd executed degrees
func (t CLIDegrees) ExecCmd(args []string) {
	err := t.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: TempConv Parse Error %s\n", err.Error())
		return
	}
	fmt.Printf("Celsius = %s\n", *celsius)
	fmt.Printf("Fahrenheit = %s\n", *fahrenheit)
	fmt.Printf("Kelvin = %s\n", *kelvin)
}

func (t CLIDegrees) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl degrees. Prints degree °C or °F or °K")
	t.set.PrintDefaults()
}

// ExecCmd executed degrees
func (t CLICounter) ExecCmd(args []string) {
	err := t.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: Count Parse Error %s\n", err.Error())
		return
	}
	c, w, l := CountCharsWordsLines(strings.NewReader(*text))
	fmt.Printf("Text:%s\nCharacters=%d, Words=%d, lines=%d\n", *text, c, w, l)
}

func (t CLICounter) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl count. Counts characters, words and lines")
	t.set.PrintDefaults()
}
