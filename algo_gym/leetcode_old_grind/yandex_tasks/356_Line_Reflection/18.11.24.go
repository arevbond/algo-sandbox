package linereflection

//  *     *
//    * * 
//      
// minx   maxx
func IsReflected(points [][]int) bool {
    mp := make(map[[2]int]struct{})
    var minX, maxX int = math.MaxInt, -1 * math.MaxInt
    for _, coords := range points {
        x := coords[0]
        minX = min(minX, x)
        maxX = max(maxX, x)

        mp[[2]int{x, coords[1]}] = struct{}{}
    }

    mid := minX + maxX
    for _, coords := range points {
        x, y := coords[0], coords[1]
        if _, ok := mp[[2]int{mid-x, y}]; !ok {
            return false
        }
    } 
    return true
}
