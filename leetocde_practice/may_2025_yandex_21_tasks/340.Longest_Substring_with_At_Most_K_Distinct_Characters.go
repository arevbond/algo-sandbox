/*
Given a string, find the length of the longest substring T that contains at most k distinct characters.

Example 1:

Input: s = "eceba", k = 2
Output: 3
Explanation: T is "ece" which its length is 3.
Example 2:

Input: s = "aa", k = 1
Output: 2
Explanation: T is "aa" which its length is 2.
*/

// e c e b b b b d
//     l
//       r 
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	var result int

	cnt := make(map[rune]int)

	runes := []rune(s)

	l := 0

	for r, ch := range runes {
		cnt[ch]++

		for len(cnt) > k {
			last := runes[l]
			cnt[last]--

			if cnt[last] == 0 {
				delete(cnt, last)
			}

			l++
		}

		result = max(result, r - l + 1)
	}


	return result
}


