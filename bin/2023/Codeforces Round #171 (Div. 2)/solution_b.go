// #B. Books

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
	var n,t int
	I("%d", &n)
	I("%d\n", &t)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		I("%d", &arr[i])
	}
	I("\n")
	// Solution

	prefixSum:=make([]int, n)
	prefixSum[0] = arr[0]
	
	for i := 0 ; i < n - 1 ; i++{
		prefixSum[i+1] = arr[i]+arr[i+1]
	}


	l := 0
	
	count:=prefixSum[n-1]

	for count > t {
		count = count - prefixSum[l]
		if count <= t{
			break
		}else{
			l++
		}
		
	}
	
	  
		O("%d",n-l)
	 
	O("\n")

}

func main() {
 	solve() 
	w.Flush()
}