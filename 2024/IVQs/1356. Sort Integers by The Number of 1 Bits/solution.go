package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if countOnes(arr[i]) == countOnes(arr[j]) {
			return arr[i] < arr[j] 
		}
		return countOnes(arr[i]) < countOnes(arr[j]) 
	})
	return arr
}

// 1 0 1 1
// 0 1 1 1

func countOnes(num int) int {
	count := 0
	for num > 0 {
		count += num & 1 // Increment count if the least significant bit is 1
		num >>= 1        // Right shift the number by 1 bit
	}
	return count
}
