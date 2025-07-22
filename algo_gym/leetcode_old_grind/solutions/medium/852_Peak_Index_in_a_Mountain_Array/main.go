package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
    l, r := 0, len(arr) - 1
    onLeft := func(indx int) bool {
    	return arr[indx] < arr[indx + 1]
    }

    for l < r {
    	m := (l + r) / 2
    	// fmt.Println(m)
    	if onLeft(m) {
    		l = m + 1
    	} else {
    		r = m
    	}
    }

    return l
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0,1,0}))
	fmt.Println(peakIndexInMountainArray([]int{0,2,1,0}))
	fmt.Println(peakIndexInMountainArray([]int{0,10,5,2}))
	fmt.Println(peakIndexInMountainArray([]int{0,2,4,10,5,3}))
	fmt.Println(peakIndexInMountainArray([]int{24,69,100,99,79,78,67,36,26,19}))
}