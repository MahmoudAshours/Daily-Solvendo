package main

import "fmt"

func main() {
	fmt.Println(stringSequence("he"))
}

func stringSequence(target string) []string {
	s := make([]string, 0)
	ss := "a"
	s = append(s, "a")
	i := 0

	for i != len(target) {
		if target[i] == ss[i] {
			i++
			s = append(s, s[len(s)-1]+"a")
			ss += "a"
		} else {
			added := string(rune((ss[len(ss)-1]-'a'+1)%26 + 'a'))
			ss = ss[:len(ss)-1] + added
			s = append(s, ss)
		}
	}
	return s[:len(s)-1]
}
