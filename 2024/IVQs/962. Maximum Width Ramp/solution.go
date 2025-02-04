package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}

func maxWidthRamp(nums []int) int {
	n := len(nums)
	stack := []int{}

	// Step 1: Build a decreasing stack of indices
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	maxWidth := 0

	// Step 2: Traverse from the end and find maximum width ramp
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if j-i > maxWidth {
				maxWidth = j - i
			}
		}
	}

	return maxWidth
}

// 5 1 2 3 4

/*
5 0
1 1
2 2
3 3
4 4
*/

/*
2 0
2 1
2 2
2 3
1 2
*/
