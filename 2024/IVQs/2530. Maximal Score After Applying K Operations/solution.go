package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

// An IntHeap is a max-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Reversed for max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the maximum element from the heap.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := &IntHeap{}
	heap.Init(h)

	// Add all elements to the max-heap.
	for _, num := range nums {
		heap.Push(h, num)
	}

	var res int64
	for i := 0; i < k; i++ {
		// Extract the largest element.
		largest := heap.Pop(h).(int)
		res += int64(largest)

		// Replace the element with ceil(largest / 3) and push it back into the heap.
		heap.Push(h, int(math.Ceil(float64(largest)/3)))
	}

	return res
}
