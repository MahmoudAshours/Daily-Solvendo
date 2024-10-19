package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findLucky([]int{1, 2, 2, 3}))
}
func findLucky(arr []int) int {
	mp := make(map[int]int, 0)
	max := -1
	for _, v := range arr {
		mp[v]++
	}
	for k, v := range mp {
		if k == v {
			max = int(math.Max(float64(k), float64(max)))
		}
	}
	return max
}
