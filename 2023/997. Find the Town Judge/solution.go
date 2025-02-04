package main 

import "math"

func main(){
	print(findJudge(3,[][]int{{1,3},{2,3}}))
}

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n ==2 {
		return -1
	}
	if len(trust) == 0 && n == 1 {
		return 1
	}
	freqMapTrusted := make(map[int]int, 0)
	freqMapTrusting := make(map[int]int, 0)
	maxVal := math.MinInt32
	key := 1
	for _, v := range trust {
		freqMapTrusted[v[1]]++
		freqMapTrusting[v[0]]++
	}
	
	for k, v := range freqMapTrusted {
		if v > maxVal {
			maxVal = v
			key = k
		}
	}

	if _,ok:=freqMapTrusting[key];ok || maxVal != n-1{
		return -1
	}
	return key
}


/**
 * Alternative  solution
 * func findJudge(n int, trust [][]int) int {
    trustSomeone := make([]bool, n)
    trustedBy := make([]int, n)

    for i := range trust {
        u := trust[i][0] - 1
        v := trust[i][1] - 1
    
        trustSomeone[u] = true
        trustedBy[v] += 1
    }

    for i := 0; i < n; i++ {
        if !trustSomeone[i] && trustedBy[i] == n - 1 {
            return i + 1
        }
    }
    return -1
}
 * */