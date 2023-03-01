package main
func main(){
	sortArray([]int{5,3,2,4})
}


func sortArray(nums []int) []int {
    sort.Ints(nums)
    return nums
}