// correct, but extra allocations
func isAnagramExtra(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// char: cnt
	mpS := make(map[rune]int)

	for _, ch := range s {
		mpS[ch]++
	}
	
	// char: cnt
	mpT := make(map[rune]int)

	for _, ch := range t {
		if _, ok := mpS[ch]; !ok {
			return false
		}
		mpT[ch]++
	}

	for ch, cntT := range mpT {
		cntS := mpS[ch]
		if cntS != cntT {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {
	cntChars := make(map[rune]int)
	for _, ch := range s {
		cntChars[ch]++
	}

	for _, ch := range t {
		cntChars[ch]--
		if cntChars[ch] == 0 {
			delete(cntChars, ch)
		}
	}

	return len(cntChars) == 0
}
