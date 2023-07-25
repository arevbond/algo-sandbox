package main

import "fmt"

func characterReplacement(s string, k int) int {
    maxLen, curLen := 1, 1
    l, r := 0, 1 
    for ; r < len(s) && k > 0; r++ {
    	
    }
}


func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}