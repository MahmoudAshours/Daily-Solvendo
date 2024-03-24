package main
func main(){
	print(minEatingSpeed([]int{3,6,7,11},8))
}

func minEatingSpeed(piles []int, h int) int {
    left, right := 1, max(piles)
    
    for left < right {
        mid := (left + right) / 2
        hours := 0
        
        for _, pile := range(piles) {
            hours += int(math.Ceil(float64(pile) / float64(mid)))
        }
        
        if hours <= h {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return right
}

func max(arr []int) int {
    max := arr[0]
    
    for _, val := range(arr) {
        if max < val {
            max = val
        }
    }
    return max
}