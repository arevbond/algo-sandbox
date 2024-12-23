package robotreturntoorigin

func judgeCircle(moves string) bool {
    coords := [2]int{0, 0}

    for _, move := range moves {
        switch move {
        case 'U':
            coords[1]++
        case 'D':
            coords[1]--
        case 'L':
            coords[0]--
        case 'R':
            coords[0]++
        }
    }

    return coords == [2]int{0, 0}
}
