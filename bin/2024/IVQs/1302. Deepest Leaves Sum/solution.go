package main

func main() {
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	maxHeight := getMaxHeight(root)

	sum := 0
	getDeepest(root, 1, maxHeight, &sum)

	return sum
}

func getDeepest(root *TreeNode, currHeight int, maxHeight int, sum *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && currHeight == maxHeight {
		*sum = *sum + root.Val
	}

	getDeepest(root.Left, currHeight+1, maxHeight, sum)
	getDeepest(root.Right, currHeight+1, maxHeight, sum)
}

func getMaxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getMaxHeight(root.Left), getMaxHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
