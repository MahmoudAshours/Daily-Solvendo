package main

import "math"

func main() {
	nums := [][]int{{2, 3}}
	print(nearestValidPoint(3, 4, nums[:]))
}
func nearestValidPoint(x int, y int, points [][]int) int {
	smallestDistance := math.MaxUint32
	smallestIndex := math.MaxUint32

	for i, v := range points {
		if x == v[0] || y == v[1] {
			manhattaenDistance := math.Abs(float64(x-v[0])) + math.Abs(float64(y-v[1]))
			if manhattaenDistance < float64(smallestDistance) || i < smallestIndex {
				smallestDistance = int(manhattaenDistance)
				smallestIndex = i
			}
		}
	}
	if smallestIndex == math.MaxUint32 {
		return -1
	}
	return smallestIndex
}
