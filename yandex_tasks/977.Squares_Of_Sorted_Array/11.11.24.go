package squaresofasortedarray

// input: nums = [-4, -1, 0, 3, 10]
// output: [0. 1, 9, 16, 100] 

// input: nums = [-7, -3, 2, 3, 11]
// output: [4, 9, 9, 49, 121] 


func sortedSquares(nums []int) []int {
    negSquars := make([]int, 0)    
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < 0 {
            negSquars = append(negSquars, nums[i] * nums[i])
        }
    }

    posSquares := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] >= 0 {
            posSquares = append(posSquares, nums[i] * nums[i])
        }
    }

    result := make([]int, 0, len(negSquars) + len(posSquares))
    for i, j := 0, 0; i < len(negSquars) || j < len(posSquares); {
        if i == len(negSquars) {
            result = append(result, posSquares...)
            break
        }

        if j == len(posSquares) {
            result = append(result, negSquars...)
            break
        }

        neg, pos := negSquars[i], posSquares[j]
        if neg > pos {
            result = append(result, neg)
            i++
        } else {
            result = append(result, pos)
            j++
        }
    }
    return result
}
