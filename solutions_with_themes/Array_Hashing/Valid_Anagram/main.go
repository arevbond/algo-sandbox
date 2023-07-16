package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dct1, dct2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dct1[s[i]]++
		dct2[t[i]]++
	}
	for k, v := range dct1 {
		if v2, _ := dct2[k]; v2 != v {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
