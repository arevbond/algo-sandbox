type Node struct {
    parentIndx int
    node *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
    stack := []*Node{&Node{parentIndx: 0, node: root}}
    familySum := root.Val // sum node values in same depth 

    for len(stack) > 0 {
        newStack := make([]*Node, 0)
        indx := 0
        nextFamilySum := 0

        for i := 0; i < len(stack); {
            curNode := stack[i]
            curSum := curNode.node.Val

            if curNode.node.Left != nil {
                newStack = append(newStack, &Node{
                    parentIndx: indx,
                    node: curNode.node.Left,
                })
                nextFamilySum += curNode.node.Left.Val
            }

            if curNode.node.Right != nil {
                newStack = append(newStack, &Node{
                    parentIndx: indx,
                    node: curNode.node.Right,
                })
                nextFamilySum += curNode.node.Right.Val
            }

            indx++

            if i + 1 < len(stack) {
                nextNode := stack[i+1]

                if curNode.parentIndx == nextNode.parentIndx {
                    if nextNode.node.Left != nil {
                        newStack = append(newStack, &Node{
                            parentIndx: indx,
                            node: nextNode.node.Left, 
                        })
                        nextFamilySum += nextNode.node.Left.Val
                    }

                    if nextNode.node.Right != nil {
                        newStack = append(newStack, &Node{
                            parentIndx: indx,
                            node: nextNode.node.Right,
                        })
                        nextFamilySum += nextNode.node.Right.Val
                    }

                    indx++

                    curSum += nextNode.node.Val

                    result := familySum - curSum
              
                    curNode.node.Val = result
                    nextNode.node.Val = result

                    i += 2
                    continue
                } 
            } 
            result := familySum - curSum
            curNode.node.Val = result
            i++
        }

        stack = newStack
        familySum = nextFamilySum
    }

    return root
}
