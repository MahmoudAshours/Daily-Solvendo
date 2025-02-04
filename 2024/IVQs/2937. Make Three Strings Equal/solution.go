package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMinimumOperations("abc", "abb", "ab"))
}

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	// var operations int
	minimumStringLenght := math.Min(float64(len(s1)), math.Min(float64(len(s2)), float64(len(s3))))
	return int(minimumStringLenght)
}
