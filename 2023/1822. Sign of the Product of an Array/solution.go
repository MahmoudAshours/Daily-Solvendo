package main

func main() {
	nums := []int{7, 3, 3, 6, 6, 6, 10, 5, 9, 2}
	print(arraySign(nums[:]))
}
func arraySign(nums []int) int {
	numsBelowZero := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			numsBelowZero++
		}
	}

	if numsBelowZero%2 == 1 {
		return -1
	}

	return 1
}
