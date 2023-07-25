package main

func checkInclusion(s1 string, s2 string) bool {
    l, count := 0, [26]int{}
    for i := range s1 {
    	count[s1[i]-97]++
    }

    for r := range s2 {
    	count[s2[r]-97]--
    	if count == [26]int{} {
    		return true
    	}
    	if r + 1 >= len(s1) {
    		count[s2[l]-97]++
    		l++
    	}
    }
    return false
}

