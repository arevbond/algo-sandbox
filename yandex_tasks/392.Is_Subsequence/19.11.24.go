package issubsequence

// ace abcde = true
//    ^
//         ^
// aec
//   ^
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    if len(t) == 0 {
        return false
    }

    indxS := 0
    for i := range t {
        if indxS == len(s) {
            return true
        }
        if t[i] == s[indxS] {
            indxS++
        }
    }
    
    return indxS == len(s)
}
