package main

import (
	"fmt"
)

func main() {
	intervals := [][]int{{1, 5}}
	newInterval := []int{2, 3}

	whole := insert(intervals, newInterval)
	for _, v := range whole {
		fmt.Printf("{%d,%d},", v[0], v[1])
	}
}
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	if len(intervals) == 0 {
		res = append(res, newInterval)
	}

	done := false

	for i := 0; i < len(intervals); i++ {
		if done || newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])

			if newInterval[0] > intervals[i][1] && i == len(intervals)-1 {
				res = append(res, newInterval)
			}

			continue
		}

		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			res = append(res, intervals[i])
			done = true
			continue
		}

		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])

		if i == len(intervals)-1 {
			res = append(res, newInterval)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
