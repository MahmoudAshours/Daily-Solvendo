package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function to handle the problem's logic
func main() {
	// Initialize a buffered scanner for fast input
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Example of how to read a single line and convert it to an integer
	n, _ := strconv.Atoi(readLine(reader))
	seats := readLine(reader)
	// Processing logic here
	result := solve(n, seats)

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
func solve(n int, seats string) int {
	k := 0
	for i, _ := range seats {
		if i < n-2 {
			if (seats[i] == '#' && seats[i] == seats[i+2]) && (seats[i+1] == '.') {
				k++
			}
		}
	}
	return k
}
