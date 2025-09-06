package taskTwentySix

import "strings"

func UniqueSymbols(s string) bool {
	rs := []rune(s)
	set := make(map[string]interface{})

	for _, val := range rs {
		sVal := strings.ToLower(string(val))
		_, ok := set[sVal]
		if ok {
			return false
		}
		set[sVal] = struct{}{}
	}

	return true
}
