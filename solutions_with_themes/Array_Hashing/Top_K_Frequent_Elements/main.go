package main

import (
	"sort"
)

type kv struct {
	k, v int
}

func topKFrequent(nums []int, k int) []int {
	mapInts := make(map[int]int)
	for _, num := range nums {
		mapInts[num]++
	}
	kvs := []kv{}
	for k, v := range mapInts {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v > kvs[j].v
	})
	var answer []int
	for i := 0; i < k; i++ {
		answer = append(answer, kvs[i].k)
	}
	return answer
}
