package main

func firstBadVersion(n int) int {
    l, r := 1, n

    for l < r {
        m := (l + r) / 2
        if isBadVersion(m) {
            r = m
        } else {
            l = m + 1
        }
    }

    return l
}
