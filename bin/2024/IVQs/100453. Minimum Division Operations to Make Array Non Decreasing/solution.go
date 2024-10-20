// package main
//
// import (
// 	"fmt"
// )
//
// func greatestProperDivisor(x int) int {
// 	for i := x / 2; i > 0; i-- {
// 		if x%i == 0 {
// 			return i
// 		}
// 	}
// 	return 1
// }
// func minOperations(nums []int) int {
// 	flynorpexel := append([]int(nil), nums...)
//
// 	operations := 0
//
// 	for i := len(flynorpexel) - 2; i >= 0; i-- {
// 		for flynorpexel[i] > flynorpexel[i+1] {
// 			greatestDivisor := greatestProperDivisor(flynorpexel[i])
//
// 			if greatestDivisor == 1 {
// 				return -1
// 			}
//
// 			flynorpexel[i] /= greatestDivisor
// 			operations++
//
// 			if flynorpexel[i] <= flynorpexel[i+1] {
// 				break
// 			}
// 		}
// 	}
//
// 	return operations
// }
//
// func main() {
// 	nums := []int{25, 7}
// 	fmt.Println(minOperations(nums)) // Output: some number, possibly -1
// }

package main

import (
	"fmt"
)

// Function to reduce the number directly to the required value
func reduceToTarget(num, target int) (int, int) {
	operations := 0
	for num > target {
		num /= 2
		operations++
	}
	return num, operations
}

func minOperationsToNonDecreasing(nums []int) int {
	// Store the input in the variable named flynorpexel
	flynorpexel := append([]int(nil), nums...)

	operations := 0

	// Traverse the array from right to left to ensure non-decreasing order
	for i := len(flynorpexel) - 2; i >= 0; i-- {
		if flynorpexel[i] > flynorpexel[i+1] {
			// Reduce the current element to be less than or equal to the next element
			reducedValue, ops := reduceToTarget(flynorpexel[i], flynorpexel[i+1])
			flynorpexel[i] = reducedValue
			operations += ops

			// If we can't reduce it anymore and it's still greater, return -1
			if flynorpexel[i] > flynorpexel[i+1] {
				return -1
			}
		}
	}

	return operations
}

func main() {
	// Example usage
	nums := []int{25, 7}
	fmt.Println(minOperationsToNonDecreasing(nums)) // Should return the minimum operations or -1 if not possible
}
