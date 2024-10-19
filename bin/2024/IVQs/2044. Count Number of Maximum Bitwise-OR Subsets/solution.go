package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func countMaxOrSubsets(nums []int) int {
	mp := make(map[int]int, 0)
	res := 0
	backtrack(nums, 0, 0, &res, &mp)
	return mp[res]
}

func backtrack(nums []int, start int, curr int, res *int, visitedFreq *map[int]int) {

	for i := start; i < len(nums); i++ {

		tmp := curr | nums[i]
		*res = int(math.Max(float64(tmp), float64(*res)))
		(*visitedFreq)[tmp]++
		backtrack(nums, i+1, tmp, res, visitedFreq)
	}

}
