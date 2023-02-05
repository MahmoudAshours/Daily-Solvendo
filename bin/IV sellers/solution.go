package main

import "math"

func main() {
	print(sellers([]int{5, 4, 6}, []int{3, 1, 5}, 1))
}

func sellers(before []int, after []int, k int) int {

	var minimumSum uint64
	lhashmap := make(map[int][]int, 0)

	for i, v := range before {

		_, present := lhashmap[v-after[i]]

		if present {
			lhashmap[v-after[i]] = append(lhashmap[v-after[i]], i)
		} else {
			lhashmap[v-after[i]] = make([]int, 0)
			lhashmap[v-after[i]] = append(lhashmap[v-after[i]], i)

		}
	}

	for _, v := range lhashmap {
		for _, arrinside := range v {
			if k != 0 {
				minimumSum += uint64(before[arrinside])
				before[arrinside] = -1
				k--
			} else {
				break
			}
		}
	}

	for i, v := range before {
		if v != -1 {
			minimumSum += uint64(math.Min(float64(before[i]), float64(after[i])))
		}
	}
	return int(minimumSum)
}
