func numberOfSubstrings(s string) int {
    freq := map[byte]int{'a': 0, 'b': 0, 'c': 0}
    left, ans := 0, 0

    for right := 0; right < len(s); right++ {
        freq[s[right]]++

        for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
            ans += len(s) - right
            freq[s[left]]--
            left++
        }
    }

    return ans
}