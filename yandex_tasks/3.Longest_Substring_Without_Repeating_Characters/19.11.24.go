package longestsubstringwithoutrepeatingcharactersc

// bacabcabc
//    ^
//   ^
func lengthOfLongestSubstring(s string) int {
    var maxLen int
    counter := make(map[rune]int)
    runes := []rune(s)
    l := 0
    for r := range runes {
        counter[runes[r]]++
        if counter[runes[r]] > 1 {
            for ; l < r && counter[runes[r]] > 1; l++ {
                counter[runes[l]]--
                if counter[runes[l]] == 0 {
                    delete(counter, runes[l])
                }
            }
        }
        maxLen = max(maxLen, len(counter))
    }
    return maxLen
}
