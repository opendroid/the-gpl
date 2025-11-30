package web

import (
	"fmt"

	"github.com/opendroid/the-gpl/logger"
	"github.com/spf13/cobra"
)

// NewServerCmd creates the server command
// eg: the-gpl server --port=8888 # Starts http on port
func NewServerCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start HTTP server",
		Long:  `Starts the HTTP server on the specified port.`,
		Run: func(cmd *cobra.Command, args []string) {
			startServer(port)
		},
	}

	cmd.Flags().IntVar(&port, "port", 8080, "Port number eg: 8080")

	return cmd
}

// startServer starts the server on port
func startServer(port int) {
	prefix := fmt.Sprintf("http://localhost:%d", port)
	logger.Log.Println("Starting server, sample URLs:")
	logger.Log.Println(prefix)
	logger.Log.Printf("%s/who\n", prefix)
	logger.Log.Printf("%s/graph\n", prefix)
	logger.Log.Printf("%s/egg\n", prefix)
	logger.Log.Printf("%s/sinc\n", prefix)
	logger.Log.Printf("%s/search?q=%%22%s%%22&Timeout=3s\n", prefix, "The%20Go%20Programming%20Language")
	logger.Log.Printf("%s/mandel\n", prefix)
	logger.Log.Printf("%s/mandelbw\n", prefix)
	logger.Log.Printf("%s/incr\n", prefix)
	logger.Log.Printf("%s/counter\n", prefix)
	Start(port)
}
