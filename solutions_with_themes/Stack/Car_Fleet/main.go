package main

import (
	"fmt"
	"sort"
	// "math"
)

func carFleet(target int, position []int, speed []int) int {
    answer := 0

    var prevTime float64 = 0
    cars := make([][]int, len(position))
    for i := 0; i < len(cars); i++ {
    	cars[i] = []int{position[i], speed[i]}
    }
    sort.Slice(cars, func(i, j int) bool {
    	return cars[i][0] < cars[j][0]
    })
    for i := len(cars) - 1; i >= 0; i-- {
    	time := float64(target - cars[i][0]) / float64(cars[i][1])
    	if time > prevTime {
    		answer++
    		prevTime = time
    	} else if time <= prevTime {
    		continue
    	}
    }
    return answer

}

func main() {
	fmt.Println(carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3}))
}