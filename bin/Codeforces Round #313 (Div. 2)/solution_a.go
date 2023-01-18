// B. Card Deck
package main

import (
	"bufio"
	"fmt"
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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		I("%d", &arr[i])
	}

	I("\n")

	for _, v := range arr {
		if v == 1 {
			O("-1\n")
			return
		}
	}
	O("1")
	return
}

func main() {
	solve()
	w.Flush()
}
