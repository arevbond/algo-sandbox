package main

import (
    "math/rand"
)

type Solution struct {
    nums []int
}

func Constructor(nums []int) Solution {
    return Solution{
        nums: nums,
    }
}

func (s *Solution) Shuffle() []int {
    usedIndices := make(map[int]struct{}, len(s.nums))

    result := make([]int, 0, len(s.nums))

    for len(usedIndices) != len(s.nums) {
        indx := rand.Intn(len(s.nums))

        if _, ok := usedIndices[indx]; ok {
            continue
        }

        usedIndices[indx] = struct{}{}

        result = append(result, s.nums[indx])
    }

    return result
}

func (s *Solution) Reset() []int {
    return s.nums
}
