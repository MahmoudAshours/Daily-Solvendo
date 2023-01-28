package main

func main() {
	print(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	freqMap := make(map[rune]int)
	for _, v := range s {
		freqMap[v]++
	}
	for i, v := range s {
		if freqMap[v] == 1 {
			return i
		}
	}
	return -1
}
