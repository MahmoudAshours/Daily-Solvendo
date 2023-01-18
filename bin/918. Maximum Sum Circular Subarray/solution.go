package main

func main() {
	print(maxSubarraySumCircular([]int{1, -2, 3, -2}))
}
func maxSubarraySumCircular(nums []int) int {
	currMax := nums[0]
	maxSum := nums[0]

	currMin := nums[0]
	minSum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		currMax = max(currMax+nums[i], nums[i])
		maxSum = max(maxSum, currMax)

		currMin = min(currMin+nums[i], nums[i])
		minSum = min(minSum, currMin)

		sum += nums[i]
	}

	if sum == minSum { // when all number is negative
		return maxSum
	}
	return max(maxSum, sum-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
