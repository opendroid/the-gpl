# Chapter 4: Composite Types

## Arrays
Arrays are fixed length sequences of one or more elements of a particular type (homogeneous elements).
If array element is comparable, then the array type is comparable. Arrays are passed by-copy to functions, 
that may be inefficient. Any changes made to by functions to the passed array is only done on the 
copy received.

By default, array elements are initialized to zero, eg
```go
{
  // Initialize few elements of an array
  var planets = [8]string{"Mercury", "Venus", 4: "Jupiter"} // Unspecified take "" values
} 
```

## Slices
Lightweight data structure `(pointer, length, capacity)` that gives access to underlying array.
Multiple slices can share the same underlying array. See image below. **Slices are not comparable**. 
To look for empty slice perform `len(s) == 0` not `s == nil`.
```go
var s []int // len(s) == 0, s == nil
s = nil     // len(s) == 0, s == nil
s = []int{} // len(s) == 0, s != nil
```
Go function should treat `len(s) == 0` and `s == nil` slices same way.

To make slice use `make([]T, len, cap)` built-in function. To append to a slice use pattern:
```go
slice = append(slice, "Whatever") // Returns updated space structure pointer.  
```


## Maps
In `Go` a __map__ is a reference to **Hash Table**, written as __map[k]V__. It is collection of
unordered key, value pairs. Features of map are:

- Allocate before using maps
- The key type **K** must be comparable  (eg: **not `slice` or `map`**) using `==`. Avoid floating-point numbers as keys.
- The value **V** can be any composite type including a `map`.
- Cannot take address of a map elements eg `_ = &ages["golang"] // compile error`.
- Order of map iteration is not specified, eg `for k, v in range {fmt.Printf("%s:%s", k, v)}` order is not guaranteed.
- `nil` map reference behaves as `empty` maps for operations **delete, len** and **range**.
- Accessing a map by subscripting always returns a value, even if key is not there.
- Return value for a non-existent key always returns zero value for its type.
- `map` is not comparable however you can use `reflect.DeepEqual(a, b)` to compare. [Playground example](https://play.golang.org/p/JbD9fVsS4jS)

### Special note
Go does not provide a `set` composite type as maps key can serve this purpose. See example below [Structs and map](#structs-and-map). 

## Structs
A `struct` groups together zero or more names values of arbitrary type as a single entity. Each value is called a
`field`. 

- The `Field order` is significant in struct definitions. [Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard) by [Dave Cheney](https://dave.cheney.net/).
- Changing field order will define different struct type.
- Fields can be exported or un-exported. A field starting with a capital letter is exported.
- A named struct cannot contain a field with same type as struct, instead use a pointer to it.
- Zero value of struct is composed of zero value of each of fields. 
- Zero value should be meaningful. eg: `bytes.Buffer` or `sync.Mutex` zero-values are ready-to-use.
- If all fields of struct are comparable, then structs are comparable. Those can be used as map keys.
- `struct embedding` mechanism allows named struct as anonymous filed. It is syntactic sugar on dot notation. eg

```go
type Point struct {X, Y int}
type Circle struct { Point; Radius int} // anonymous center
type Wheel struct { Circle; Spokes int}  
var w Wheel = Wheel{Circle{Point{8, 10}, 5}, 10} // Wheel with Circle at (8, 10) of radius = 5 with 10 spokes
```   
Check on [playground](https://play.golang.org/p/w99AaWcePFA). 

### Structs and map
A struct with no fields is called `empty struct`. Some go programmers use it as a value of map that represents a set.

```go
seen := make(map[string]struct{}) // set of strings
if _, ok = seen["go"]; !ok {
  if seen[s] == struct{} 
}
```

## Javascript Object Notation (JSON)
JSON is useful notation to send and receive structured information. Go's standard library [encoding/json](https://golang.org/pkg/encoding/json/) provides 
support to parse and marshal JSON to golang `struct`. See [playground JSON example](https://play.golang.org/p/Cp3Uo7qS4O9)
The library uses `field tags`, a metadata string associated with field at compile time. The field tags are `space` separated
 `key:value` pairs. In example below `json` key controls behavior of `encoding/json` and `xml` of `encoding/xml`.

```go
	type State struct {
		Name         string  `xml:"name,omitempty" json:"name,omitempty"`
		Capital      string  `xml:"capital,omitempty" json:"capital,omitempty"`
		PopulationMM float64 `xml:"population_millions,omitempty" json:"population_millions,omitempty"`
		GDPBillions  int16   `xml:"gdp_billions,omitempty" json:"gdp_billions,omitempty"`
	}
```
Note that names of struct fields that json encoder should manipulate should be exported i.e capitalized. 
This metadata will provide hints to library to use a JSON as:
```json
{
	"name": "California",
	"capital": "Sacramento",
	"population_millions": 40000000,
	"gdp_billions": 4000
}
```
An example of XML data for this struct is:
```xml
<State>
	<name>California</name>
	<capital>Sacramento</capital>
	<population_millions>4e+07</population_millions>
	<gdp_billions>4000</gdp_billions>
</State>
```

### Marshal
`json.Marshal` provides a way to convert a golang struct to string with no extraneous white spaces.

### Unmarshal
Converts a []byte, i.e stringified JSON, in a golang memory location.

**Note:** Similar utilities are provided to parse XML as well. See [playground XML example](https://play.golang.org/p/TIBzdIyyNTm).

## Pictures
![Solar System Bodies](../content/media/slicesBodies.jpeg?raw=true "Bodies in Solar System")
