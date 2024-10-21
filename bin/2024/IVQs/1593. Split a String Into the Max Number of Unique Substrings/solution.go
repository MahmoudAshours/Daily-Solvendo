package main

import "fmt"

func main() {
	fmt.Println(maxUniqueSplit("wwwzfvedwfvhsww"))
	// fmt.Println(maxUniqueSplit("aa"))
}

func maxUniqueSplit(s string) int {
	return backtrack(s, make(map[string]bool))
}

func backtrack(s string, seen map[string]bool) int {
	if len(s) == 0 {
		return 0
	}

	maxCount := 0
	for i := 1; i <= len(s); i++ {
		substr := s[:i]

		if !seen[substr] {
			seen[substr] = true
			count := 1 + backtrack(s[i:], seen)
			maxCount = max(maxCount, count)
			delete(seen, substr)
		}
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
