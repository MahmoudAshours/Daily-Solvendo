package main
func main(){
	print(findKthPositive([]int{2,3,4,7,11},5))
}


func findKthPositive(arr []int, k int) int {
    if k < arr[0] {  // handle edge case of k being outside of the array on the left
        return k
    }
    if k > arr[len(arr)-1] {  // outside on the right
        return len(arr) + k
    }
    // k is inside of the array
    l := 0
    r := len(arr)-1
    for l < r {
        mid := l + (r-l + 1)/2
        if arr[mid] - 1 - mid < k {
            l = mid
        } else {
            r = mid-1
        }
    }
    return l + 1 + k
}