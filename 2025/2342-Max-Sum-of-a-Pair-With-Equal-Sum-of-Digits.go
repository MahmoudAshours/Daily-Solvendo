func maximumSum(nums []int) int {
    res, mp := -1, [82]int{}
    for _, v := range nums {
        val, sum := v, 0
        for val > 0 {
            sum += val % 10 
            val /= 10
        }
        if mp[sum] > 0 {
            res = max(res, v + mp[sum])
        } 
        mp[sum] = max(mp[sum], v)
    } 
    return res
}