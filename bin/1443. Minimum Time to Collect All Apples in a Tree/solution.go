package main

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	hasApple := []bool{false, false, true, false, true, true, false}
	minTime(8, edges[:], hasApple[:])
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	edgesRec := make([][]int, n)
	for _, element := range edges {
		edgesRec[element[0]] = append(edgesRec[element[0]], element[1])
		edgesRec[element[1]] = append(edgesRec[element[1]], element[0])
	}
	visited := make([]bool, n)
	visited[0] = true
	return depthFirstSearch(0, edgesRec, hasApple, visited)
}

func depthFirstSearch(current int, edgesRec [][]int, hasApple []bool, visited []bool) (ret int) {
	appleIncrease := 0
	for _, adj := range edgesRec[current] {
		if visited[adj] {
			continue
		}
		visited[adj] = true
		ret += depthFirstSearch(adj, edgesRec, hasApple, visited)
		visited[adj] = false
	}
	if (hasApple[current] && current != 0) || (ret > 0 && current != 0) {
		appleIncrease = 2
	}
	ret += appleIncrease
	return
}
