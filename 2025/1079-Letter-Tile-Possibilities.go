/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var sum int
func sumRootToLeaf(root *TreeNode) int {
    sum = 0
    traverse(root,[]int{})
    return sum
}

func traverse(root *TreeNode,arr []int){
if root.Left == nil && root.Right == nil{
    sum = sum + binToDec(append(arr,root.Val))
    return
} 
if root.Left != nil{
    traverse(root.Left,append(arr,root.Val))
}
if root.Right != nil{
    traverse(root.Right,append(arr,root.Val))
}
}

func binToDec(arr []int) int{
    f := 1
    s := 0
    for i := len(arr) - 1 ; i >=0 ; i-- {
        s = s + arr[i] * f
        f = f * 2
    }
    return s
}