package main

import (
	"fmt"
)

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	result := 0

	for start := 0; start < n; start++ {
		charCount := make([]int, 26)

		for end := start; end < n; end++ {
			charCount[s[end]-'a']++

			valid := false
			for _, count := range charCount {
				if count >= k {
					valid = true
					break
				}
			}

			if valid {
				result++
			}
		}
	}

	return result
}

func main() {
	// Example usage
	fmt.Println(numberOfSubstrings("abacb", 2)) // Output: 4
	fmt.Println(numberOfSubstrings("abcde", 1)) // Output: 15
}
