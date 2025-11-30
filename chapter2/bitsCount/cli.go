package bitsCount

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewBitsCmd creates the bits command
// eg: the-gpl bits --n=0x1234BAD0 # counts number of 1 bits in n
func NewBitsCmd() *cobra.Command {
	var bits uint64

	cmd := &cobra.Command{
		Use:   "bits",
		Short: "Count bits",
		Long:  `Counts number of 1 bits in a 64-bit integer.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\tThere are %d one bits in 0X%016X\n", BitCountEachOne(bits), bits)
		},
	}

	cmd.Flags().Uint64Var(&bits, "n", 0xBAD0FACEC0FFEE, "A 64-bit int")

	return cmd
}
