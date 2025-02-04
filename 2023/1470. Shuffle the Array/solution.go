package main

func main(){
	for _, v := range shuffle([]int{2,5,1,3,4,7},3) {
	 	print(v)
	 } 
}

func shuffle(nums []int, n int) []int {
    shuffled := make([]int, 0)
    for i := 0; i < n; i++ {
    	shuffled = append(shuffled,nums[i],nums[i+n])
    }
    return shuffled
}