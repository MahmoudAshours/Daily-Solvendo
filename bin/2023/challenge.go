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
	arr1 := make([]int, n)

	for i := 0; i < n; i++ {
		I("%d", &arr[i])
	}

	I("\n")
	for i := 0; i < n; i++ {
		I("%d", &arr1[i])
	}
	sort.Ints(arr)
	sort.Ints(arr1)
	
	prod7oska:=arr[len(arr1)-1] * arr1[0]
	prodWael:=arr1[len(arr1)-1] * arr[0]
		
	
	if(prod7oska<prodWael){
	O("7oSkaa")
	}else if(prod7oska>prodWael) {
	O("Wael")
	}else{
	O("Tie")
	}
	O("\n")
}

func main() {
	solve()
	w.Flush()
}