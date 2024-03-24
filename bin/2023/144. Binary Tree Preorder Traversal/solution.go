package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var node TreeNode
	node.Val = 1
	preorderTraversal(&node)
}

/**
 * Definition for a binary tree node.
 */

func preorderTraversal(root *TreeNode) []int {
	// Create an empty slice to store the result.
	result := make([]int, 0)

	// Return an empty slice if the root is nil.
	if root == nil {
		return result
	}

	// Create a stack to store the nodes we have to visit.
	stack := make([]*TreeNode, 0)

	// Push the root node onto the stack.
	stack = append(stack, root)

	// Keep visiting nodes until the stack is empty.
	for len(stack) > 0 {
		// Pop the top node from the stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Append the node's value to the result slice.
		result = append(result, node.Val)

		// Push the right child onto the stack.
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// Push the left child onto the stack.
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
