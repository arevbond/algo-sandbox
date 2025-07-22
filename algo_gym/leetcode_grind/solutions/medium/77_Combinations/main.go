func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var backtrack(start, comb) 
	backtrack = func(startint, comb []int) {
		if len(comb) == k {
			res = append(res, comb)
		}
	}   
}