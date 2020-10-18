package webserver

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve"
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
	serve.Add("server", cmd)
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
	logger.Println("Starting server, sample URLs:")
	logger.Println(prefix)
	logger.Printf("%sgraph\n", prefix)
	logger.Printf("%segg\n", prefix)
	logger.Printf("%ssinc\n", prefix)
	logger.Printf("%smandel\n", prefix)
	logger.Printf("%smandelbw\n", prefix)
	logger.Printf("%sincr\n", prefix)
	logger.Printf("%scounter\n", prefix)
	logger.Println(prefix+`/echo?q="You can echo this back"`)
	logger.Println(prefix+`post?q="Go"&r="Vote"&year="2020"`)
	Start(port)
}