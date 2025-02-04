func dfs(curr, parent int, adjList [][]int, ans *int64, seats int) int {
    var numOfRepresentative, total = 0, 0
    for _, next := range adjList[curr] {
        if next == parent {continue}
        numOfRepresentative = dfs(next, curr, adjList, ans, seats)
        *ans += int64(math.Ceil(float64(numOfRepresentative) / float64(seats)))
        total += numOfRepresentative
    }
    return total + 1
}

func minimumFuelCost(roads [][]int, seats int) int64 {
    var n int = len(roads)
    adjList := make([][]int, n + 1)
    for _, edge := range roads {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0]);
    }
    var ans int64 = 0
    dfs(0, -1, adjList, &ans, seats)
    return ans
}