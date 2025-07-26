func minSteps(s string, t string) int {
	if len(s) < len(t) {
		s, t = t, s 
	}
	
	mpS, mpT := make(map[rune]int), make(map[rune]int)    

	for _, ch := range s {
		mpS[ch]++
	}

	for _, ch := range t {
		mpT[ch]++
	}

	var result int

	seen := make(map[rune]struct{})

	for ch, cntS := range mpS {
		if _, ok := seen[ch]; ok {
			continue 
		}
		cntT := mpT[ch]

		result += max(cntS, cntT) - min(cntS, cntT)

		seen[ch] = struct{}{}
	}

	for ch, cntT := range mpT {
		if _, ok := seen[ch]; ok {
			continue 
		}
		cntS := mpS[ch]

		result += max(cntS, cntT) - min(cntS, cntT)

		seen[ch] = struct{}{}

	}

	return result
}
