package main

func main(){
print(convert("PAYPALISHIRING",3))
}
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result []uint8
	var ri int
	for r := 0; r < numRows; r++ {
		for i := r; i < len(s); {
			result = append(result, s[i])
			if i != ri+i && ri+i < len(s) {
				result = append(result, s[ri+i])
			}
			i += 2 * (numRows - 1)
		}
		ri = 2 * (numRows - 2 - r)
	}
	return string(result)
}
