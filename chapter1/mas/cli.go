package mas

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewMasCmd creates the mas command
// eg: the-gpl mas --fn=array # Tests the array example
//
//	the-gpl mas --fn=comp --n1=123 --n2=345 # compares n1 and n2 and computes n1 - n2
func NewMasCmd() *cobra.Command {
	var callMethod string
	var n1, n2 int

	cmd := &cobra.Command{
		Use:   "mas",
		Short: "Maps Arrays Slices examples",
		Long:  `Run various examples for Maps, Arrays, and Slices.`,
		Run: func(cmd *cobra.Command, args []string) {
			switch callMethod {
			case "array":
				IterateOverArray()
			case "comp":
				compResult, diff := CompareNumbers(n1, n2)
				fmt.Printf("mas:CompareNumbers: ints: %d == %d, => %t, differance: %d\n", n1, n2, compResult, diff)
			case "slice":
				AddToSlices()
			}
		},
	}

	cmd.Flags().StringVar(&callMethod, "fn", "array", "[array comp slice]")
	cmd.Flags().IntVar(&n1, "n1", 123, "first number")
	cmd.Flags().IntVar(&n2, "n2", -46, "second number")

	return cmd
}
