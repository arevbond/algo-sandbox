type TimeMap struct {
    values map[string][]string
    times map[string][]int
}


func Constructor() TimeMap {
	return TimeMap{values: make(map[string][]string),
					times: make(map[string][]int)}    
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.values[key] = append(this.values[key], value)
    this.times[key] = append(this.times[key], timestamp)
}


func findIndex(times []int, time int) int {
	l, r := 0, len(times) - 1
	for l < r {
		m := (l + r) / 2
		if times[m] >= time {
			r = m
		} else {
			l = m + 1
		}
	}
	if times[l] <= time {
		return l
	}
	return l - 1
}

func (this *TimeMap) Get(key string, timestamp int) string {
    times, ok := this.times[key]
    if !ok {
    	return ""
    }
    indx := findIndex(times, timestamp)
    if indx == -1 {
    	return ""
    }
		values := this.values[key]
    return values[indx]
}	