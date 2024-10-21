package main

import "fmt"

// import (
// "fmt"
// "sort"
// )
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	max := 0
	for _, n := range nums {
		m[n]++
		if m[n] > max {
			max = m[n]
		}
	}
	res := make([][]int, max+1)
	for k, v := range m {
		res[v] = append(res[v], k)
	}
	ans := make([]int, 0, k)
	ct := 0
	for i := max; i > 0; i-- {
		for j := 0; j < len(res[i]); j++ {
			ans = append(ans, res[i][j])
			ct++
			if ct == k {
				return ans
			}
		}
	}
	return ans
}
