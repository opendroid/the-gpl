package livecaption

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewSTTCmd creates the stt command
// eg: the-gpl stt --port=9999
func NewSTTCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "stt",
		Short: "Speech to Text",
		Long: `Convert Speech on a specific RTP port to text.
First run this command:
ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ar 48000 -f s16le udp://localhost:9999`,
		Run: func(cmd *cobra.Command, args []string) {
			p := fmt.Sprintf(":%d", port)
			StreamRTPPort(p, os.Stdout)
		},
	}

	cmd.Flags().IntVar(&port, "port", defaultRTPPort, "RTP Port")

	return cmd
}
