package main

import "sort"

func main() {
	nums := []int{7, 3, 3, 6, 6, 6, 10, 5, 9, 2}
	print(maxIceCream(nums[:], 56))
}
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	count := 0
	if coins < costs[0] {
		return 0
	}
	for i := 0; coins > 0 && i < len(costs); i++ {
		if coins-costs[i] >= 0 {
			coins -= costs[i]
			count += 1
		}
	}
	return count
}
