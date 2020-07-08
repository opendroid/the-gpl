package mas

import (
	"fmt"
)

// Stringer a type assertion takes a value and tries to create another version in specified explicit type
type Stringer interface {
	StringMe() string
}

type fakeString struct {
	content string
}

func (s *fakeString) StringMe() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.StringMe())
	}
}

func learnStringerInterface() {
	s := &fakeString{content: "Awesome content"}
	s1 := &fakeString{"Default content Name"}
	printString(s)
	printString(s1)
	printString("Just a string Sir")
}
