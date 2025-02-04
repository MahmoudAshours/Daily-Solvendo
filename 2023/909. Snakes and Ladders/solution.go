package main

func main(){
	print(snakesAndLadders([][]int{{-1,-1,-1,-1,-1,-1},{-1,-1,-1,-1,-1,-1},{-1,-1,-1,-1,-1,-1},{-1,35,-1,-1,13,-1},{-1,-1,-1,-1,-1,-1},{-1,15,-1,-1,-1,-1}}))
}

func snakesAndLadders(board [][]int) int {
    n := len(board)
    need := make(map[int]int)
    need[1] = 0
    bfs := []int{1}
    for j := 0; j < len(bfs); j++ {
        x := bfs[j]
        for i := x + 1; i < x + 7; i++ {
            a, b := (i - 1) / n, (i - 1) % n
            nxt := 0
            if a % 2 == 0 {
                nxt = board[len(board)-a-1][b]
            } else {
                nxt = board[len(board)-a-1][len(board)-b-1]
            }
            temp := i
            if nxt > 0 { temp = nxt }
            if temp == n * n { return need[x] + 1 }
            if _, ok := need[temp]; !ok {
                need[temp] = need[x] + 1
                bfs = append(bfs, temp)
            }
        }
        
    }
    return -1
}