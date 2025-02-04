package main

func main() {
	nums := []int{-1, 2, 9}
	print(subarraysDivByK(nums, 2))
}

func subarraysDivByK(nums []int, k int) int {
	subArraysCount := 0
	prefSum := 0
	freq := make(map[int]int)
	freq[0] = 1
	for _, v := range nums {
		prefSum += v
		remainder := prefSum % k
		if remainder < 0 {
			remainder += k
		}
		subArraysCount += freq[remainder]
		freq[remainder] += 1
	}
	return subArraysCount
}
