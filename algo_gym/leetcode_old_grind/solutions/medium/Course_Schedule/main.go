package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for _, data := range prerequisites {
		crs, pre := data[0], data[1]
		preMap[crs] = append(preMap[crs], pre)
	}
	visitedMap := make(map[int]bool)
	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if _, ok := visitedMap[crs]; ok {
			return false
		}
		if len(preMap[crs]) == 0 {
			return true
		}
		visitedMap[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visitedMap, crs)
		preMap[crs] = []int{}
		return true
	}
	for crs := 0; crs < numCourses; crs++ {
		if !dfs(crs) {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println([]int{} == nil)
	fmt.Println(canFinish(5, [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 3},
		[]int{1, 4}, []int{3, 4}}))
}
