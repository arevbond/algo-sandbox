package main

import "fmt"

func frequencyDictionry(str string) map[string]int {
    counter := make(map[string]int)

    runes := []rune(str)

    for l := 0; l < len(runes); {
        var cnt int
        cur := runes[l]
        
        r := l
        for ; r < len(runes) && runes[r] == cur; r++ {
            cnt++
        }
        counter[string(cur)] = max(cnt, counter[string(cur)])
        l = r
    }

    return counter
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(frequencyDictionry("aafbaaaaffc"))
    fmt.Println(frequencyDictionry("aaaaaa"))
    fmt.Println(frequencyDictionry("abc"))
}
