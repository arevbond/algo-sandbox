package main

import (
	"unicode"
)

/*
   1. left and right pointers
   2. skip non alphabet chars
   3. make left and right chars to lower
*/
func isPalindrome(s string) bool {
    runes := []rune(s) 
    l, r := 0, len(runes) - 1

    for l < r {
        if !unicode.IsLetter(runes[l]) && !unicode.IsNumber(runes[l]) {
            l++
            continue
        }

        if !unicode.IsLetter(runes[r]) && !unicode.IsNumber(runes[r]) {
            r--
            continue
        } 

        if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
            return false
        }
        l++
        r--
    }
    return true
}
