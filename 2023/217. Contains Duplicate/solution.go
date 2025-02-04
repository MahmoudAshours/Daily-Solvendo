package main

func main() {
	print(containsDuplicate([]int{1, 2, 3, 2}))
}

type void struct{}

func containsDuplicate(nums []int) bool {
	set := map[int]void{}

	for _, v := range nums {
		set[v] = void{}
	} 
	if len(nums) == len(set) {
		return false
	}
	return true
}
