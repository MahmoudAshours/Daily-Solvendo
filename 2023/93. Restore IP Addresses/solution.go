package main

import (
	"math"
	"strconv"
)

func main() {
	res := restoreIpAddresses("25525511135")
	for _, v := range res {
		print(v)
	}
}

func restoreIpAddresses(s string) []string {
	valid := []string{}

	if len(s) > 12 {
		return valid
	}

	var backtracking func(int, int, string)
	backtracking = func(starting int, numberOfDots int, curIP string) {
		if numberOfDots == 4 && starting == len(s) {
			valid = append(valid, curIP[:len(curIP)-1])
			return
		}
		if numberOfDots > 4 {
			return
		}
		for i := starting; i < int(math.Min(float64(starting+3), float64(len(s)))); i++ {
			val, _ := strconv.Atoi(s[starting : i+1])
			if val < 256 && (starting == i || s[starting] != '0') {
				backtracking(i+1, numberOfDots+1, curIP+s[starting:i+1]+".")
			}
		}
	}
	backtracking(0, 0, "")
	return valid
}
