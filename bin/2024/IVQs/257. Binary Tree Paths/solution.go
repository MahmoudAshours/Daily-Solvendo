package main

import "strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	s := make([]string, 0)
	traverse(root, &s, "")
	return s
}

func traverse(root *TreeNode, arr *[]string, str string) {
	if root == nil {
		return
	}

	str += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, str)
		return
	}

	traverse(root.Left, arr, str+"->")
	traverse(root.Right, arr, str+"->")

}
