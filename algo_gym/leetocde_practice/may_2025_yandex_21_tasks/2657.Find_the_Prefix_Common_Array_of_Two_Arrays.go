func findThePrefixCommonArray(A []int, B []int) []int {
	seen := make(map[int]struct{})

	result := make([]int, len(A))

	count := 0

	for i := 0; i < len(A); i++ {
		a, b := A[i], B[i]

		if _, ok := seen[a]; ok {
			count++
		} else {
			seen[a] = struct{}{}
		}

		if _, ok := seen[b]; ok {
			count++
		} else {
			seen[b] = struct{}{}
		}

		result[i] = count
	}
	
	return result
}
