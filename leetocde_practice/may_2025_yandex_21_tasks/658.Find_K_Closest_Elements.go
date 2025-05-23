// sorting - O(n*logN) time
func findClosestElementsSORT(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if abs(a - x) == abs(b -x) {
			return a < b
		}
		return abs(a - x) < abs(b - x)
	})

	slices.Sort(arr[:k])

	return arr[:k]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

// 1. find most closest number in arr
// 2. pick l and r pointers around most closest number
// 3. compare arr[l] and arr[r] for closeness
func findClosestElements(arr []int, k int, x int) []int {
	closestIndx := 0

	for i, num := range arr {
		closestNum := arr[closestIndx]

		if abs(closestNum-x) == abs(num-x) {
			if num < closestNum {
				closestIndx = i
			}
		} else {
			if abs(num-x) < abs(closestNum-x) {
				closestIndx = i
			}
		}
	}

	l := closestIndx 
	r := closestIndx + 1
	k = k - 1
	for k > 0 {
		if (l - 1) < 0 {
			r++
			k--
			continue
		} else if r == len(arr) {
			l--
			k--
			continue
		}

		if abs(arr[l-1]-x) == abs(arr[r]-x) {
			if arr[l-1] < arr[r] {
				l--
			} else {
				r++
			}
		} else {
			if abs(arr[l-1]-x) < abs(arr[r]-x) {
				l--
			} else {
				r++
			}
		}
		k--
	}

	return arr[l:r]
}

