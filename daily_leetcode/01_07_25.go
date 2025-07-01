// 3330. Find the Original Typed String I

func possibleStringCount(word string) int {
	result := 1

	runes := []rune(word)

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			result++
		}
	}

	return result
}
