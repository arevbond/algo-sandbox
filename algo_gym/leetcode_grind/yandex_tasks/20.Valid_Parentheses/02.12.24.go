package validparentheses

/*

    ()(){()}     (({}))
    
    stack = [(({] 
    1. for ch range s 
    2. if ch is opening => append this to stack
    3. if ch is closing:
    4. check stack not empty
    5. last element in stack is pair for current bracket

    pairs: {'(': ')', ...}

*/
func isValid(s string) bool {
    openToClose := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    stack := make([]rune, 0)

    for _, ch := range s {
        if _, ok := openToClose[ch]; !ok {
            if len(stack) == 0 {
                return false
            }
            lastOpenBracket := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            if ch != openToClose[lastOpenBracket] {
                return false
            } 
        } else {
            stack = append(stack, ch)
        }
    }

    return len(stack) == 0
}
