package main
func main(){
	print(strStr("sadbutsad","sad"))
}

func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {return -1}
    
    if haystack == needle {return 0}
    
    if len(needle) == 0 {return 0}
    
    if len(haystack) == 0 {return -1}
    
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle[0] {
            j, begin := i, 0
            
            for {
                j++
                begin++
                
                if begin == len(needle) || j == len(haystack) || haystack[j] != needle[begin] {
                    break
                }
            }

            if begin == len(needle) {
                return j - len(needle)
            }
        }
    }
    
    return -1
}