func isVowel(c byte)bool{
    if strings.Contains("aeiou",string(c)){
        return true
    }
    return false
}

func vowelCountSatisfy(vowelCount map[byte]int)bool{
    for _,c:=range([]byte("aeiou")){
        if vowelCount[c]==0{
            return false
        }
    }
    return true
}


func countOfSubstrings(word string, k int) int64 {
    n:=len(word)
    nextConsonant:=make([]int, n)
    ans:=0

    // computing nextConsonant Index for every position
    computeNextConsonant:=func(){
        lastConsonantIdx:=n
        for i:=n-1;i>=0;i--{
            nextConsonant[i]=lastConsonantIdx
            if !isVowel(word[i]){
                lastConsonantIdx=i
            }
        }
    }

    computeNumSubstring:=func(){
        r:=0
        consonantCount:=0
        vowelCount:=make(map[byte]int)

        for l:=0;l<n;l++{
            for r<n && (consonantCount <k || vowelCountSatisfy(vowelCount)==false){
                // It means word[i:r] is not satisfying. So we include current(rth) character as well.
                vowelCount[word[r]]++
                if !isVowel(word[r]){
                    consonantCount++
                }
                r++
            }
            if r>=n && (consonantCount <k || vowelCountSatisfy(vowelCount)==false){
                // Not possible to achieve the requirements.
                break
            }

            // If reached here consonantCount >= k
            if consonantCount==k{
                ans += nextConsonant[r-1]-(r-1)
            }
            vowelCount[word[l]]--
            if !isVowel(word[l]){
                consonantCount--
            }
        }
    }

    computeNextConsonant()
    computeNumSubstring()
    return int64(ans)
}