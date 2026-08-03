package web

import "fmt"

// bookChapters is the single source of truth for the nine book chapters. It is
// rendered on /chapters (chaptersHandler) and reused by chapterContext to give
// the AI tutor chapter-aware context.
var bookChapters = []ChapterEntry{
	{1, "Tutorial", "Goroutines, channels, CLI utilities, the Lissajous GIF, a Dialogflow bot, and Speech-to-Text.", chapterURL(1)},
	{2, "Program Structure", "Bit counting via three strategies, temperature-conversion types, and package-level variables.", chapterURL(2)},
	{3, "Basic Data Types", "Mandelbrot PNG, 3-D surface plots as SVG, and string utilities.", chapterURL(3)},
	{4, "Composite Types", "JSON marshalling, HTML templating, and GitHub issue search.", chapterURL(4)},
	{5, "Functions", "HTML traversal, a web crawler, topological sort, variadic max/min, generic MaxOf/MinOf.", chapterURL(5)},
	{6, "Methods", "An IntSet bit-vector: Union, Intersect, Difference, SymmetricDifference.", chapterURL(6)},
	{7, "Interfaces", "Writer implementations, CountWriter, BroadcastWriters, and a temperature flag.", chapterURL(7)},
	{8, "Goroutines & Channels", "TCP services (clock, reverb, chat, FTP), concurrent du, and web search with context.", chapterURL(8)},
	{9, "Concurrency & Shared Variables", "sync.Mutex SafeBank, sync.RWMutex RWBank, sync.Once Icon, and a Memo cache.", chapterURL(9)},
}

// Prompt is one suggested tutor question. Label is the short chip text; Q is the
// full question prefilled into the textarea. JSON tags feed the front-end.
type Prompt struct {
	Label string `json:"label"`
	Q     string `json:"q"`
}

// chapterPrompts maps a chapter <select> value ("" = no chapter) to its suggested
// questions. Keys match exactly the options rendered in ask.gohtml.
var chapterPrompts = map[string][]Prompt{
	"": {
		{"buffered vs unbuffered", "What's the difference between a buffered and an unbuffered channel?"},
		{"when to use a mutex", "When should I use a sync.Mutex instead of a channel?"},
		{"select statement", "How does the select statement work in Go?"},
	},
	"1": {
		{"goroutines", "What is a goroutine and how do I start one?"},
		{"the Lissajous demo", "How does the Lissajous GIF program generate its animation?"},
		{"reading stdin", "How do I read input line by line from standard input in Go?"},
	},
	"2": {
		{"iota & constants", "How does iota work when declaring a set of constants?"},
		{"bit counting", "What are the different ways to count the set bits in an integer?"},
		{"named types", "Why define a named type like Celsius instead of using a plain float64?"},
	},
	"3": {
		{"rune vs byte", "What's the difference between a rune and a byte in Go?"},
		{"complex numbers", "How do I work with complex numbers, like in the Mandelbrot demo?"},
		{"floating point", "How do floating-point precision issues show up in Go?"},
	},
	"5": {
		{"variadic functions", "How do variadic functions work in Go?"},
		{"closures", "What is a closure and when is it useful?"},
		{"defer", "How does defer work and when do deferred calls run?"},
	},
	"6": {
		{"pointer vs value receiver", "When should a method use a pointer receiver instead of a value receiver?"},
		{"method sets", "What is a method set and how does it affect interface satisfaction?"},
		{"IntSet bit-vector", "How does the IntSet bit-vector represent a set of integers?"},
	},
	"7": {
		{"interface satisfaction", "How does a type satisfy an interface in Go?"},
		{"empty interface", "What is the empty interface and when should I use it?"},
		{"type assertions", "How do type assertions and type switches work?"},
	},
	"8": {
		{"channel directions", "What are send-only and receive-only channel types for?"},
		{"pipelines", "How do I build a pipeline of goroutines connected by channels?"},
		{"context cancellation", "How do I cancel running goroutines with a context?"},
	},
	"9": {
		{"data races", "What is a data race and how do I detect one in Go?"},
		{"Mutex vs RWMutex", "When should I use sync.RWMutex instead of sync.Mutex?"},
		{"sync.Once", "What is sync.Once used for?"},
	},
}

// chapterContext returns a one-line context string for the AI tutor describing
// the selected chapter, or "" when no (or an unknown) chapter is selected.
func chapterContext(chapter string) string {
	if chapter == "" {
		return ""
	}
	for _, c := range bookChapters {
		if fmt.Sprintf("%d", c.Number) == chapter {
			return fmt.Sprintf("This question relates to Chapter %d — %s: %s", c.Number, c.Title, c.Description)
		}
	}
	return ""
}
