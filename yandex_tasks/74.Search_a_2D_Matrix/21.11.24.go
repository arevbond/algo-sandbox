package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
    R := len(matrix)
    C := len(matrix[0])

    good := func(indx int) bool {
        x, y := indx / C, indx % C
        return matrix[x][y] >= target
    }

    left, right := 0, R * C 
    
    for left < right {
        mid := (left + right) / 2
        if good(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    if left == R * C {
        return false
    }
    x, y := left / C, left % C
    return matrix[x][y] == target
}
