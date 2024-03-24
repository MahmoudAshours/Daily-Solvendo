func bestTeamScore(scores []int, ages []int) int {
    players := [][]int{}
    
    for i := 0; i < len(scores); i++ {
        players = append(players, []int{ages[i], scores[i]})
    }
    
    sort.Slice(players, func(i, j int) bool {
		if players[i][0] > players[j][0] {
			return true
		} else if players[i][0] < players[j][0] {
			return false
		}
		return players[i][1] > players[j][1]
	})
    
    dp := make([]int, len(scores))
    for i := 0; i < len(scores); i++ {
        dp[i] = 0
    }
    ans := 0
    
    for i := 0; i < len(scores); i++ {
        dp[i] = players[i][1]
        score := dp[i]
        for j := 0; j < i; j++ {
            if players[i][0] == players[j][0] || players[i][1] <= 
players[j][1] {
                dp[i] = max(dp[i], dp[j] + players[i][1])
            }
            score = max(score, dp[i])
        }
        ans = max(ans, score)
    }
    return ans
}
func max(x, y int) int{
    if x > y {
        return x
    }
    return y
}
