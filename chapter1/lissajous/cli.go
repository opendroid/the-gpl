package lissajous

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// NewLissajousCmd creates the lissajous command
// eg: the-gpl lissajous --cycles=2 --size=1025 --frames=10 --file=~/Downloads.gif
func NewLissajousCmd() *cobra.Command {
	var cycles, size, frames int
	var outFileName string

	cmd := &cobra.Command{
		Use:   "lissajous",
		Short: "Generate Lissajous figures",
		Long:  `Generates a GIF animation of Lissajous figures.`,
		Run: func(cmd *cobra.Command, args []string) {
			f, err := os.Create(outFileName)
			if err != nil {
				fmt.Printf("ExecCmd: Write to file %s: %s\n", outFileName, err.Error())
				return
			}
			saveSquareImage(cycles, size, frames, f)
		},
	}

	cmd.Flags().IntVar(&cycles, "cycles", 2, "number of cycles")
	cmd.Flags().IntVar(&size, "size", 512, "size of square image")
	cmd.Flags().IntVar(&frames, "frames", 10, "number of frames")
	cmd.Flags().StringVar(&outFileName, "file", "lis.gif", "name of output file")

	return cmd
}

// saveImageSquare tests lissajous package
func saveSquareImage(cycles, size, frames int, w io.Writer) {
	config := Config{
		Cycles:     cycles,
		Resolution: 0.000001,
		Size:       size,
		NFrames:    frames,
		DelayMS:    10,
	}
	Lissajous(w, config)
}
