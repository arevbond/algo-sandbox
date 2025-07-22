package isomorphicstrings

// 
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mpChStoChT := make(map[byte]byte)
    mpChTtoChS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        chS, chT := s[i], t[i]

        expectedChT, ok1 := mpChStoChT[chS]
        expectedChS, ok2 := mpChTtoChS[chT]
        if ok1 != ok2 {
            return false
        }

        if ok1 && (expectedChS != chS || expectedChT != chT) {
            return false
        }
        
        mpChStoChT[chS] = chT
        mpChTtoChS[chT] = chS
    }
    return true
}
