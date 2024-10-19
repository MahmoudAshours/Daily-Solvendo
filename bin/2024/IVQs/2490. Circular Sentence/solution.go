package main

import "strings"

func main() {
	print(isCircularSentence("leetcode exercises sound delightful"))
}

func isCircularSentence(sentence string) bool {
	words := strings.Fields(sentence)
	firstWord := words[0]
	lastWord := words[len(words)-1]
	if firstWord[0] != lastWord[len(lastWord)-1] {
		return false
	}
	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]
		if currentWord[len(currentWord)-1] != nextWord[0] {
      print(currentWord);
			return false
		}
	}
	return true
}
