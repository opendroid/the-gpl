package chapter5

import (
	"sort"
)

// Toposort sorts a  map into a topo sorted string of dependencies
func Toposort (m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func (items []string) // recursive func, required for recursion
	visitAll = func (items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}


// ToposortMap topo-sorts using map.
//   Exercise 5.10: Rewrite topoSort to use maps instead of slices and eliminate the initial sort.
//   Verify that the results, though nondeterministic, are valid topological orderings.
func ToposortMap(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func (items map[string]bool) // recursive func
	visitAll = func (items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order
}