package main

func main() {
	print(tribonacci(25))
}

func tribonacci(n int) int {
	trib := []int{0, 1, 1}
    if n == 0 {
        return 0
    }
	for i := 3; i < n+1; i++ {
		trib[0] , trib[1] , trib[2] = trib[1],trib[2],sum(trib)
	}
	return trib[2]
}
func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
