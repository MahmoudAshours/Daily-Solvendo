package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSteps("0111"))
}
func minimumSteps(s string) int {
	minSwaps := 0
	countBlack := 0

	for _, char := range s {
		if char == '1' {
			countBlack++
		} else {
			// When a white ball is encountered, add the number of black balls before it to the swap count
			minSwaps += countBlack
		}
	}

	return minSwaps
}
