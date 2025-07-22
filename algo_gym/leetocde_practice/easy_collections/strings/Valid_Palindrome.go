func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), "")
	
	runes := []rune(s)
	l, r := 0, len(runes) - 1

	for l < r {
		if !unicode.IsLetter(runes[l]) && !unicode.IsDigit(runes[l]) {
			l++ 
			continue
		}

		if !unicode.IsLetter(runes[r]) && !unicode.IsDigit(runes[r]) {
			r--
			continue
		}


		if runes[l] != runes[r] {
			return false
		}

		l++
		r--

	}

	return true
}
