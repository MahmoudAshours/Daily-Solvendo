package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var p TreeNode
	var q TreeNode
	p = TreeNode{
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	q = TreeNode{
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	print(isSameTree(&p, &q))
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
