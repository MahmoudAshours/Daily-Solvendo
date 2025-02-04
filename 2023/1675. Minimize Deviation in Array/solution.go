package main 

func main(){
   print(minimumDeviation([]int{1,2,3,4}))
}

type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumDeviation(nums []int) int {
    n, minNum, deviation := len(nums), math.MaxInt32, math.MaxInt32
    h := make(IntHeap, n)
    for i := range nums {
        if nums[i] % 2 == 0 {
            h[i] = nums[i]
        } else {
            h[i] = nums[i] * 2
        }
        minNum = min(minNum, h[i])
    }
	heap.Init(&h)
    for {
        maxNum := heap.Pop(&h).(int)
        deviation = min(deviation, maxNum - minNum)
        if maxNum % 2 == 1 {
            break
        }
        maxNum >>= 1
        heap.Push(&h, maxNum)
        minNum = min(minNum, maxNum)
    }
    return deviation
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}