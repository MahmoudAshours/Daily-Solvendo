
package main

import "fmt"

func main() {
	nums := sortArrayByParityII([]int{1, 2, 4, 7, 3})
	for _, v := range nums {
		fmt.Println(v)
	}
}

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1 // i for even positions, j for odd positions

	for i < len(nums) && j < len(nums) {
		// If the number at even index is odd, swap it with the number at the odd index
		if nums[i]%2 == 1 && nums[j]%2 == 0 {
			swap(&nums[i], &nums[j])
		}
		// Move i to the next even index if it's correctly placed
		if nums[i]%2 == 0 {
			i += 2
		}
		// Move j to the next odd index if it's correctly placed
		if nums[j]%2 == 1 {
			j += 2
		}
	}
	return nums
}

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}
