func maxDistToClosest(seats []int) int {
	forwardArr := make([]int, len(seats))
	copy(forwardArr, seats)
	
	closestIndx := len(seats) - 1
	for i := 0; i < len(forwardArr); i++ {
		if forwardArr[i] == 1 {
			forwardArr[i] = -1	
			closestIndx = i
			continue
		}
		
		forwardArr[i] = abs(closestIndx - i)
	}

	result := 0
	closestIndx = 0
	for i := len(forwardArr) - 1; i >= 0; i-- {
		if forwardArr[i] == -1 {
			closestIndx = i
			continue
		}

		forwardArr[i] = min(forwardArr[i], abs(closestIndx - i))	

		result = max(result, forwardArr[i])
	}

	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -1 * a
}
