/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	    
}

func isSameTree(rootOne *TreeNode, rootTwo *TreeNode) bool {
	if rootOne == nil && rootTwo == nil {
		return true
	}

	if rootOne == nil || rootTwo == nil {
		return false
	}

	if rootOne.Val != rootTwo.Val {
		return false
	}

	return isSameTree(rootOne.Left, rootTwo.Left) && isSameTree(rootOne.Right, rootTwo.Right)
}
