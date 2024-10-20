package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n, _ := strconv.Atoi(readLine(reader))

	arr1 := readIntArray(reader)

	arr2 := readIntArray(reader)

	result := solve(n, arr1, arr2)

	// Output the result
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
func solve(n int, arr []int, arr2 []int) int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] > arr2[j]
	})

	N := len(arr)
	i := 0

	for i < N-1 && arr2[i] >= arr[i] {
		i++
	}

	for j := i; j < N-1; j++ {
		if arr2[j] < arr[j+1] {
			return -1
		}
	}
	return arr[i]
}
