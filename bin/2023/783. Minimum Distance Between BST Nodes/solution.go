package main

func main(){

}

 
  type TreeNode struct {
      Val int
       Left *TreeNode
       Right *TreeNode
   }
 
func minDiffInBST(root *TreeNode) int {
    minDiff, preValue := math.MaxInt32, -1
    dfs(root, &minDiff, &preValue)
    return minDiff
}

func dfs(node *TreeNode, minDiff *int, preValue *int) {
    if node == nil {
        return 
    }

    dfs(node.Left, minDiff, preValue)

    if *preValue != -1 && node.Val - *preValue < *minDiff {
        *minDiff = node.Val - *preValue
    }
    *preValue = node.Val

    dfs(node.Right, minDiff, preValue)
}