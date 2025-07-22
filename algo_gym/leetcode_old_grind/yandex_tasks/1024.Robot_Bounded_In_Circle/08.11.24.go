package robotboundedincircle

func isRobotBounded(instructions string) bool {
    dirs := []rune{'N', 'E', 'S', 'W'}
    dirIndx := 0
    coords := [2]int{0, 0}

    for i := 0; i < 2; i++ {
        instructions += instructions
    }

    for _, instruction := range instructions {
        switch instruction {
        case 'G':
            coords = proccessMove(dirs[dirIndx], coords)
        case 'L':
            dirIndx--
            if dirIndx < 0 {
                dirIndx = len(dirs) - 1
            }
        case 'R':
            dirIndx++
            if dirIndx == len(dirs) {
                dirIndx = 0
            }
        }
    }
    return coords == [2]int{0, 0}
}

func proccessMove(dir rune, coords [2]int) [2]int {
    switch dir {
    case 'N':
        coords[1]++
    case 'S':
        coords[1]--
    case 'E':
        coords[0]++
    case 'W':
        coords[0]--
    }
    return coords
}
