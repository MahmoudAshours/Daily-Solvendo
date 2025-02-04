package main

func main() {
}
func canSortArray(nums []int) bool {
	prev := 0
	prevBig := 0
	currBig := 0
	for i := 0; i < len(nums); i++ {
		if bitsCount(nums[i]) != prev {
			prevBig = currBig
			currBig = nums[i]
			prev = bitsCount(nums[i])
		} else {
			if nums[i] > currBig {
				currBig = nums[i]
			}
		}
		if nums[i] < prevBig {
			return false
		}
	}
	return true
}
func bitsCount(n int) int {

	result := 0
	for n > 0 {
		result += n % 2
		n = n / 2
	}
	return result
}
