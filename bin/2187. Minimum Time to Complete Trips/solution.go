package main
func main(){
	print(minimumTime([]int{1,2,3},5))
}

func minimumTime(time []int, totalTrips int) int64 {
	left, right := 0, math.MaxInt
	for left < right {
		mid := (left + right) >> 1
		if check(time, totalTrips, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func check(time []int, totalTrips, mid int) bool {
	cnt := 0
	for _, t := range time {
		cnt += mid / t
		if cnt >= totalTrips {
			return true
		}
	}
	return false
}