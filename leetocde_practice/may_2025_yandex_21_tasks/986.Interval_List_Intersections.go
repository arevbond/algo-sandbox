func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int		
	
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		intersection := getIntersection(firstList[i], secondList[j])
		
		if intersection[0] <= intersection[1] {
			result = append(result, intersection)
		}

		if intersection[1] == firstList[i][1] {
			i++
		} else {
			j++
		}
	}
	
	return result
}

func getIntersection(firstList []int, secondList []int) []int {
	intersection := []int{0, 0}

	// start
	intersection[0] = max(firstList[0], secondList[0])

	// end
	intersection[1] = min(firstList[1], secondList[1])

	return intersection
}
