package removeinvalidparentheses

import ("math")

func removeInvalidParentheses(s string) []string {
    result := make(map[int]map[string]struct{})
//    result := map[int]map[string]struct{}{}

    var cntBrackets int
    for _, ch := range s {
        if ch == '(' || ch == ')' {
            cntBrackets++
        }
    }

    runes := []rune(s)
    var backtrack func(indx int, openN, closeN int, cur string)
    backtrack = func(indx int, openN, closeN int, cur string) {
        if indx == len(runes) {
            if openN == closeN {
                removals := cntBrackets - (openN + closeN)
                if _, ok := result[removals]; !ok {
                    result[removals] = make(map[string]struct{})
                }
                result[removals][cur] = struct{}{}
                return
            } 
        }

        if openN < closeN {
            return
        }

        for i := indx; i < len(runes); i++ {
            if runes[i] == '(' {
                backtrack(i+1, openN, closeN, cur)
                backtrack(i+1, openN+1, closeN, cur + string(runes[i]))
            } else if runes[i] == ')' {
                backtrack(i+1, openN, closeN, cur)
                backtrack(i+1, openN, closeN+1, cur + string(runes[i]))
            } else {
                backtrack(i+1, openN, closeN, cur + string(runes[i]))
            }
        }
        return
    }

    backtrack(0, 0, 0, "")

    minRemovals := math.MaxInt
    for removals := range result {
        minRemovals = min(minRemovals, removals)
    }

    ans := make([]string, 0, len(result[minRemovals]))
    for str := range result[minRemovals] {
        ans = append(ans, str)
    }

    return ans 
}

func min(a, b int) int {
    if a > b {
        return a
    }
    return b
}
