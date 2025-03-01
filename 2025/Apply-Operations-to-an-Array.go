func applyOperations(nums []int) []int {
    for i,v := range nums[:len(nums)-1] {
        if v == nums[i+1]{
            nums[i] = nums[i] * 2
            nums[i+1] = 0
        }
    }
 
    zeroPtr := 0  
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[zeroPtr] = nums[zeroPtr], nums[i]  
            zeroPtr++  
        }
}
    return nums
}