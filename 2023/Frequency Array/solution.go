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
	freq := map[int]int{} 

	I("%d\n", &n)
	for i := 0; i < n; i++ {
		var in int
		I("%d", &in) 
		freq[in]++
	}

	I("\n")

	if len(freq) != n {

		O("ne krasivo")
	} else {

		O("prekrasnyy")
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
