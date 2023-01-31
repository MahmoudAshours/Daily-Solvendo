package main

import "sort"

func main(){
	print(majorityElement([]int{1}))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	maxCount:=0
	minCount:=0
	key:=nums[0]
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] {
			minCount++
			if minCount >= maxCount{
				maxCount = minCount
				key = nums[i]
			}
		} else { 
			minCount = 0
		}  	
	}    
	return key
}