package main

import "container/list"

func main() {

}

type Pair struct {
	First  int
	Second *TreeNode
	Sum    int
}

func solve(root *TreeNode) {
	if root == nil {
		return
	}

	q := list.New()
	cpy := list.New()
	levelSum := []int{}
	level := 0

	levelSum = append(levelSum, root.Val)
	q.PushBack(Pair{level, root, root.Val})
	cpy.PushBack(Pair{level, root, root.Val})

	for q.Len() > 0 {
		n := q.Len()
		lvlSum := 0

		for n > 0 {
			front := q.Remove(q.Front()).(Pair)
			node := front.Second
			sibSum := 0

			if node.Left != nil {
				sibSum += node.Left.Val
			}
			if node.Right != nil {
				sibSum += node.Right.Val
			}

			if node.Left != nil {
				q.PushBack(Pair{level + 1, node.Left, sibSum})
				cpy.PushBack(Pair{level + 1, node.Left, sibSum})
			}

			if node.Right != nil {
				q.PushBack(Pair{level + 1, node.Right, sibSum})
				cpy.PushBack(Pair{level + 1, node.Right, sibSum})
			}

			lvlSum += sibSum
			n--
		}

		levelSum = append(levelSum, lvlSum)
		level++
	}

	for cpy.Len() > 0 {
		front := cpy.Remove(cpy.Front()).(Pair)
		lvl := front.First
		node := front.Second
		sibSum := front.Sum

		node.Val = levelSum[lvl] - sibSum
	}
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	solve(root)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
