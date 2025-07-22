package tasks

// Time: O(N) Space: O(1)
func backspaceCompare(s string, t string) bool {
	sI := nextValidChar(s, len(s))
	tI := nextValidChar(t, len(t))

	for sI >= 0 || tI >= 0 {
		if sI < 0 || tI < 0 {
			return false
		}

		if s[sI] != t[tI] {
			return false
		}

		sI = nextValidChar(s, sI)
		tI = nextValidChar(t, tI)
	}

	return true
}

func nextValidChar(s string, indx int) int {
	shifts := 0

	for i := indx - 1; i >= 0; i-- {
		if s[i] == '#' {
			shifts++
		} else  if shifts > 0 {
			shifts--
		} else {
			return i
		}
	}

	return -1
}

// Time: O(N), Space: O(N)
func backspaceCompare_O-N(s string, t string) bool {
	return processBackspace(s) == processBackspace(t)	
}

func processBackspace(s string) string {
	stack := []rune{}

	for _, ch := range s {
		if ch == '#' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if ch != '#' {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}
