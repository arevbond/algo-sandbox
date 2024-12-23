package partitionlabels

func partitionLabels(s string) []int {
    indxMap := make(map[rune]int)
    for i, ch := range s {
        indxMap[ch] = i
    }
    
    result := make([]int, 0)
    lastIndx, start := 0, 0
    for i, ch := range s {
        lastIndx = max(lastIndx, indxMap[ch])
        if i == lastIndx {
            result = append(result, i-start+1)
            start = i + 1
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
