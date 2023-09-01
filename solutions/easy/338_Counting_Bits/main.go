package main 

import (
	"fmt"
	"strconv"
)

func countBits(n int) []int {
    result := []int{}
    for x := 0; x <= n; x++ {
    	cnt := countBit(x)
    	result = append(result, cnt)
    }
    return result
}

func countBit(x int) int {
	var ans int
	binaryX := strconv.FormatInt(int64(x), 2)
	for _, ch := range []rune(binaryX) {
		if ch == '1' {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}