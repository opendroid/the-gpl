# Chapter 5 — Functions

Examples from Chapter 5 of *The Go Programming Language*, covering first-class functions,
recursion, multiple return values, variadic functions, closures, and error handling.

## CLI

```bash
the-gpl parse --type=<type> --site=<url>
```

| `--type` value | Function called | What it does |
|---|---|---|
| `outline` | `ParseOutline` | Prints an outline of HTML tag structure |
| `links` | `ParseLinks` | Lists all unique `href` links |
| `images` | `ParseImages` | Lists all unique `<img src>` URLs |
| `css` | `ParseCss` | Lists all unique CSS `<link>` URLs |
| `scripts` | `ParseScripts` | Lists all unique `<script src>` URLs |
| `text` | `ParseText` | Prints text nodes (excluding script/style) |
| `pretty` | `PrettyHTML` | Pretty-prints indented HTML |
| `crawl` | `Crawl` | Downloads all pages in domain to `--dir` |

## Key Functions

### HTML Traversal (Exercise 5.1)

```go
// E51FindLinks recursively collects href values from an HTML parse tree.
links := chapter5.E51FindLinks([]string{}, rootNode)
```

Uses `golang.org/x/net/html` for parsing. Recursion replaces a loop to demonstrate
how the same tree walk can be expressed either way.

### Web Crawler (Exercise 5.13)

```go
count, err := chapter5.Crawl("https://example.com", "/tmp/crawl")
```

Downloads all pages within the same domain, mirroring the directory structure locally.

### String Expansion (Exercise 5.9)

```go
// Expand replaces $var in s by calling f("var").
result := chapter5.Expand("Hello, $name!", func(v string) string {
    return map[string]string{"name": "Go"}[v]
})
// result == "Hello, Go!"
```

### Topological Sort

```go
order := chapter5.TopologicalSort(map[string][]string{
    "algebra":  {"arithmetic"},
    "calculus": {"algebra"},
})
```

Classic DFS-based sort demonstrating closures capturing the `seen` map.

### Variadic Min/Max

```go
chapter5.MaxInt(3, 1, 4, 1, 5, 9)        // 9
chapter5.MaxIntOf(3, 1, 4, 1, 5, 9)      // 9  (requires at least one arg)
chapter5.MinInt(3, 1, 4, 1, 5, 9)        // 1
chapter5.Join(", ", "a", "b", "c")        // "a, b, c"
```

### Hyperlinks Sort Interface

```go
type Hyperlinks []string  // implements sort.Interface, ignores http/https scheme
```

## Running Tests

```bash
go test ./chapter5/...
go test -v -run TestCrawl ./chapter5/...
```

## Go Features Demonstrated

- Recursive functions on tree-structured data (`html.Node`)
- First-class functions passed as arguments (`func(float64, float64) float64`)
- Closures capturing variables from enclosing scope
- Multiple return values and named results
- Variadic functions (`...int`)
- `defer` for cleanup (file close, mutex unlock)
- Error values and wrapping with `%w`
- `golang.org/x/net/html` — HTML tokeniser and parse tree
- `sort.Interface` implementation on a named slice type
