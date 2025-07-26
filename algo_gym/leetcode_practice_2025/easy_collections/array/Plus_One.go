func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if digit + carry <= 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
			carry = 1
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
