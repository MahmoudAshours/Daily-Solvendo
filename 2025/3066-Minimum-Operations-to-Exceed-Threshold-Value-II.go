import (
	"container/heap"
)

func minOperations(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	// Push all elements into the min heap
	for _, num := range nums {
		heap.Push(h, num)
	}

	count := 0
	for h.Len() > 1 && (*h)[0] < k {
		a := heap.Pop(h).(int)  // Smallest element
		b := heap.Pop(h).(int)  // Second smallest element

		newVal := 2*a + b
		heap.Push(h, newVal)  // Push the new value back into heap

		count++
	}

	return count
}

// MinHeap implementation
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
