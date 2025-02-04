package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeFancyString("aaabaaaa"))
}
func makeFancyString(s string) string {
	arr := []byte{s[0]}
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] != s[i] || s[i] != s[i+1] {
			arr = append(arr, s[i])
		}
	}
	if len(s) >= 2 {
		arr = append(arr, s[len(s)-1])
	}
	return string(arr)
}
