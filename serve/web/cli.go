package web

import (
	"fmt"
	"log/slog"

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
	slog.Info("Starting server, sample URLs:")
	slog.Info(prefix)
	slog.Info(fmt.Sprintf("%s/who", prefix))
	slog.Info(fmt.Sprintf("%s/graph", prefix))
	slog.Info(fmt.Sprintf("%s/egg", prefix))
	slog.Info(fmt.Sprintf("%s/sinc", prefix))
	slog.Info(fmt.Sprintf("%s/search?q=%%22%s%%22&Timeout=3s", prefix, "The%20Go%20Programming%20Language"))
	slog.Info(fmt.Sprintf("%s/mandel", prefix))
	slog.Info(fmt.Sprintf("%s/mandelbw", prefix))
	slog.Info(fmt.Sprintf("%s/incr", prefix))
	slog.Info(fmt.Sprintf("%s/counter", prefix))
	Start(port)
}
