package main

func main(){
shortestAlternatingPaths(3,[][]int{{0,1},{1,2}})
}
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) 
[]int {
    redCities := make([][]int, n) 
    blueCities := make([][]int, n)
    for _, edge := range redEdges {
        redCities[edge[0]] = append(redCities[edge[0]], edge[1])
    }
    for _, edge := range blueEdges {
        blueCities[edge[0]] = append(blueCities[edge[0]], edge[1])
    }
    redDist := make([]int, n)
    blueDist := make([]int, n)
    for i := 0; i < n; i++ {
        redDist[i] = math.MaxInt
        blueDist[i] = math.MaxInt
    }
    redDist[0], blueDist[0] = 0, 0
    Queue := [][]int{}
    Queue = append(Queue, []int{0, 'r'})
    Queue = append(Queue, []int{0, 'b'})
    for len(Queue) > 0 {
        curr := Queue[0]
        Queue = Queue[1:]
        if curr[1] == 'r' {
            for _, next := range blueCities[curr[0]] {
                if blueDist[next] == math.MaxInt {
                    blueDist[next] = redDist[curr[0]] + 1
                    Queue = append(Queue, []int{next, 'b'})
                }
            }
        } else {
            for _, next := range redCities[curr[0]] {
                if redDist[next] == math.MaxInt {
                    redDist[next] = blueDist[curr[0]] + 1
                    Queue = append(Queue, []int{next, 'r'})
                }
            }
        }
    }
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = redDist[i]
        if blueDist[i] < ans[i] {
            ans[i] = blueDist[i]
        }
        if ans[i] == math.MaxInt {
            ans[i] = -1
        }
    }
    return ans
}
