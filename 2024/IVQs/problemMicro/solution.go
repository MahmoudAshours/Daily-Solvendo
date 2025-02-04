package main

import "fmt"

func main() {
	// i j
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// 1 , 1 , 1 , 1 , 1
	// _|_|
	fmt.Println(trapWater(height))
}
func trapWater(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		left := height[i]
		right := height[i]
		for j := 0; j < i; j++ {
			left = max(left, height[j])
		}
		for j := i + 1; j < len(height); j++ {
			right = max(right, height[j])
		}
		res += min(left, right) - height[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**

**/
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
