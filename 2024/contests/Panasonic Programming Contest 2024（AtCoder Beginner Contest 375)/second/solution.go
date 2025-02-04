package main

import (
	"bufio"
	"fmt"
	"math"
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
	arr := make([][]float64, 0)
	for n > 0 {
		coord := readLine(reader)
		splitted := strings.Split(coord, " ")
		first, _ := strconv.ParseFloat(splitted[0], 64)
		second, _ := strconv.ParseFloat(splitted[1], 64)
		arr = append(arr, []float64{first, second})
		n--
	}
	result := solve(x, arr)

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
func solve(n int, arr [][]float64) float64 {
	startPoint := []float64{0, 0}
	var res float64

	res += calcDistance(startPoint, arr[0])

	for i, _ := range arr {
		if i < n-1 {
			println(n)
			res += calcDistance(arr[i], arr[i+1])
		}
	}
	res += calcDistance(startPoint, arr[len(arr)-1])
	return res
}

func calcDistance(arr1 []float64, arr2 []float64) float64 {
	first := float64(arr1[0]-arr2[0]) * float64(arr1[0]-arr2[0])
	second := float64(arr1[1]-arr2[1]) * float64(arr1[1]-arr2[1])
	return math.Sqrt(first + second)
}
