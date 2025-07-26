// correct, but invalid
func reverse1(x int) int {
	isNegative := x < 0
	x = abs(x)
	n := 1
	k := x
	for k > 9 {
		n *= 10
		k /= 10
	}

	result := 0

	pow := 0.0
	for n > 0 {
		digit := x / n
		number := digit * int(math.Pow(10, pow))
		if result + number > math.MaxInt32 {
			return 0
		}
		result += number
		x=abs(((x / n)*n) - x)
		n = n/10
		pow++
	}
	
	if isNegative {
		return -1 * result
	}
	return result
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

// incorrect
func reverse(x int) int {
	isNegative := x < 0
	x = abs(x)
	n := 1
	k := x
	for k > 9 {
		n *= 10
		k /= 10
	}

	result := int32(0)

	pow := 0.0
	for n > 0 {
		digit := x / n
		number := int32(digit) * int32(math.Pow(10, pow))
		if result + number > math.MaxInt32 {
			return 0
		}
		result += number
		x=abs(((x / n)*n) - x)
		n = n/10
		pow++
	}
	
	if isNegative {
		return -1 * int(result)
	}
	return int(result)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
