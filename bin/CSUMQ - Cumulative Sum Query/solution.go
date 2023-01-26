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
	var q int
	I("%d\n", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		I("%d", &arr[i]) 
	}
	I("\n")
	I("%d\n", &q)

	for k := 0; k < q; k++ {
		sum := 0
		var i int 
		var j int
		I("%d%d", &i,&j)
 		I("\n") 
 		var slice []int
 		if j == n{
 			slice = arr[i:j]
 		}else{
 			slice = arr[i:j+1]
 		}
		for _, v := range slice {
			sum+=v
		} 
		O("%d\n",sum)
	}

}

func main() {
	 
	solve()
	w.Flush()
}
