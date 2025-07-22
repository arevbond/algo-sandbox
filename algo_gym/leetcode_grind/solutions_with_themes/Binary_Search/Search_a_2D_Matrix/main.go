func searchMatrix(matrix [][]int, target int) bool {
	indx := -1
	for i, row := range matrix {
		if target <= row[len(row)-1] && target >= row[0] {
			indx = i
		}
	}
	if indx == -1 {
		return false
	} else {
		l, r := 0, len(matrix[i]) - 1
		for l < r {
			m := ( l + r ) / 2
			if matrix[i][m] >= target {
				r = m
			} else {
				l = m + 1
			}
		}
		return indx != -1
	}
	return 
}