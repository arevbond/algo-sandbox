package removeelement

func removeElement(nums []int, val int) int {
    indxToWrite := 0
    
    for r := range nums {
        for ; indxToWrite < len(nums) && nums[indxToWrite] != val; indxToWrite++{}

        if indxToWrite < len(nums) && r > indxToWrite && nums[r] != val {
            nums[indxToWrite], nums[r] = nums[r], nums[indxToWrite]
        }
    }

    if indxToWrite < len(nums) && nums[indxToWrite] != val {
        indxToWrite++
    }

    return indxToWrite
}

func removeElement2(nums []int, val int) int {
    l := 0
    for r := 0; r < len(nums); r++ {
        if l < len(nums) && l < r && nums[l] == val && nums[l] != nums[r] {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }
        for ; l < len(nums) && nums[l] != val; l++ {}
    }
    return l
}
