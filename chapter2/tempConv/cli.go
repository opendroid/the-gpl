package tempConv

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewTempCmd creates the temp command
// eg: the-gpl temp --c=12 # Converts 12°C to °F
func NewTempCmd() *cobra.Command {
	var c, f, k float64

	cmd := &cobra.Command{
		Use:   "temp",
		Short: "Temperature Conversion",
		Long:  `Converts temperatures between Celsius, Fahrenheit, and Kelvin.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\t%s is %s = %s\n", Celsius(c), Celsius(c).ToF(), Celsius(c).ToK())
			fmt.Printf("\t%s is %s = %s\n", Fahrenheit(f), Fahrenheit(f).ToC(), Fahrenheit(f).ToK())
			fmt.Printf("\t%s is %s = %s\n", Kelvin(k), Kelvin(k).ToC(), Kelvin(k).ToF())
		},
	}

	cmd.Flags().Float64Var(&c, "c", float64(FreezingPointC), "°Celsius")
	cmd.Flags().Float64Var(&f, "f", float64(FreezingPointC), "°Fahrenheit")
	cmd.Flags().Float64Var(&k, "k", float64(FreezingPointC), "°Kelvin")

	return cmd
}
