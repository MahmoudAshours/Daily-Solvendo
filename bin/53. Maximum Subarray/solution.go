package main 

import (
	"math"
	"sort"
)

func main(){
	print(maxSubArray([]int{5,4,-1,7,8}))
}


func maxSubArray(nums []int) int {
    prefixSum := make([]int, 0)
	sum := nums[0]
	prefixSum = append(prefixSum,sum)
	for i := 0; i < len(nums) - 1; i++ {
			sum = int(math.Max(float64(sum + nums[i+1]),float64(nums[i+1])))    	
			prefixSum = append(prefixSum,sum)
}    	
	sort.Ints(prefixSum)
    return prefixSum[len(nums) - 1]
}