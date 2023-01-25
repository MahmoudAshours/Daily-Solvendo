package main

import "math"

func main() {
	print(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
}
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	dist1 := make([]int, n)
	dist2 := make([]int, n)
	for i := 0; i < n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}

	traverse(node1, &edges, 0, &dist1)
	traverse(node2, &edges, 0, &dist2)
	maxDis := math.MaxInt32
	label := -1
	for i := 0; i < n; i++ {
		// node i is reachable from both node1 and node2
		if dist1[i] > -1 && dist2[i] > -1 {
			mx := max(dist1[i], dist2[i])
			if mx < maxDis {
				maxDis = mx
				label = i
			}
		}
	}

	return label
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

func traverse(curr int, edges *[]int, dis int, dist *[]int) {
	if (*dist)[curr] != -1 {
		return
	}

	(*dist)[curr] = dis
	if (*edges)[curr] == -1 {
		return
	}

	traverse((*edges)[curr], edges, dis+1, dist)
}
