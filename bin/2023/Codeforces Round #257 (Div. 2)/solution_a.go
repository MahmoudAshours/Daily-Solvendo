// A. Jzzhu and Children
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
	var n, m int
	I("%d", &n)
	I("%d\n", &m)

	maxIndex := 1
	maxNum := math.MinInt32

	for i := 1; i <= n; i++ {
		var x int
		I("%d", &x)
		ceil := int(math.Ceil(float64(x) / float64(m)))
		if ceil >= maxNum {
			maxNum = ceil
			maxIndex = i
		}
	}

	O("%d", maxIndex)
	I("\n")

	return
}

func main() {
	solve()
	w.Flush()
}
