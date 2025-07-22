func firstUniqChar(s string) int {
	// char: [indx, cnt]
	mp := make(map[rune][2]int)

	for i, ch := range s {
		if _, ok := mp[ch]; !ok {
			mp[ch] = [2]int{i, 0}
		}
		pair := mp[ch]
		mp[ch] = [2]int{pair[0], pair[1]+1}
	}

	result := len([]rune(s)) - 1
	hasUniq := false

	for _, pair := range mp {
		indx, cnt := pair[0], pair[1]
		if cnt > 1 {
			continue 
		}

		hasUniq = true

		result = min(result, indx)
	}
	if hasUniq {
		return result
	}

	return -1
}
