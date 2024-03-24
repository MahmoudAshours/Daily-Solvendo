package main

func main(){
	print(isAnagram("anagram","nagaram"))
}

func isAnagram(s string, t string) bool {
	countChars := func(ss string) [26]int {
		var count [26]int
		for _, ch := range ss {
			count[ch-'a']++
		}
		return count
	}
	return countChars(s) == countChars(t)
}