package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	mapString := make(map[string][]string)
	for _, str := range strs {
		sliceStr := strings.Split(str, "")
		sort.Strings(sliceStr)
		sortedStr := strings.Join(sliceStr, "")
		mapString[sortedStr] = append(mapString[sortedStr], str)
	}
	var answer [][]string
	for _, slicesStrs := range mapString {
		answer = append(answer, slicesStrs)
	}
	return answer
}
