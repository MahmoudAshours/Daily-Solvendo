package main

func main(){
findAnagrams("cbaebabacd","abc")
}
func findAnagrams(s string, p string) []int {
	var res []int

	if len(s) < len(p) {
		return res
	}

	m := make(map[byte]int)
	sMap := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}

	for i := 0; i < len(s); i++ {
		if i-len(p) >= 0 {
			sMap[s[i-len(p)]]--

			if sMap[s[i-len(p)]] == 0 {
				delete(sMap, s[i-len(p)])
			}
		}

		sMap[s[i]]++

			if len(m) == len(sMap) {
				isTheSame := true

				for key, val := range m {
					if v, ok := sMap[key]; !ok {
						isTheSame = false
						break
					} else {
						if v != val {
							isTheSame = false
							break
						}
					}
				}

				if isTheSame {
					res = append(res, i+1-len(p))
				}
			}
	}

	return res
}
