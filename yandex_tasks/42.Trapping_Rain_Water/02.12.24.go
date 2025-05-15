package trappingrainwater

/*
    [0,1,0,2,1,0,1,3,2,1,2,1]

    1. find max height => this will be end for two loops
    2. var leftMaxHeight; first loop left -> maxHeight:
    3. max(leftMaxHeight); if curHeight < leftMaxHeight; pour water leftMax - cur
    4. var rightMaxHeight; second loop maxHeight <- right:
    5. max(rightMaxHeight); if curHeight < rightMaxHeight: pour water rightMax - cur

*/

func trap(height []int) int {
    var maxHeight int
    var indxMaxHeight int
    for i, h := range height {
        if h > maxHeight {
            maxHeight = h
            indxMaxHeight = i
        }
    }

    var water int

    var leftMaxHeight int
    for i := 0; i < indxMaxHeight; i++ {
        curHeight := height[i]
        leftMaxHeight = max(leftMaxHeight, curHeight)
        if curHeight < leftMaxHeight {
            water += leftMaxHeight - curHeight
        }
    }

    var rightMaxHeight int
    for i := len(height) - 1; i > indxMaxHeight; i-- {
        curHeight := height[i]
        rightMaxHeight = max(rightMaxHeight, curHeight)
        if curHeight < leftMaxHeight {
            water += min(leftMaxHeight, maxHeight)
        }
    }
    return water
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
