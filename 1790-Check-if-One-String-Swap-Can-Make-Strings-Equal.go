func areAlmostEqual(s1 string, s2 string) bool {
	ind, count := -1, 0
	for i := 0; i < len(s2) && count <= 2; i++ {
		if s1[i] == s2[i] {
			continue
		} else if ind >= 0 && !(s1[ind] == s2[i] && s1[i] == s2[ind]) {
			return false
		} else {
			ind = i
		}
		count++
	}
	return count == 2 || count == 0
}
