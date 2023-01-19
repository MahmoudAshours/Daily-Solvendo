// A. Jzzhu and Children
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		I("%d", &arr[i])
	}
	I("\n")
	sort.Ints(arr)
	arr[0], arr[n-1] = arr[n-1], arr[0]
	for _, v := range arr {
		O("%d ", v)
	}

	return
}

func main() {
	solve()
	w.Flush()
}
