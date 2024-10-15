package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Push(-1)
	obj.Push(2)
	obj.Push(5)
	obj.Push(4)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3)
	fmt.Println(param_4)
}

type MinStack struct {
	stack []int
	min   []int
	top   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min(val, (this.min)[this.top]))
	}
	this.stack = append(this.stack, val)
	this.top++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.min = this.min[:this.top]
	this.top--
	return
}

func (this *MinStack) Top() int {
	return (this.stack)[this.top]
}

func (this *MinStack) GetMin() int {
	return (this.min)[this.top]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
