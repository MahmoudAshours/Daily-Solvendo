package main
func main(){
	print(minDistance("horse","ros"))
}

type State struct {
	i int16
	j int16
}

func minDistance(word1 string, word2 string) int {
	n := int16(len(word1))
	m := int16(len(word2))
	visited := make([][]bool, n+1)
	for i := range visited {
		visited[i] = make([]bool, m+1)
	}
	cur := []State{{}}
	var next []State
	addNext := func(i, j int16) {
		if !visited[i][j] {
			visited[i][j] = true
			next = append(next, State{i, j})
		}
	}
	for result := 0; ; cur, next, result = next, cur[:0], result+1 {
		for _, state := range cur {
			for state.i != n && state.j != m && word1[state.i] == word2[state.j] {
				state.i++
				state.j++
			}
			if state.i == n && state.j == m {
				return result
			}
			if state.i == n {
				addNext(n, state.j+1)
			} else if state.j == m {
				addNext(state.i+1, m)
			} else {
				addNext(state.i, state.j+1)
				addNext(state.i+1, state.j)
				addNext(state.i+1, state.j+1)
			}
		}
	}
}