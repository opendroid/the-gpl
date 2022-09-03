package chapter5

import (
	"testing"
)

// preRequisites for various courses in the university
var preRequisites = map[string][]string{
	"algorithms":             {"data structures"},
	"calculus":               {"linear algebra"},
	"compilers":              {"data structures", "formal languages", "computer organization"},
	"data structures":        {"discrete math"},
	"databases":              {"data structures"},
	"discrete math":          {"intro to  programming"},
	"formal languages":       {"discrete math"},
	"networks":               {"operating systems"},
	"operating systems":      {"data structures", "computer organization"},
	"programming  languages": {"data structures", "computer organization"},
}

// preRequisites2 for various courses in the university
var preRequisites2 = map[string]map[string]bool{
	"algorithms":             {"data structures": true},
	"calculus":               {"linear algebra": true},
	"compilers":              {"data structures": true, "formal languages": true, "computer organization": true},
	"data structures":        {"discrete math": true},
	"databases":              {"data structures": true},
	"discrete math":          {"intro to  programming": true},
	"formal languages":       {"discrete math": true},
	"networks":               {"operating systems": true},
	"operating systems":      {"data structures": true, "computer organization": true},
	"programming  languages": {"data structures": true, "computer organization": true},
}

// TestTopologicalSort sorts a test map of courses
//
//	cd chapter5
//	go test -run TestTopologicalSort -v
func TestTopologicalSort(t *testing.T) {
	t.Run("Slice sorted order", func(t *testing.T) {
		topology := TopologicalSort(preRequisites)
		for i, course := range topology {
			t.Logf("%d: %s", i+1, course)
		}
	})
}

// TestTopologicalSortMap sorts a test map of courses
//
//	cd chapter5
//	go test -run TestTopologicalSortMap -v
func TestTopologicalSortMap(t *testing.T) {
	t.Run("Map sorted order Run 1", func(t *testing.T) {
		topology := TopologicalSortMap(preRequisites2)
		for i, course := range topology {
			t.Logf("%d: %s", i+1, course)
		}
	})
	t.Run("Map sorted order Run 2", func(t *testing.T) {
		topology := TopologicalSortMap(preRequisites2)
		for i, course := range topology {
			t.Logf("%d: %s", i+1, course)
		}
	})
}
