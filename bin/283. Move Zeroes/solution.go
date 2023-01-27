package main

func main(){
	 moveZeroes([]int{0,1,0,3,12})
}
func moveZeroes(nums []int)  {
	j := 0
    for i, v := range nums {
    	if v!=0{
    		nums[i] , nums[j] = nums[j] , nums[i]
    		j++
    	}
    }
    for _, v := range nums {
    	print(v)
    }
}
 