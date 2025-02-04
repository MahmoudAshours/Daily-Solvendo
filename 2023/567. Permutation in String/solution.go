package main
func main(){
	print(checkInclusion("ab","eidbaooo"))
}
func checkInclusion(s1 string, s2 string) bool {
    chars1 := [26]byte{}
    for i := range s1 {
        chars1[s1[i]-'a']++
    }

    chars2 := [26]byte{}
    lo, hi, ls2 := 0, 0, len(s2)
    for hi < ls2 {
        ch := s2[hi]-'a'
        chars2[ch]++
        hi++
        if chars2[ch] > chars1[ch] {
            for lo < hi && chars2[ch] > chars1[ch] {
                chars2[s2[lo]-'a']--
                lo++
            }
        }
        if isEqual(chars1, chars2) {
            return true
        }
    }

    return false
}

func isEqual(chars1, chars2 [26]byte) bool {
    for i := range chars1 {
        if chars1[i] != chars2[i] {
            return false
        }
    }
    return true
}