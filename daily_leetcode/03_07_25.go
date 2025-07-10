// 3304. Find the K-th Character in String Game I

func kthCharacter(k int) byte {
	z := 'z'

	word := []rune("a")

	for len(word) <= k {
		for _, ch := range word {
			newCh := ch + 1
			if ch == z {
				newCh = 'a'
			}
			word = append(word, newCh)

			if len(word) == k {
				break	
			}
		}
	}

	return byte(word[k-1])
}
