// Package chapter8, Goroutines and Channels, provides examples of sharing by communicating.
package chapter8

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// NewServiceCmd creates the service command
// eg: the-gpl service --sp="clock:9999" # Starts a clock service on port 9999.
func NewServiceCmd() *cobra.Command {
	var serverPort string

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Start various services (clock, reverb, ftp, chat)",
		Long:  `Starts a specified service on a given port.`,
		Run: func(cmd *cobra.Command, args []string) {
			params := strings.Split(serverPort, ":")
			if len(params) < 2 {
				fmt.Printf("Invalid input: %q. Expect -sp=\"clock:9999\"", serverPort)
				return
			}
			service := params[0]
			port, _ := strconv.Atoi(params[1])
			switch service {
			case "clock":
				slog.Info("Started service", "service", service, "port", port)
				ClockServer(port)
			case "reverb":
				slog.Info("Started service", "service", service, "port", port)
				ReverbServer(port)
			case "ftp":
				slog.Info("Started service", "service", service, "port", port)
				FTPServer(port)
			case "chat":
				slog.Info("Started service", "service", service, "port", port)
				ChatService(port)
			default:
				slog.Info("Service not implemented", "service", service)
			}
		},
	}

	cmd.Flags().StringVar(&serverPort, "sp", "clock:9999", "service-type:port eg: \"clock:9999\" or \"reverb:9998\" or \"ftp:9997\"  or \"chat:9998\"")

	return cmd
}

// NewClientCmd creates the client command
// eg: the-gpl client --cp="clock:9999"
func NewClientCmd() *cobra.Command {
	var clientPort string

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Connect to services",
		Long:  `Connects to a running service.`,
		Run: func(cmd *cobra.Command, args []string) {
			params := strings.Split(clientPort, ":")
			if len(params) < 2 {
				fmt.Printf("Invalid input: %q. Expect -cp=\"clock:9999\"", clientPort)
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
		},
	}

	cmd.Flags().StringVar(&clientPort, "cp", "clock:9999", "client-type:port eg \"clock:9999\" or \"reverb:9998\" or \"chat:9998\"")

	return cmd
}

// NewDuCmd creates the du command
// eg: the-gpl du --dir="."
func NewDuCmd() *cobra.Command {
	var dir string
	var verbose bool

	cmd := &cobra.Command{
		Use:   "du",
		Short: "Disk Usage",
		Long:  `Calculate disk usage of a directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = DU(dir, verbose)
		},
	}

	cmd.Flags().StringVar(&dir, "dir", ".", "du -dir:\".\"")
	cmd.Flags().BoolVar(&verbose, "v", false, "du -v:false")

	return cmd
}
