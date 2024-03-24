package main

import "sort"

func main() {
	nums := []int{3, 6, 2, 3}
	print(largestPerimeter(nums))
}
func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	for index := range nums {
		if index != len(nums)-2 {
			if nums[index+2]+nums[index+1] <= nums[index] {
				continue
			} else {
				return nums[index] + nums[index+1] + nums[index+2]
			}
		} else {
			break
		}
	}
	return 0
}
