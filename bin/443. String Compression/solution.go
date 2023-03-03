package main
func main(){
	print(compress("a","a","b","b","c","c","c"))
}

func compress(chars []byte) int {
	// is - index in source slice, ir - index in resulting slice
	is, ir := 0, 0
	n := len(chars)

	for is < n {
		// moving through all the repeating characters
		i := is + 1
		for i < n && chars[i] == chars[is] {
			i++
		}

		// updating array
        // setting the next letter in resulting array
		chars[ir] = chars[is]
		ir++
		if cnt := i - is; cnt > 1 {
			digits := []byte(strconv.FormatInt(int64(cnt), 10))
			for j := 0; j < len(digits); j++ {
				chars[ir+j] = digits[j]
			}
			ir += len(digits)
		}
        // moving source index to the next unique letter
		is = i
	}
    
	chars = chars[:ir]
	return ir
}