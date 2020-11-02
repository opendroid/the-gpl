package chapter8

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve/shell"
	"strconv"
	"strings"
)

// CliServer wrapper for *flag.FlagSet
type CliServer struct {
	set *flag.FlagSet
}

// CliClient wrapper for *flag.FlagSet
type CliClient struct {
	set *flag.FlagSet
}

// cmdServer allows to refer call send this module the CliServer argument
var cmdServer CliServer
var serverPort *string // Flag that stores value for -type="clock:port"

// cmdClient allows to refer call send this module the CliClient argument
var cmdClient CliClient
var clientPort *string // Flag that stores value for -type="clock:port"

// InitCli initialize the services APIs
func InitCli() {
	cmdServer.set = flag.NewFlagSet("service", flag.ContinueOnError)
	serverPort = cmdServer.set.String("sp", "clock:9999", "server-type:port eg \"clock:9999\"")
	shell.Add("service", cmdServer)

	cmdClient.set = flag.NewFlagSet("client", flag.ContinueOnError)
	clientPort = cmdClient.set.String("cp", "clock:9999", "client-type:port eg \"clock:9999\"")
	shell.Add("client", cmdClient)
}

// Implement CLI for server side
// ExecCmd executes a specific command
func (m CliServer) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: service Parse Error %s\n", err.Error())
		return
	}
	params := strings.Split(*serverPort, ":")
	if len(params) < 2 {
		fmt.Printf("Invalid input: %q. Expect -sp=\"clock:9999\"", *serverPort)
		return
	}
 	service := params[0]
	port, _ := strconv.Atoi(params[1])
	switch service {
	case "clock":
		fmt.Printf("Started %q service on %d\n", service, port)
		ClockServer(port)
	default:
		fmt.Printf("service %s not implemented\n", service)
	}
}

func (m CliServer) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl service -sp=\"clock:9999\" # Starts a clock service on port 9999.")
	m.set.PrintDefaults()
}

// Implement CLI for client side
// ExecCmd executes a client CLI command with syntax
//	the-gpl service -cp="clock:9999"
func (m CliClient) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: client Parse Error %s\n", err.Error())
		return
	}
	params := strings.Split(*clientPort, ":")
	if len(params) < 2 {
		fmt.Printf("Invalid input: %q. Expect -cp=\"clock:9999\"", *clientPort)
		return
	}
	client := params[0]
	port, _ := strconv.Atoi(params[1])
	switch client {
	case "clock":
		fmt.Printf("Started %q client on %d\n", client, port)
		ClockClient(port)
	default:
		fmt.Printf("client %s not implemented\n", client)
	}
}

func (m CliClient) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl service -cp=\"clock:9999\" # Listens to a clock service on port 9999.")
	m.set.PrintDefaults()
}
