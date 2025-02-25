const mod = 1e9 + 7

func numOfSubarrays(arr []int) int {
    even := 1 // There's one prefix sum (0) which is even
    odd := 0
    prefixSum := 0
    result := 0

    for _, num := range arr {
        prefixSum += num
        if prefixSum%2 == 0 {
            // If the current prefix sum is even, the number of subarrays with odd sum is the number of odd prefix sums so far
            result = (result + odd) % mod
            even++
        } else {
            // If the current prefix sum is odd, the number of subarrays with odd sum is the number of even prefix sums so far
            result = (result + even) % mod
            odd++
        }
    }

    return result

}