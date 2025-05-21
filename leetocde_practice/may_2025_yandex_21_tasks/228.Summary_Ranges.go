func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	ranges := make([]string, 0)

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] + 1 {
			end = nums[i]
		} else {
			ranges = append(ranges, convertRange(start, end))

			start, end = nums[i], nums[i]
		}
	}

	ranges = append(ranges, convertRange(start, end))

	return ranges
}

func convertRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}

	return fmt.Sprintf("%d->%d", start, end)
}
