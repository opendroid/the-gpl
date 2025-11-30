package chapter7

import (
	"fmt"
	"strings"

	"github.com/opendroid/the-gpl/chapter2/tempConv"
	"github.com/spf13/cobra"
)

// NewDegreesCmd creates the degrees command
// eg: the-gpl degrees --c=-20 --f=-20 --k=-20
func NewDegreesCmd() *cobra.Command {
	var c, f, k float64

	cmd := &cobra.Command{
		Use:   "degrees",
		Short: "Temperature conversion utilities",
		Long:  `Convert temperatures between Celsius, Fahrenheit, and Kelvin.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Note: The original implementation used custom flags that returned *tempConv.Celsius etc.
			// Here we simplify by taking float values and converting them.
			// However, the original code used CelsiusFlag which might have had side effects or specific parsing.
			// Looking at the original code:
			// celsius = CelsiusFlag("c", -20, "degrees convert to Celsius", degrees.set)
			// It seems it was just a helper to register the flag.
			// But wait, the original code printed *celsius, *fahrenheit, *kelvin.
			// If the user didn't provide -c, it used the default.
			// The original implementation printed ALL of them.

			// To match original behavior, we need to convert the float values to their respective types.
			cVal := tempConv.Celsius(c)
			fVal := tempConv.Fahrenheit(f)
			kVal := tempConv.Kelvin(k)

			fmt.Printf("Celsius = %s\n", cVal)
			fmt.Printf("Fahrenheit = %s\n", fVal)
			fmt.Printf("Kelvin = %s\n", kVal)
		},
	}

	cmd.Flags().Float64Var(&c, "c", -20, "degrees convert to Celsius")
	cmd.Flags().Float64Var(&f, "f", -20, "degrees convert to Fahrenheit")
	cmd.Flags().Float64Var(&k, "k", -20, "degrees convert to Kelvin")

	return cmd
}

// NewCountCmd creates the count command
// eg: the-gpl count --text="some text"
func NewCountCmd() *cobra.Command {
	var text string

	cmd := &cobra.Command{
		Use:   "count",
		Short: "Count characters, words, and lines",
		Long:  `Counts characters, words and lines in an input text string.`,
		Run: func(cmd *cobra.Command, args []string) {
			c, w, l := CountCharsWordsLines(strings.NewReader(text))
			fmt.Printf("Text:%s\nCharacters=%d, Words=%d, lines=%d\n", text, c, w, l)
		},
	}

	cmd.Flags().StringVar(&text, "text", "The quick brown 狐狸\n jumps over the lazy 狗", `count characters, words and lines in -text=""`)

	return cmd
}
