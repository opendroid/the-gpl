// Package main is the-gpl entry point for all chapters.
package main

import (
	"github.com/opendroid/the-gpl/chapter1/bot"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"github.com/opendroid/the-gpl/chapter1/lissajous"
	"github.com/opendroid/the-gpl/chapter1/livecaption"
	"github.com/opendroid/the-gpl/chapter1/mas"
	"github.com/opendroid/the-gpl/chapter2/bitsCount"
	"github.com/opendroid/the-gpl/chapter2/tempConv"
	"github.com/opendroid/the-gpl/chapter5"
	"github.com/opendroid/the-gpl/chapter7"
	"github.com/opendroid/the-gpl/serve/shell"
	"github.com/opendroid/the-gpl/serve/web"
	"os"
)

// main initializes all  modules. then enables commands eg:
//   go run main.go server -port=8081
func main() {
	// Init modules - Sets up CLIDegrees Interface
	bitsCount.InitCli()
	bot.InitCli()
	channels.InitCli()
	chapter5.InitCli()
	chapter7.InitCli()
	lissajous.InitCli()
	livecaption.InitCli()
	mas.InitCli()
	tempConv.InitCli()
	web.InitCli()


	// Execute commands
	shell.ExecCLICmd(os.Args[:])
}



