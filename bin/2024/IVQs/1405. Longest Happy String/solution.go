package main

import (
	"container/heap"
	"strings"
)

// CharFreq represents a character and its frequency
type CharFreq struct {
	freq int
	char byte
}

// PriorityQueue is a custom heap implementation for CharFreq
type PriorityQueue []*CharFreq

// The following methods implement heap.Interface
func (pq PriorityQueue) Len() int { return len(pq) }

// We want a max-heap, so we use > instead of <
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq > pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*CharFreq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	// Initialize the priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Push non-zero frequencies to the priority queue
	if a > 0 {
		heap.Push(&pq, &CharFreq{freq: a, char: 'a'})
	}
	if b > 0 {
		heap.Push(&pq, &CharFreq{freq: b, char: 'b'})
	}
	if c > 0 {
		heap.Push(&pq, &CharFreq{freq: c, char: 'c'})
	}

	// Initialize the result string builder
	var result strings.Builder

	// Main loop: continue while there are at least 2 characters left
	for pq.Len() > 1 {
		// Get the character with the highest frequency
		freq1 := heap.Pop(&pq).(*CharFreq)

		// If the result string has 0 or 1 character, we can safely add the current char
		if result.Len() <= 1 {
			result.WriteByte(freq1.char)
			freq1.freq--
			if freq1.freq > 0 {
				heap.Push(&pq, freq1)
			}
		} else {
			// Check if adding the current char would create three consecutive same characters
			if result.String()[result.Len()-1] == freq1.char && result.String()[result.Len()-2] == freq1.char {
				// If so, use the second most frequent character instead
				freq2 := heap.Pop(&pq).(*CharFreq)

				result.WriteByte(freq2.char)
				freq2.freq--

				// Push back the updated frequencies
				if freq2.freq > 0 {
					heap.Push(&pq, freq2)
				}
				if freq1.freq > 0 {
					heap.Push(&pq, freq1)
				}
			} else {
				// If no three consecutive same characters, add the current char
				result.WriteByte(freq1.char)
				freq1.freq--
				if freq1.freq > 0 {
					heap.Push(&pq, freq1)
				}
			}
		}
	}

	// Handle the case when only one character type is left
	if pq.Len() == 1 {
		freq1 := heap.Pop(&pq).(*CharFreq)

		// Add at most two of the remaining character
		if freq1.freq <= 1 {
			result.WriteByte(freq1.char)
		} else {
			result.WriteByte(freq1.char)
			result.WriteByte(freq1.char)
		}
	}

	// Return the final diverse string
	return result.String()
}

