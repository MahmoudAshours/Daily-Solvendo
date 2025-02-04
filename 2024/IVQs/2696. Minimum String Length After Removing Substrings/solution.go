package main

import (
	"fmt"
	"strings"
)

func main() {
  fmt.Println(minLength("ACBBD"))
}

func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		if strings.Contains(s, "AB") {
			s = strings.Replace(s, "AB", "", 1)
		} else if strings.Contains(s, "CD") {
			s = strings.Replace(s, "CD", "", 1)
		}
	}
	return len(s)
}
