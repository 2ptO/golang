package word

import "strings"

// UseCount returns the frequency of each word in a given string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the total number of words in a given string
func Count(s string) int {
	// write the code for this func
	return len(strings.Fields(s))
}
