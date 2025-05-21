func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	mpS, mpT := make(map[rune]rune), make(map[rune]rune)

	runesS, runesT := []rune(s), []rune(t)

	for i := 0; i < len(runesS); i++ {
		chS, chT := runesS[i], runesT[i]

		if tmpChT, ok := mpS[chS]; ok && tmpChT != chT {
			return false
		}

		if tmpChS, ok := mpT[chT]; ok && tmpChS != chS {
			return false
		}

		mpS[chS] = chT
		mpT[chT] = chS
	}

	return true
}
