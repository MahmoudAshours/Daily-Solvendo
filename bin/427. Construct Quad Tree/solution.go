package main
func main(){
	construct([][]int{{1,1}})
}


  type Node struct {
       Val bool
       IsLeaf bool
       TopLeft *Node
       TopRight *Node
       BottomLeft *Node
       BottomRight *Node
   }
 

func construct(grid [][]int) *Node {
    return dfs(grid)
}

func dfs(grid [][]int) *Node {
    head := &Node{}
    prev := grid[0][0]
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid); j++ {
            if grid[i][j] != prev {
                // If the grid have both 0 and 1,
                // we split it into 4 subgrid
                mid := len(grid) / 2
                topLeft := make([][]int, 0, mid)
                topRight := make([][]int, 0, mid)
                bottomLeft := make([][]int, 0, mid)
                 bottomRight := make([][]int, 0, mid)
                for k := 0; k < mid; k++ {
                    topLeft = append(topLeft, grid[k][:mid])
                    topRight = append(topRight, grid[k][mid:])
                    bottomLeft = append(bottomLeft, grid[mid+k][:mid])
                    bottomRight = append(bottomRight, grid[mid+k][mid:])
                }
                head.TopLeft = dfs(topLeft)
                head.TopRight = dfs(topRight)
                head.BottomLeft = dfs(bottomLeft)
                head.BottomRight = dfs(bottomRight)
                return head
            }
        }
    }

    // If the grid consists of only one value,
    // it will be a leaf node
    if prev == 1 {
        head.Val = true
    } else {
        head.Val = false
    }
    head.IsLeaf = true
    return head
}