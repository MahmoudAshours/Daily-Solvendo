package main
import "fmt"
func main(){
	for _, v := range findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}) {
		print(v)
	}
}
func findAllConcatenatedWordsInADict(words []string) []string {
    dict := make(map[string]bool)

    for _, word := range words {
        dict[word] = true
    }

    res := []string{}
    for _, word := range words {
        wordLen := len(word)
        dp := make([]bool, wordLen+1)
        dp[0] = true

        for end := 1; end <= wordLen; end++ {
            start := 0
            if end == wordLen {
                start = 1 
            }
 
            for ; start < end && !dp[end]; start++ {
                if dp[start] && dict[word[start:end]] {
                    dp[end] = true
                }
            }

            if dp[wordLen] {
                res = append(res, word)
            }
        }
    }
    return res
}