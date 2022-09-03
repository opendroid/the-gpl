// Package structs provide example of a struct and its receiver or pointer receiver methods.
// Example demonstrates the difference between value and pointer receiver methods.
// The value receiver methods are called on a copy of the struct.
// Whereas, the pointer receiver methods are called on the original struct.
package structs

import "fmt"

// Thakur is a family member is example of pointer receiver methods.
type Thakur struct {
	Age  int
	Name string
}

// NewThakur creates a member
func NewThakur(name string, age int) *Thakur {
	return &Thakur{Name: name, Age: age}
}

// GetAThakur creates a Thakur
func GetAThakur() Thakur {
	return Thakur{9000, "Ajay"}
}

// ChangeToHeadOfHousehold changes the name and age
func (ath *Thakur) ChangeToHeadOfHousehold() {
	ath.Name, ath.Age = "Sofia", 40
}

// Describe method pattern to show name and age fields
func (ath *Thakur) Describe() {
	fmt.Printf("Thakur: %s: %d\n", ath.Name, ath.Age)
}

// GotMarried adds a Mrs in beginning
func (ath *Thakur) GotMarried() {
	ath.Name = "Mrs. " + ath.Name
}

// ThoughtIGotMarried is just the same
func (ath *Thakur) ThoughtIGotMarried() {
	ath.Name = "Mrs. " + ath.Name
}

// ThakurCopy is a copy of Thakur, does not modify the original.
type ThakurCopy Thakur

// Describe method pattern to show name and age fields
func (ath ThakurCopy) Describe() {
	fmt.Printf("Thakur: %s: %d\n", ath.Name, ath.Age)
}

// ValueDescribe shows name after copying
func (ath ThakurCopy) ValueDescribe() {
	fmt.Printf("Value describe Thakur: %s: %d\n", ath.Name, ath.Age)
}

// ThoughtIGotMarried does not change the original. Ignore the [lint warning].
//
// [lint warning]: https://golangci-lint.run/usage/false-positives/
//
//nolint:staticcheck
func (ath ThakurCopy) ThoughtIGotMarried() {
	ath.Name = "Mrs. " + ath.Name
}

// GotMarried adds a Mrs in beginning, it is ineffective, ignore the lint warning.
//
//nolint:staticcheck
func (ath ThakurCopy) GotMarried() {
	ath.Name = "Mrs. " + ath.Name
}

// ChangeToHeadOfHousehold changes the name and age
//
//nolint:staticcheck
func (ath ThakurCopy) ChangeToHeadOfHousehold() {
	ath.Name, ath.Age = "Sofia", 40
}
