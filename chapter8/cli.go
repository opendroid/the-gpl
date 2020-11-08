package chapter8

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve/shell"
	"strconv"
	"strings"
)

// CliService wrapper for *flag.FlagSet
type CliService struct {
	set *flag.FlagSet
}

// CliClient wrapper for *flag.FlagSet
type CliClient struct {
	set *flag.FlagSet
}

// CliDu wrapper for *flag.FlagSet for "du" - disk usage command
type CliDu struct {
	set *flag.FlagSet
}

// cmdServer allows to refer call send this module the CliService argument
var cmdServer CliService
var serverPort *string // Flag that stores value for -sp="clock:port"

// cmdClient allows to refer call send this module the CliClient argument
var cmdClient CliClient
var clientPort *string // Flag that stores value for -cp="clock:port"

// cmdDu allows to refer call send this module the CliDu argument
var cmdDu CliDu
var dir *string // Flag that stores value for -dir="./"
var verbose *bool // Flag that stores value for -dir="./"


// InitCli initialize the services APIs
func InitCli() {
	cmdServer.set = flag.NewFlagSet("service", flag.ContinueOnError)
	serverPort = cmdServer.set.String("sp", "clock:9999", "service-type:port eg: \"clock:9999\" or \"reverb:9998\" or \"ftp:9997\"  or \"chat:9998\"")
	shell.Add("service", cmdServer)

	cmdClient.set = flag.NewFlagSet("client", flag.ContinueOnError)
	clientPort = cmdClient.set.String("cp", "clock:9999", "client-type:port eg \"clock:9999\" or \"reverb:9998\" or \"chat:9998\"")
	shell.Add("client", cmdClient)

	cmdDu.set = flag.NewFlagSet("du", flag.ContinueOnError)
	dir = cmdDu.set.String("dir", ".", "du -dir:\".\"")
	verbose = cmdDu.set.Bool("v", false, "du -v:false")
	shell.Add("du", cmdDu)
}

// Implement CLI for server side
// ExecCmd executes a specific service command
func (m CliService) ExecCmd(args []string) {
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
		fmt.Printf("Started %q service on port %d\n", service, port)
		ClockServer(port)
	case "reverb":
		fmt.Printf("Started %q service on port %d\n", service, port)
		ReverbServer(port)
	case "ftp":
		fmt.Printf("Started %q service on port %d\n", service, port)
		FTPServer(port)
	case "chat":
		fmt.Printf("Started %q service on port %d\n", service, port)
		ChatService(port)
	default:
		fmt.Printf("service %s not implemented\n", service)
	}
}

// DisplayHelp for the services
func (m CliService) DisplayHelp() {
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
	case "reverb":
		fmt.Printf("Started %q client on %d\n", client, port)
		ReverbClient(port)
	default:
		fmt.Printf("client %s not implemented\n", client)
	}
}

// DisplayHelp for the service modules
func (m CliClient) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl service -cp=\"clock:9999\" # Listens to a clock service on port 9999.")

	m.set.PrintDefaults()
}

// Implement CLI for du - disc Usage
// ExecCmd executes a client CLI command with syntax
//	the-gpl service -cp="clock:9999"
func (m CliDu) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: du Error %s\n", err.Error())
		return
	}
	_ = DU(*dir,  *verbose)
}

// DisplayHelp for the service modules
func (m CliDu) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl du -dir=\".\" # Print directory usage in 'dir'")
	m.set.PrintDefaults()
}