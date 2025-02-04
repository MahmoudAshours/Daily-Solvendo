package main

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {0, 3}}
	labels := "bbbb"
	countSubTrees(3, edges[:], labels)
}
func dfs(curr int, adjList [][]int, count []int, labels string, ans []int) {
	if ans[curr] == 0 {
		ans[curr] = 1
		for _, next := range adjList[curr] {
			temp := make([]int, 26)
			dfs(next, adjList, temp, labels, ans)
			for i := 0; i < 26; i++ {
				count[i] += temp[i]
			}
		}
		count[labels[curr]-'a']++
		ans[curr] = count[labels[curr]-'a']
	}
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	adjList := make([][]int, n)
	count := make([]int, 26)
	ans := make([]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	dfs(0, adjList, count, labels, ans)
	return ans
}
