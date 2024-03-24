package main

func main(){
singleNonDuplicate([]int{1,1,2,3,3,4,4,8,8})
}

func singleNonDuplicate(nums []int) int {
    for i := 0; i < len(nums)-2; i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
