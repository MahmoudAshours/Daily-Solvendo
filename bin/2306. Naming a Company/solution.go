package main

func main(){
	print(distinctNames([]string{"coffee","donuts","time","toffee"}))
}

func distinctNames(ideas []string) int64 {
    // Group tails by head
    headTails := make([]map[string]bool, 26)
    for i := range headTails {
        // Initialize all of our maps so we don't get NPE below
        headTails[i] = map[string]bool{}
    }
    for _, i := range ideas {
        // Offset to indexes 0-25
        headTails[i[0] - 'a'][i[1:]] = true
    }

    var count int64
    for i := range headTails {
        // Skip heads with no tail
        if len(headTails[i]) == 0 {
            continue
        }
        for j := i + 1; j < 26; j++ {
            // Again, skip heads with no tail
            if len(headTails[j]) == 0 {
                continue
            }

            // Count duplicate tails
            var dupes int
            for t := range headTails[i] {
                if headTails[j][t] {
                    dupes++
                }
            }

            // Our formula as described in the approach section
            count = count + int64(2 * (len(headTails[i]) - dupes) * (len(headTails[j]) - dupes))
        }
    }
    return count
}