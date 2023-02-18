package main



type TreeNode struct {
       Val int
       Left *TreeNode
       Right *TreeNode
}
 
func main(){
	node := TreeNode{
		Val:4,
	}
	node.Left = &TreeNode{
		Val:2,
		Left : &TreeNode{
			Val:1,
		},
		Right: &TreeNode{
			Val : 3,
		},
	}
	node.Right = &TreeNode{
		Val:7,
		Left : &TreeNode{
			Val:6,
		},
		Right: &TreeNode{
			Val : 9,
		},
	}
	invertTree(&node)
}


func invertTree(root *TreeNode) *TreeNode {
	 if root != nil {
        root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    }
    return root
}









