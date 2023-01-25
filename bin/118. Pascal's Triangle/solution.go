package main

func main() {
	for _, v := range generate(5) {
		for _, d := range v {
			print(d)
		}
		print(" ")
	}
}

func generate(numRows int) [][]int {
	generated := make([][]int, 0)
	generated = append(generated, []int{1})
	counter := 0
	for k := 0; k < numRows-1; k++ {
		temp := append([]int{0}, generated[len(generated)-1]...)
		temp = append(temp, 0)
		row := make([]int, 0)
		for i := 0; i <= len(generated[len(generated)-1]); i++ {
			row = append(row, temp[i]+temp[i+1])
		}
		generated = append(generated, row)
		counter++
	}
	return generated
}
