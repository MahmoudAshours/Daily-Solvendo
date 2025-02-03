func longestMonotonicSubarray(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    inc := 1
    dec := 1 
    maxm := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            inc++
            dec = 1
        }
        if nums[i] < nums[i-1] {
            dec++
            inc = 1
        }
        if nums[i] == nums[i-1] {
            inc = 1
            dec = 1
        }
        maxm = max(inc, dec, maxm)
    }
    return maxm
}