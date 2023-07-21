package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
    answer := make([]int, len(temperatures))
    for i := 0; i < len(answer); i++ {
    	answer[i] = 0
    }

    

    return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}