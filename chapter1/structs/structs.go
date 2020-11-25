// Package structs provide example of a struct and its receiver or pointer receiver methods.
package structs

import "fmt"

// Thakur family name
type Thakur struct {
	Age  int
	Name string
}

// NewThakur creates a member
func NewThakur(name string, age int) *Thakur {
	return &Thakur{Name: name, Age: age}
}

// GetAThakur creates a Thankur
func GetAThakur() Thakur {
	return Thakur{9000, "Ajay"}
}

// ChangeThakur changes the name and age
func ChangeThakur(aName *Thakur) {
	aName.Name, aName.Age = "Sofia", 40
}

// Describe method pattern to show name and age fields
func (ath *Thakur) Describe() {
	fmt.Printf("Thakur: %s: %d\n", ath.Name, ath.Age)
}

// ValueDescribe shows name after copying
func (ath Thakur) ValueDescribe() {
	fmt.Printf("Value describe Thakur: %s: %d\n", ath.Name, ath.Age)
}

// GotMarried adds a Mrs in beginning
func (ath *Thakur) GotMarried() {
	ath.Name = "Mrs. " + ath.Name
}

// ThoughtIGotMarried is just the same
func (ath Thakur) ThoughtIGotMarried() {
	ath.Name = "Mrs. " + ath.Name
}
