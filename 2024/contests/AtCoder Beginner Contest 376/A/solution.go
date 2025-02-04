package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := strings.Split(readLine(reader), " ")

	x, _ := strconv.Atoi(n[0])

	xx, _ := strconv.Atoi(n[1])
	arr := make([]int, 0)
	for x > 0 {
		reader, _ := reader.ReadString(' ')
		fmt.Println(arr)
		s, _ := strconv.Atoi(reader)
		arr = append(arr, s)
		x--
	}

	result := solve(xx, arr)

	fmt.Fprintln(writer, result)
}

// Reads a single line from the input
func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// Reads a line of integers from the input and returns them as a slice
func readIntArray(reader *bufio.Reader) []int {
	line := readLine(reader)
	parts := strings.Fields(line)
	arr := make([]int, len(parts))
	for i, v := range parts {
		arr[i], _ = strconv.Atoi(v)
	}
	return arr
}

// Problem-specific logic goes in the solve function
func solve(k int, arr []int) int {
	ss := arr[0]
	count := 0
	for _, v := range arr {
		if v-ss >= k {
			ss = v
			count++
		}
	}
	return count
}
