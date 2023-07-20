package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	newS := ""
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			newS += strings.ToLower(string(c))
		}
	}
	i, j := 0, len(newS)-1
	for i < j {
		if newS[i] != newS[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	//fmt.Println(unicode.IsLetter(' '))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("0P"))
}
