package main

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"github.com/opendroid/the-gpl/chapter1/graphs"
	mas "github.com/opendroid/the-gpl/chapter1/mas"
	server "github.com/opendroid/the-gpl/chapter1/wwwexamples"
	"github.com/opendroid/the-gpl/chapter2"
	"os"
)

// Maps argument="methodToCall" to function call.
var funcMap = map[string]func(){
	"callMas":         callMas,
	"fetchSites":      fetchSites,
	"saveSquareImage": saveSquareImage,
	"tempConv":        tempConv,
	"tryDefer":        tryDefer,
	"server":          setupWebServer,
}

// main: go run main.go -func="methodToCall"
func main() {
	flag.Parse()
	if funcMap[*callMethod] != nil {
		funcMap[*callMethod]() // execute func
	} else {
		fmt.Println("Usage: LangGO -func={callMas|fetchSites|saveSquareImage|..}")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

// Flag to run
var callMethod = flag.String("func",
	"",
	"Method to call and test. (Required)")

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

func tempConv() {
	boilingPointC := chapter2.BoilingPointF.ToC().String()
	fmt.Printf("tempConv: %s = %s\n", chapter2.BoilingPointF.String(), boilingPointC)
}
