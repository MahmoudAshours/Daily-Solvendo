package main

import (
	"strconv"
	"math"
)

func main(){
	print(addBinary("1010","1011"))
}
func addBinary(a string, b string) string {
	if len(a) < len(b) {a, b = b, a}
	lna, lnb := len(a), len(b)
	ans := make([]byte, lna+1)
	var A, B, C byte
	for i := 0; i <= lna; i++ {
		ia, ib, ians := lna-1-i, lnb-1-i, lna-i
		if ia >= 0 {A = a[ia] - '0'} else {A = 0}
		if ib >= 0 {B = b[ib] - '0'} else {B = 0}
		ans[ians], C = (A ^ B ^ C) + '0', (A & B) + (C & (A ^ B))
	}
	if ans[0] == '0' {return string(ans[1:])}
	return string(ans)
}