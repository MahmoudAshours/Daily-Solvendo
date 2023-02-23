package main
func main(){
	print(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10},5))
}
func shipWithinDays(weights []int, days int) int {
    // Defining a function to check if it's possible to send
    // all the packages withing a given days with some capacity
	fitsCapacity := func(cap int) bool {
		s, d := 0, 0
		for i := 0; i < len(weights) && d < days; i++ {
			w := weights[i]
			if s+w <= cap {
				s += w
			} else {
				d++
				s = w
			}
		}
		return d < days
	}

    // Calculating possible values for the capacity
    // - For min capacity - it can't be less than the max weight, 
    // because it would be impossible to the max weight otherwise
    // - For max capacity - it's a sum of all the weights 
    // to be able to send everything in any number of days (e.g. in 1 day)
	i, j := 0, 0
	for _, w := range weights {
		if w > i {
			i = w
		}
		j += w
	}

    // Using binary search from the min cap to the max cap
    // to find the answer
	for i < j {
		h := int(uint(i+j) >> 1)
		if !fitsCapacity(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}