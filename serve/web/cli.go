package web

import (
	"flag"
	"fmt"

	"github.com/opendroid/the-gpl/logger"
	"github.com/opendroid/the-gpl/serve/shell"
)

type CLI struct {
	set *flag.FlagSet
}

// bitCountCmd allows to refer call send this module the CLI argument
var cmd CLI
var port *int

// InitCli for command: the-gpl server
//
//	eg: the-gpl server -port=8888 # Starts http on port
func InitCli() {
	cmd.set = flag.NewFlagSet("server", flag.ContinueOnError)
	port = cmd.set.Int("port", 8080, "Port number eg: 8080")
	shell.Add("server", cmd)
}

// ExecCmd run server from CLI
func (s CLI) ExecCmd(args []string) {
	err := s.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: Server %s\n", err.Error())
		return
	}
	startServer(*port)
}

// DisplayHelp prints help on command line for lissajous module
func (s CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl server. At the -port=")
	s.set.PrintDefaults()
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
