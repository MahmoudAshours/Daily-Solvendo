package main

func main() {
}
func rotateString(s string, goal string) bool {
	goalLength := len(goal)
	i := 0
	for i < goalLength {
		if s == goal {
			return true
		}
		s = s[1:] + string(s[0])
		i++
	}
	return false
}

// abcde -> bcdea -> cdeab
