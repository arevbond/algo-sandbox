func removeDuplicates(nums []int) int {
        w := 1

        for r := 1; r < len(nums); r++ {
                if w >= len(nums) {
                        return w
                }

                for ; w < len(nums) && nums[w-1] < nums[w] ; w++ {}

                for ; r < len(nums) && nums[r-1] == nums[r]; r++ {}

                if r > w && r < len(nums) && w < len(nums) && nums[r] != nums[w] {
                	nums[w] = nums[r]
                	w++
                }

        }

        return w
}
