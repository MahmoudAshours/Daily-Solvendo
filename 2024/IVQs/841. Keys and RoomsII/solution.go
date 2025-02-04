package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}

func canVisitAllRooms(rooms [][]int) bool {

	adjList := make(map[int][]int, 0)
	for i, v := range rooms {
		for _, room := range v {
			adjList[i] = append(adjList[i], room)
		}
	}
	return bfs(adjList)

}

func bfs(adjlist map[int][]int) bool {
	queue := make([]int, 0)
	visited := make(map[int]bool, len(adjlist))
	visited[0] = true

	queue = append(queue, 0)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, v := range adjlist[curr] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
	i := 0
	for i < len(adjlist) {
		if !visited[i] {
			return false
		}
		i++
	}
	return true
}
