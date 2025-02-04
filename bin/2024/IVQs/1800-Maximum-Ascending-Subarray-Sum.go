func maxAscendingSum(nums []int) int {
    maxSum := nums[0]
    localSum := nums[0]
    for i,v := range nums[:len(nums)-1]{
        if v < nums[i+1]{
            localSum += nums[i+1]
        }else{
            maxSum = max(maxSum,localSum)
            localSum = nums[i+1]
        }
    }
    maxSum = max(maxSum,localSum)
    return maxSum
}