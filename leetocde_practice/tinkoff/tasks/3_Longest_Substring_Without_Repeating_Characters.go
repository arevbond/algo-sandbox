package tasks

// Time: O(N^2)
func lengthOfLongestSubstringN2(s string) int {
	result := 0

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		mp := make(map[rune]struct{})
		cur := 0
		for j := i; j < len(runes); j++ {
			if _, ok := mp[runes[j]]; ok {
				break
			}

			mp[runes[j]] = struct{}{}
			cur++
		}

		result = max(result, cur)
	}

	return result
}

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	result := 0

	for l := 0; l < len(runes); {
		set := make(map[rune]int)

		r := l
		indxToBack := -1
		for ; r < len(runes); r++ {
			if indx, ok := set[runes[r]]; ok {
				indxToBack = indx + 1
				break
			}
			set[runes[r]] = r
		}
		result = max(result, r-l)
		if indxToBack != -1 {
			l = indxToBack
		} else {
			l = r
		}
	}

	return result
}

func lengthOfLongestSubstringPAST(s string) int {
	var maxLen int
	counter := make(map[rune]int)
	runes := []rune(s)
	l := 0
	for r := range runes {
		counter[runes[r]]++
		if counter[runes[r]] > 1 {
			for ; l < r && counter[runes[r]] > 1; l++ {
				counter[runes[l]]--
				if counter[runes[l]] == 0 {
					delete(counter, runes[l])
				}
			}
		}
		maxLen = max(maxLen, len(counter))
	}
	return maxLen
}
