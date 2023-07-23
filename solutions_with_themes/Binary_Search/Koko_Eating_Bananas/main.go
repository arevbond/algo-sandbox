package main

import (
	"fmt"
	"math"
)

func Round(x float64) int {
	ans := math.Trunc(x)
	if x - ans > 0 {
		ans++
	} 
	return int(ans)
}

func hasInTime(piles []int, h, k int) bool {
	time := 0
	for _, pile := range piles {
		time += Round(float64(pile)/float64(k))
	}
	return time <= h
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 10000000000
	for l < r {
		m := (l + r) / 2
		if hasInTime(piles, h, m) {
			r = m
		} else {
			l = m +1
		}
	}
    return l
}


func main() {
	// fmt.Println(Round(1.2))
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5))
	fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6))
}