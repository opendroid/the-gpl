package main

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/audio"
	"github.com/opendroid/the-gpl/chapter1/bot"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"github.com/opendroid/the-gpl/chapter1/graphs"
	mas "github.com/opendroid/the-gpl/chapter1/mas"
	server "github.com/opendroid/the-gpl/chapter1/wwwexamples"
	"github.com/opendroid/the-gpl/chapter2"
	"os"
)

// main: go run main.go -func="setupWebServer"
func main() {
	// Flag in Main
	if len(os.Args) < 2 {
		printArgsHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "func":
		execCmd(os.Args[2:])
	case "bot":
		bot.ExecCmd(os.Args[2:])
	case "audio":
		audio.ExecCmd(os.Args[2:])
	case "temp":
		chapter2.ExecTempConvCmd(os.Args[2:])
	case "bits":
		chapter2.ExecBitsCountCmd(os.Args[2:])
	}
}

// Maps argument="methodToCall" to function call.
var funcMap = map[string]func() {
	"mas":    callMas,
	"fetch":  fetchSites,
	"image":  saveSquareImage,
	"defer":  tryDefer,
	"server": setupWebServer,
}

// cmd flag set for methods
var cmd = flag.NewFlagSet("func", flag.ContinueOnError)
// callMethod method to call.
var callMethod = cmd.String("name", "", "Method to call and test.")
// execCmd execute function maps
func execCmd(args []string)  {
	err := cmd.Parse(args)
	if err != nil {
		_ = fmt.Errorf("func error %s", err)
		os.Exit(1)
	}
	if fn, ok := funcMap[*callMethod]; ok {
		fn() // execute func
	} else {
		fmt.Printf("%s func %s does not exist\n", os.Args[0], *callMethod)
		os.Exit(1)
	}
}
// printArgsHelp Prinst help for all arguments
func printArgsHelp() {
	fmt.Println("Usage: the-gpl func")
	cmd.PrintDefaults()
	fmt.Print("\tMethods: ")
	for k := range funcMap {
		fmt.Printf("%s ", k)
	}
	fmt.Println("\nUsage: the-gpl bot")
	bot.Cmd.PrintDefaults()
	fmt.Println("\nUsage: the-gpl audio. First run this command")
	fmt.Println(`ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ar 48000 -f s16le udp://localhost:9999`)
	audio.Cmd.PrintDefaults()
	fmt.Println("\nUsage: the-gpl temp. Coverts c to f and visa versa")
	chapter2.Cmd.PrintDefaults()
	fmt.Println("\nUsage: the-gpl bits. Counts 1 bits in 64-bit int")
	chapter2.BitCountCmd.PrintDefaults()
}

//  deferC Deferred functions may read and assign to the returning function's
//  named return values. Returns 2. then defer multiples that
func deferC() (i int) {
	defer func() { i = i * 3 }() // Operates on return value
	return 2
}

func tryDefer() {
	fmt.Printf("deferC: %d\n", deferC())
}

// callMas: Tests the mass functions
func callMas() {
	mas.IterateOverArray()
	compResult, diff := mas.CompareNumbers(5, 2)
	fmt.Printf("mas:CompareNumbers: ints: %d == %d, => %t, differance: %d\n", 5, 3, compResult, diff)
	mas.AddToSlices()
}

// saveImageSquare tests graphs package
func saveSquareImage() {
	config := graphs.LissajousConfig{
		Cycles:     2,
		Resolution: 0.000001,
		Size:       512,
		NFrames:    12,
		DelayMS:    10,
	}
	graphs.Lissajous(config, os.Stdout)
}

// setupWebServer tests server package
func setupWebServer() {
	fmt.Println("Startgin server, try commands like:")
	fmt.Println("http://localhost:8080/")
	fmt.Println("http://localhost:8080/graph")
	fmt.Println("http://localhost:8080/egg")
	fmt.Println("http://localhost:8080/sinc")
	fmt.Println("http://localhost:8080/mandel")
	fmt.Println("http://localhost:8080/mandelbw")
	fmt.Println("http://localhost:8080/incr")
	fmt.Println("http://localhost:8080/counter")
	fmt.Println("http://localhost:8080/post?q=\"Ajay\"&r=\"Thakur\"&son=Aiden")
	server.ServerMethodOne()
}

func fetchSites() {
	testSites := [...]string{"https://google.com", "https://youtube.com", "https://facebook.com",
		"https://qq.com", "https://amazon.com", "https://usense.io",
	}
	sitesChan := make(chan string) // Make 1 channel only
	for _, site := range testSites {
		go channels.Fetch(site, sitesChan)
	}
	// Expect 5 responses in channel
	for range testSites {
		fmt.Printf("%s\n", <-sitesChan)
	}
}

