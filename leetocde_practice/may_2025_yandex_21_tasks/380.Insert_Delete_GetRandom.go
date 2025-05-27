import "math/rand"

type RandomizedSet struct {
	nums []int    
	numToIndx map[int]int
}


func Constructor() RandomizedSet {
	return RandomizedSet{nums: make([]int, 0), numToIndx: make(map[int]int),} 
}


func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.numToIndx[val]; ok {
		return false
	}

	s.nums = append(s.nums, val)
	s.numToIndx[val] = len(s.nums) - 1	

	return true
}


func (s *RandomizedSet) Remove(val int) bool {
	indx, ok := s.numToIndx[val]
	if !ok {
		return false
	}

	lastDigit := s.nums[len(s.nums) - 1]

	s.nums[indx] = lastDigit
	s.numToIndx[lastDigit] = indx

	delete(s.numToIndx, val)
	s.nums = s.nums[:len(s.nums) - 1]

	return true
}


func (s *RandomizedSet) GetRandom() int {
	if len(s.nums) == 0 {
		return -1
	}

	return s.nums[rand.Intn(len(s.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
