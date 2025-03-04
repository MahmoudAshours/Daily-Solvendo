package main

import "fmt"

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
}
func validPath(n int, edges [][]int, start int, end int) bool {
    m := make(map[int] []int)
    
    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }
    
    alreadyDid := make(map[int] bool) // preventing infinite loops
    stack := []int{ start }
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if pop == end {
            return true
        }
        
        if !alreadyDid[pop] {
            alreadyDid[pop] = true
            stack = append(stack, m[pop]...)
        }
    }
    
    return false
}

