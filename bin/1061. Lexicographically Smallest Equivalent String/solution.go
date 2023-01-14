package main

func main() {
	print(smallestEquivalentString("parker", "morris", "parser"))
}
func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	par := make([]int, 26, 26)
	for i := range par {
		par[i] = i
	}

	var find func(int) int

	find = func(i int) int {
		if par[i] == i {
			return i
		}
		par[i] = find(par[i])
		return par[i]
	}

	union := func(i, j int) {
		pi := find(i)
		pj := find(j)
		if pi == pj {
			return
		}

		if pi < pj {
			par[pj] = pi
		} else {
			par[pi] = pj
		}
	}

	for i := range s1 {
		c1 := int(s1[i] - 'a')
		c2 := int(s2[i] - 'a')
		union(c1, c2)
	}

	ans := ""

	for _, v := range baseStr {
		ans += string(rune(find(int(v-'a')) + 'a'))
	}
	return ans
}
