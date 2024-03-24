// B. Card Deck
package main

import (
	"bufio"
	"fmt"
	"math"

	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

func solve() {
	var n int
	I("%d\n", &n)
	res := make([]int, n)
	deck := make([]int, n)
	out := make([]int, 0)
	for i := 0; i < n; i++ {
		I("%d", &deck[i])
	}

	I("\n")
	res = recursiveGreedy(deck, out)
	for _, v := range res {
		O("%d ", v)
	}
	O("\n")
}

func main() {
	var t int
	I("%d\n", &t)
	for t > 0 {
		solve()
		t--
	}
	w.Flush()
}

func getMax(arr []int) int {
	maxValue := math.MinInt
	maxIndex := 0
	for i, v := range arr {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return maxIndex
}

func recursiveGreedy(arr []int, out []int) []int {
	if len(arr) == 0 {
		return out
	}
	max := getMax(arr)
	out = append(out, arr[max:]...)
	arr = append([]int{}, arr[:max]...)
	return recursiveGreedy(arr, out)
}
