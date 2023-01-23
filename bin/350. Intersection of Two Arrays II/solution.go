package main

func main() {
	for _, v := range intersect([]int{1, 2, 2, 1}, []int{2, 2}) {
		print(v)
	}
}
func intersect(nums1 []int, nums2 []int) []int {
	common := make(map[int]int)
	res := make([]int, 0, 1)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for _, n := range nums1 {
		common[n]++
	}

	for _, n := range nums2 {
		if v, ok := common[n]; ok && v > 0 {
			res = append(res, n)
			common[n]--
		}
	}

	return res
}
