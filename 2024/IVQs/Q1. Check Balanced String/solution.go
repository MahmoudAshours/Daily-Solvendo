package main

func main() {
}
func isBalanced(num string) bool {
	evensum := 0
	oddsum := 0
	for i, v := range num {
		s, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			evensum += s
		} else {
			oddsum += s
		}
	}
	return evensum == oddsum
}
