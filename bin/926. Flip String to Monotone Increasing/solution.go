package main

import "math"

func main() {
	print(minFlipsMonoIncr("00011000"))
}

func minFlipsMonoIncr(s string) int {
	zerosCount := 0
	for _, v := range s {
		if v == rune('0') {
			zerosCount++
		}
	}

	ans := zerosCount

	for _, v := range s {
		if v == rune('0') {
			zerosCount--
			ans = int(math.Min(float64(ans), float64(zerosCount)))
		} else {
			zerosCount++
		}
	}

	return ans
}
