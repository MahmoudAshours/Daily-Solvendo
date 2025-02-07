func queryResults(limit int, queries [][]int) []int {
	n := len(queries)
	result := make([]int, 0, n)
	colorMap := map[int]int{}
	ballMap := map[int]int{}

	for _, query := range queries {
		ball := query[0]
		color := query[1]

		if prevColor, ok := ballMap[ball]; ok {
			colorMap[prevColor]--
			if colorMap[prevColor] == 0 {
				delete(colorMap, prevColor)
			}
		}
        
		ballMap[ball] = color
		colorMap[color]++
		result = append(result, len(colorMap))
	}

	return result
}