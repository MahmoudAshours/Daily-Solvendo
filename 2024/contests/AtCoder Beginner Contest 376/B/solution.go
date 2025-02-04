package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := strings.Split(readLine(reader), " ")

	length, _ := strconv.Atoi(n[0])

	instr, _ := strconv.Atoi(n[1])

	arr := makeRange(1, length)

	lPost := 1
	RPos := 2
	res := 0
	for instr > 0 {
		s := strings.Split(readLine(reader), " ")
		noTimes, _ := strconv.Atoi(s[1])
		solve(rune(s[0][0]), noTimes, arr, &res, &lPost, &RPos)
		instr--
	}

	fmt.Fprintln(writer, res)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
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
func solve(char rune, destination int, arr []int, res *int, lPos *int, rPos *int) int {
	if char == 'R' {
		if destination != *rPos {
			*res += int(math.Abs(float64(*rPos - destination)))
			*rPos = destination
		}
	} else {
		if destination != *lPos {
			*res += len(arr) - int(math.Abs(float64(*lPos-destination)))
			*lPos = destination
		}

	}
	return *res
}
