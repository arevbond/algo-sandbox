// 1. find minX and maxX
// 2. make set of points 
// 3. calculate totalSumX
// 4. for each point find mirror point with equal Y and x - totalSumX
func IsReflected(points [][]int) bool {
	set := make(map[[2]int]struct{})	
	minX := math.MaxInt
	maxX := -1 * math.MaxInt
	for _, point := range points {
		x, y := point[0], point[1]

		minX = min(x, minX)
		maxX = max(x, maxX)

		set[[2]int{x, y}] = struct{}{}
	}

	totalSumX := minX + maxX

	for point := range set {
		x1, y1 := point[0], point[1]
		
		if _, ok :=  set[[2]int{(totalSumX - x1), y1}]; !ok {
			return false
		}
	}

	return true
}
