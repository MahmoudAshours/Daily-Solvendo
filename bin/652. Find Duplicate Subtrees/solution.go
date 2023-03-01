package main
func main(){
	
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    seen := make(map[[3]int]*[2]int) //map[[3]int{Val, leftId, rightId}] [2]int{id, timesSeen}
    ans := make([]*TreeNode, 0)

    id_count := 1

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int { // returns the tree id
        var id int
        if root == nil {
            return id
        }

        left := dfs(root.Left)
        right := dfs(root.Right)

        key := [3]int{root.Val, left, right}
        val, ok := seen[key]
        if !ok { // not seen
            id = id_count
            seen[key] = &[2]int{id, 1}
            id_count++
            return id
        } 

        id = val[0]
        if val[1] == 1 { // seen once
            ans = append(ans, root)
            seen[key][1]++
        }

        return id
    }

    dfs(root)
    return ans
}