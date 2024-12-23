package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	hashSet := make(map[rune]bool)
	l := 0
	maxLen := 0
	runes := []rune(s)
	for r := range runes {
		for hashSet[runes[r]] {
			hashSet[runes[l]] = false
			l++
		}
		hashSet[runes[r]] = true
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	hashSet := make(map[rune]bool)
	l := 0
	maxLen := 0
	runes := []rune(s)
	for r, _ := range runes {
		for {
			if _, ok := hashSet[runes[r]]; ok {
				delete(hashSet, runes[l])
				l++
			} else {
				break
			}
		}
		hashSet[runes[r]] = true
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring1("dvdf"))
}
