package main

import "sort"

func main() {
	nums := []int{1, 2, 4}
	print(canMakeArithmeticProgression(nums[:]))
}
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[0] - arr[1]
	for i := 1; i < len(arr)-1; i++ {
		if diff != (arr[i] - arr[i+1]) {
			return false
		}
	}
	return true
}
