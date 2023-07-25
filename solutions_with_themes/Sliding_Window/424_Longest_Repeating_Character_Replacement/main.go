package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func characterReplacement(s string, k int) int {
    // maxLen = 0
    dic := make(map[byte]int)
    res := 0
    l, r, mostF := 0, 0, 0
    for ; r < len(s); r++ {
    	dic[s[r]] += 1
        
        mostF = max(mostF, dic[s[r]])

        if (r - l + 1) - mostF > k {
            dic[s[l]]--
            l++
        }
        res = max(res, r - l +1)
    }
    return res
}


func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	// fmt.Println(characterReplacement("AABABBA", 1))
}