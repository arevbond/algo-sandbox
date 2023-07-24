package main

import "fmt"

func contains(x, str string) bool {
	for _, ch := range str {
		if string(ch) == x {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var curLen, maxLen int
	l, r := 0, 1
	for ; r < len(s); r++ {
		for contains(string(s[r]), s[l:r]) && l < r {
			l++
		}
		curLen = len(s[l:r+1])
		maxLen = max(curLen, maxLen)
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
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}