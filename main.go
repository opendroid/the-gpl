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
	"github.com/opendroid/the-gpl/chapter8"
	"github.com/opendroid/the-gpl/cmd"
	"github.com/opendroid/the-gpl/serve/web"
)

// main initializes all modules and executes the root command.
func main() {
	rootCmd := cmd.GetRootCmd()

	// Chapter 1
	rootCmd.AddCommand(bot.NewBotCmd())
	rootCmd.AddCommand(channels.NewFetchCmd())
	rootCmd.AddCommand(lissajous.NewLissajousCmd())
	rootCmd.AddCommand(livecaption.NewSTTCmd())
	rootCmd.AddCommand(mas.NewMasCmd())

	// Chapter 2
	rootCmd.AddCommand(bitsCount.NewBitsCmd())
	rootCmd.AddCommand(tempConv.NewTempCmd())

	// Chapter 5
	rootCmd.AddCommand(chapter5.NewParseCmd())

	// Chapter 7
	rootCmd.AddCommand(chapter7.NewDegreesCmd())
	rootCmd.AddCommand(chapter7.NewCountCmd())

	// Chapter 8
	rootCmd.AddCommand(chapter8.NewServiceCmd())
	rootCmd.AddCommand(chapter8.NewClientCmd())
	rootCmd.AddCommand(chapter8.NewDuCmd())

	// Serve
	rootCmd.AddCommand(web.NewServerCmd())

	// Execute
	cmd.Execute()
}
