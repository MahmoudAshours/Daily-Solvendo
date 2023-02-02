package main
func main(){
	print(isAlienSorted([]string{"hello","leetcode"},""hlabcdefgijkmnopqrstuvwxyz""))
}

func isAlienSorted(words []string, order string) bool {
	m := make(map[string]int)
	for i := 0; i < len(order); i++ {
		m[string(order[i])] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if !isLexicographicallyOrder(words[i], words[i+1], m) {
			return false
		}
	}
	return true
}

func isLexicographicallyOrder(s1, s2 string, m map[string]int) bool {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return m[string(s1[i])] < m[string(s2[i])]
		}
	}
	return len(s1) <= len(s2)
}