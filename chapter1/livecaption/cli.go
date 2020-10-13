package livecaption

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve"
	"os"
)

// Cmd allows to refer call send this module the CLI argument

// CLI wrapper for *flag.FlagSet. Implements serve.CmdHandlers CLI interface.
type CLI struct {
	set *flag.FlagSet
}

// cmd embeds the CLI interface so it can be invoked
var cmd CLI
var port *int // flag for stt RTP port

// InitCli for the "the-gpl stt" command
func InitCli() {
	cmd.set = flag.NewFlagSet("stt", flag.ContinueOnError)
	port = cmd.set.Int("port", defaultRTPPort, "RTP Port")
	serve.Add("stt", cmd) // Register with serve module
}

// ExecCmd run stt command dispatched from CLI
func (a CLI) ExecCmd(args []string) {
	err := a.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: Audio Parse Error %s\n", err.Error())
		return
	}
	p := fmt.Sprintf(":%d", *port)
	StreamRTPPort(p, os.Stdout)
}

// DisplayHelp prints help on command line for the livecaption module
func (a CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl stt. First run this command")
	fmt.Println(`ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ar 48000 -f s16le udp://localhost:9999`)
	a.set.PrintDefaults()
}