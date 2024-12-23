package main

import "fmt"

/*

   eceba  k = 3
   result:  eceb

   1. mp counter, r l pointers
   2. if len(mp) <= k update maxLen
   3. if len(mp) < cur window then move r
   4. if len(mp) > cur winrow then move l while len(mp) == k

*/

func LengthOfLongestSubstringKDistinct(s string, k int) int {
    mp := make(map[rune]int)
    var result int
    runes := []rune(s)
    l := 0
    for r, ch := range runes {
        mp[ch]++
       
        if len(mp) <= k {
            result = max(result, r - l + 1)
        }

        for ; len(mp) > k && l < r; l++ {
            mp[runes[l]]--
            if mp[runes[l]] == 0 {
                delete(mp, runes[l])
            }
        }

    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(LengthOfLongestSubstringKDistinct("eceba", 3))
    fmt.Println(LengthOfLongestSubstringKDistinct("WORDL", 4))
}
