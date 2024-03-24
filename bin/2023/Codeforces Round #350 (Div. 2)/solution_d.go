// D1. Magic Powder - 1

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
	// Problem Initialization
	var n,mp int
	I("%d", &n)
	I("%d\n", &mp)

	gramsNeeded := make([]int, n)
	gramsPresent := make([]int, n)

	for i := 0; i < n; i++ {
		I("%d", &gramsNeeded[i])
	}
	I("\n")
	for i := 0; i < n; i++ {
		I("%d", &gramsPresent[i])
	}
	I("\n")
	// Solution

	gramsDivision:= make([]int, n)
	
	for i, v := range gramsPresent {
		gramsDivision = append(gramsDivision , int(v/gramsNeeded[i]))
	}
	
	sort.Ints(gramsDivision)

 	for _, v := range gramsDivision {
		O("%d",v)
	}

	O("\n")
}

func main() {  
	solve() 
	w.Flush()
}