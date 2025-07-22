func longestCommonPrefix(strs []string) string {
	result := ""

	minLen := math.MaxInt32
	for _, str := range strs {
		minLen = min(minLen, len(str))
	}

	for i := 0; i <= minLen; i++ {
		if i >= len(strs[0]) {
			return result
		}
		curCh := strs[0][i]
		for _, str := range strs {
			if i >= len(str) || str[i] != curCh {
				return result
			}
		}
		result += string(curCh)
	}
	return result
}
