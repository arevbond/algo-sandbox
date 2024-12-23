package permutationinstring

// eidbaooo 
// --
//  --
//   -- 
//    --
func checkInclusion(s1 string, s2 string) bool {
   counter := make(map[rune]int)
    for _, ch := range s1 {
        counter[ch]++
    }
    
    // size of sliding window
    k := len(s1)
    runes := []rune(s2)
    l := 0
    for r, cur := range runes {
        counter[cur]--
        
        if counter[cur] == 0 {
            delete(counter, cur)
        }
    
        if len(counter) == 0 {
            return true
        }

        if r - l + 1 == k {
            counter[runes[l]]++
            if counter[runes[l]] == 0 {
                delete(counter, runes[l])
            }
            l++
        }
    }
    return false
}
