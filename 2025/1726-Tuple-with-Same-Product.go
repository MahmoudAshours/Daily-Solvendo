func tupleSameProduct(nums []int) int {
	ans := 0
	mapChar := map[int]int{}
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			mapChar[nums[i] * nums[j]]++
		}
	}

	for _, Char := range mapChar {
		ans += (Char - 1) * Char * 4
	}

	return ans
}