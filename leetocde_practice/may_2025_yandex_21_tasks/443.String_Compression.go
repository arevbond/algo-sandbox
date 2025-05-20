func compress(chars []byte) int {
	indx := 1
	l := 0

	for ; l < len(chars); {
		r := l + 1
		for ; r < len(chars) && chars[r] == chars[r-1]; r++ {}

		count := r - l

		if count > 1 {
			for _, ch := range strconv.Itoa(count) {
				chars[indx] = byte(ch)
				indx++
			}
		}

		if r >= len(chars) {
			break
		}

		chars[indx] = chars[r]
		indx++

		l = r
	}

	return indx
}
