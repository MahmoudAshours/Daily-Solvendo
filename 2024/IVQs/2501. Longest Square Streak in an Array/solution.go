package main

import (
	"math"
	"sort"
)

func main() {
}
func longestSquareStreak(nums []int) int {
	mp := make(map[int]int)
	sort.Ints(nums)
	res := -1

	for _, num := range nums {
		sqrt := int(math.Sqrt(float64(num)))

		if sqrt*sqrt == num {
			if val, exists := mp[sqrt]; exists {
				mp[num] = val + 1
				if mp[num] > res {
					res = mp[num]
				}
			} else {
				mp[num] = 1
			}
		} else {
			mp[num] = 1
		}
	}

	return res
}
