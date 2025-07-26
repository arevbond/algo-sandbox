func checkInclusion(s1 string, s2 string) bool {
	cnt := make(map[rune]int)

	for _, ch := range s1 {
		cnt[ch]++
	}

	runes := []rune(s2)

	l := 0

	k := len([]rune(s1))

	for r, cur := range runes {
		cnt[cur]--

		if cnt[cur] == 0 {
			delete(cnt, cur)
		}

		if r - l + 1 > k {
			last := runes[l]
			cnt[last]++

			if cnt[last] == 0 {
				delete(cnt, last)
			}

			l++
		}

		if len(cnt) == 0 {
			return true
		}
	}

	return false
}
