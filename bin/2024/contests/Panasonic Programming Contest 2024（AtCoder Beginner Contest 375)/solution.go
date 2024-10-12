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

	n, _ := strconv.Atoi(readLine(reader))
	x := n
	arr := make([][]rune, 0)
	for n > 0 {
		line := readLine(reader)
		runearr := []rune(line)
		arr = append(arr, runearr)
		n--
	}
	result := solve(x, arr)
	printGrid(result)
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
func solve(n int, grid [][]rune) [][]rune {
	// Create a copy of the grid to store changes simultaneously
	newGrid := make([][]rune, n)
	for i := range newGrid {
		newGrid[i] = make([]rune, n)
		copy(newGrid[i], grid[i])
	}

	// Perform the operation for each i from 1 to N/2
	for i := 1; i <= n/2; i++ {
		for x := i - 1; x <= n-i; x++ {
			for y := i - 1; y <= n-i; y++ {
				newGrid[y][n-1-x] = grid[x][y]
			}
		}
	}

	return newGrid
}
func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
