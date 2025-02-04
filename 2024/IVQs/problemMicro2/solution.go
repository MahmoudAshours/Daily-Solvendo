package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	reverseKGroup(node, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	if k == 0 || head == nil {
		return head
	}

	x := head
	currentHead := head

	for x != nil{
		count++
		if count == k {
			recurs(currentHead, k)
			currentHead = x.Next
			count = 0
		}
		x = x.Next
	}
	return head
}

func recurs(head *ListNode, k int) {
	count := 0
	start := head
	tmp := head

	if k == 0 {
		return
	}

	for tmp != nil {
		count++
		if count == k {
			swap(&tmp.Val, &start.Val)
			break
		}
		tmp = tmp.Next
	}
	if head.Next != nil {
		recurs(head.Next, k-2)
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
