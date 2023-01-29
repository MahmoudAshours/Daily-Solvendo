package main

import "math"

func main(){
	obj := Constructor(2)
	obj.Put(1,1) 
	obj.Put(2,2)
	print(obj.Get(1))
	obj.Put(3,3)
	print(obj.Get(2))
	print(obj.Get(3))
	obj.Put(4,4)
	print(obj.Get(1))
	print(obj.Get(3))
	print(obj.Get(4))


}type node struct {
	key       int
	val       int
	frequency int
}

type LFUCache struct {
	nodes        map[int]*list.Element
	lists        map[int]*list.List
	cap          int
	minFrequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:          capacity,
		nodes:        make(map[int]*list.Element),
		lists:        make(map[int]*list.List),
		minFrequency: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.nodes[key]; !ok {
		return -1
	}
	currNode := this.nodes[key].Value.(*node)
	this.lists[currNode.frequency].Remove(this.nodes[key])
	currNode.frequency++
	newList := this.lists[currNode.frequency]
	if _, ok := this.lists[currNode.frequency]; !ok {
		newList = list.New()
	}
	this.lists[currNode.frequency] = newList
	newNode := newList.PushFront(currNode)
	this.nodes[key] = newNode
	if currNode.frequency-1 == this.minFrequency && this.lists[currNode.frequency-1].Len() == 0 {
		this.minFrequency++
	}
	return currNode.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	} 
	if val, ok := this.nodes[key]; ok {
		currNode := val.Value.(*node)
		currNode.val = value
		this.Get(key)
		return
	}
 
	if this.cap == len(this.nodes) { 
		minList := this.lists[this.minFrequency]
		backNode := minList.Back()
		delete(this.nodes, backNode.Value.(*node).key)
		minList.Remove(backNode)
	}
    this.minFrequency = 1
	currNode := &node{
		key:       key,
		val:       value,
		frequency: 1,
	}
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newNode := this.lists[1].PushFront(currNode)
	this.nodes[key] = newNode
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */