package main

func main() {
	parent := []int{-1, 0, 0, 1, 1, 2}
	print(longestPath(parent, "abacbe"))
}
func longestPath(parent []int, s string) int {
	if len(parent) == 1 {
		return 1
	}
	travel := make([][]int, len(parent))
	for i := 1; i < len(parent); i++ {
		travel[parent[i]] = append(travel[parent[i]], i)
	}
	var result = 0
	var dfs func(int) (int, byte)
	dfs = func(index int) (int, byte) {
		if len(travel[index]) == 0 {
			return 1, s[index]
		}
		nodeMax := [2]int{}
		for _, v := range travel[index] {
			l, n := dfs(v)
			if n != s[index] {
				if l > nodeMax[0] {
					nodeMax[0] = l
					if nodeMax[0] > nodeMax[1] {
						nodeMax[0], nodeMax[1] = nodeMax[1], nodeMax[0]
					}
				}
			}
		}
		if nodeMax[0]+nodeMax[1]+1 > result {
			result = nodeMax[0] + nodeMax[1] + 1
		}

		return nodeMax[1] + 1, s[index]
	}
	dfs(0)
	return result
}
