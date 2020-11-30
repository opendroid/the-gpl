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
//   eg: the-gpl server -port=8888 # Starts http on port
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
	prefix := fmt.Sprintf("http://localhost:%d/", port)
	logger.Log.Println("Starting server, sample URLs:")
	logger.Log.Println(prefix)
	logger.Log.Printf("%swho\n", prefix)
	logger.Log.Printf("%sgraph\n", prefix)
	logger.Log.Printf("%segg\n", prefix)
	logger.Log.Printf("%ssinc\n", prefix)
	logger.Log.Printf("%ssearch?q=%%22%s%%22&imeout=3s\n", prefix, "The%20Go%20Programming%20Language")
	logger.Log.Printf("%smandel\n", prefix)
	logger.Log.Printf("%smandelbw\n", prefix)
	logger.Log.Printf("%sincr\n", prefix)
	logger.Log.Printf("%scounter\n", prefix)
	logger.Log.Println(prefix + `echo?q=%22You%20can%20echo%20this%20back%22`)
	logger.Log.Println(prefix + `post?q=%22Go%22&r=%22Vote%22&year=%222020%22&q=%22Lang%22&`)
	Start(port)
}
