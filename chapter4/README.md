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

![Solar System Bodies](../content/media/slicesBodies.jpeg?raw=true "Bodies in Solar System")
