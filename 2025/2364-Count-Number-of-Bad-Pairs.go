func longestPath(parent []int, s string) int {
     \tif len(parent) == 1 {
\t\treturn 1
\t}
\ttravel := make([][]int, len(parent))
\tfor i := 1; i < len(parent); i++ {
\t\ttravel[parent[i]] = append(travel[parent[i]], i)
\t}
\tvar result = 0
\tvar dfs func(int) (int, byte)
\tdfs = func(index int) (int, byte) {
\t\tif len(travel[index]) == 0 {
\t\t\treturn 1, s[index]
\t\t}
\t\tnodeMax := [2]int{}
\t\tfor _, v := range travel[index] {
\t\t\tl, n := dfs(v)
\t\t\tif n != s[index]{
\t\t\t\tif l > nodeMax[0]{
\t\t\t\t\tnodeMax[0] = l
\t\t\t\t\tif nodeMax[0] > nodeMax[1]{
\t\t\t\t\t\tnodeMax[0], nodeMax[1] = nodeMax[1], nodeMax[0]
\t\t\t\t\t}
\t\t\t\t}
\t\t\t}
\t\t}
\t\tif nodeMax[0]+nodeMax[1]+1 > result{
\t\t\tresult = nodeMax[0]+nodeMax[1]+1
\t\t}

\t\treturn nodeMax[1] +1, s[index]
\t}
    dfs(0)
    return result
}