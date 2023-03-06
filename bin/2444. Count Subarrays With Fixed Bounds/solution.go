package main 


func main(){
	print(countSubarrays([]int{1,3,5,2,7,5},1,5))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var ans int64
	lastMin, lastMax := -1, -1
	leftmost := 0
	for i, d := range nums {
		if d < minK || d > maxK {
			leftmost = i + 1
			lastMin = -1
			lastMax = -1
			continue
		}
		if d == minK {
			lastMin = i
		}
		if d == maxK {
			lastMax = i
		}
        if lastMin != -1 && lastMax != -1 {
            ans += int64(min(lastMin, lastMax)-leftmost+1)
        }
	}
	
	return ans
}
