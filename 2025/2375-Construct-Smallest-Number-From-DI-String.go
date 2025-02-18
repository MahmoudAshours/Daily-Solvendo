
import (
\t\fmt\
)

// ListNode represents a node in the queue
type Node struct {
\trow, col, cost int
\tnext           *Node
}

// Queue represents a linked-list-based queue
type Queue struct {
\tfront, rear *Node
\tsize        int
}

// CreateNode creates a new ListNode
func CreateNode(row, col, cost int) *Node {
\treturn &Node{row: row, col: col, cost: cost, next: nil}
}

// CreateQueue creates a new Queue
func CreateQueue() *Queue {
\treturn &Queue{front: nil, rear: nil, size: 0}
}

// Enqueue adds a node to the rear of the queue
func (q *Queue) Enqueue(row, col, cost int) {
\tnewNode := CreateNode(row, col, cost)
\tif q.rear == nil {
\t\tq.front = newNode
        q.rear = newNode
\t} else {
\t\tq.rear.next = newNode
\t\tq.rear = newNode
\t}
\tq.size++
}

// Unshift adds a node to the front of the queue
func (q *Queue) Unshift(row, col, cost int) {
\tnewNode := CreateNode(row, col, cost)
\tif q.front == nil {
\t\tq.front = newNode
        q.rear = newNode
\t} else {
\t\tnewNode.next = q.front
\t\tq.front = newNode
\t}
\tq.size++
}

// Dequeue removes and returns a node from the front of the queue
func (q *Queue) Dequeue() *Node {
\tif q.front == nil {
\t\treturn nil
\t}
\tnode := q.front
\tq.front = q.front.next
\tif q.front == nil {
\t\tq.rear = nil
\t}
\tq.size--
\treturn node
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
\treturn q.size == 0
}

// IsInBounds checks if the position is within the grid bounds
func IsInBounds(r, c, m, n int) bool {
\treturn r >= 0 && r < m && c >= 0 && c < n
}

// MinimumObstacles finds the minimum obstacles encountered in a grid
func minimumObstacles(grid [][]int) int {
\tm, n := len(grid), len(grid[0])

\t// Directions: right, left, down, up
\tdirections := [4][2]int{
\t\t{0, 1}, // right
\t\t{0, -1}, // left
\t\t{1, 0},  // down
\t\t{-1, 0}, // up
\t}

\t// Create a queue for BFS
\tq := CreateQueue()
\tq.Enqueue(0, 0, 0)  // Start at (0, 0) with cost 0
\tgrid[0][0] = -1      // Mark as visited

\t// BFS loop
\tfor !q.IsEmpty() {
\t\tnode := q.Dequeue()
\t\tr, c, cost := node.row, node.col, node.cost

\t\t// If we've reached the bottom-right corner
\t\tif r == m-1 && c == n-1 {
\t\t\treturn cost
\t\t}

\t\t// Explore all 4 possible directions
\t\tfor _, dir := range directions {
\t\t\tnr, nc := r+dir[0], c+dir[1]
\t\t\tif IsInBounds(nr, nc, m, n) && grid[nr][nc] != -1 {
\t\t\t\tif grid[nr][nc] == 1 {
\t\t\t\t\tq.Enqueue(nr, nc, cost+1) // Cell has obstacle, add to back
\t\t\t\t} else {
\t\t\t\t\tq.Unshift(nr, nc, cost) // Free cell, add to front
\t\t\t\t}
\t\t\t\tgrid[nr][nc] = -1 // Mark as visited
\t\t\t}
\t\t}
\t}

\treturn 0 // If no valid path found
}
