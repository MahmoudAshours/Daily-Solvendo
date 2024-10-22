package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(kthLargestLevelSum(createBST([]int{5, 8, 9, 2, 1, 3, 7, 4, 6}), 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		// If the tree is empty, create a new node with the value
		return &TreeNode{Val: val}
	}

	// Insert the value into the left or right subtree depending on its value
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

// Function to create a BST from an array
func createBST(arr []int) *TreeNode {
	var root *TreeNode
	for _, val := range arr {
		root = insert(root, val)
	}
	return root
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	arr := getNodesAtSameLevel(root)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[k]

}
func getNodesAtSameLevel(root *TreeNode) []int64 {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []int64{}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}
		levelSum := 0
		// Process all nodes in the current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Collect values of nodes at the current level
			currentLevel = append(currentLevel, node.Val)

			// Enqueue left and right children if they exist
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		fmt.Println(currentLevel)
		for _, v := range currentLevel {
			levelSum += v
		}
		result = append(result, int64(levelSum))
	}

	return result
}
