package main
func main(){
	
}
type project struct {
    Cap int
    Pro int
}

type capheap []project

func (ch capheap) Len() int {
    return len(ch)
}

func (ch capheap) Less(i, j int) bool {
    return ch[i].Cap < ch[j].Cap
}

func (ch capheap) Swap(i, j int) {
    ch[i], ch[j] = ch[j], ch[i]
}

func (ch *capheap) Push(x interface{}) {
    *ch = append(*ch, x.(project))
}

func (ch *capheap) Pop() interface{} {
    l := len(*ch)
    
    item := (*ch)[l-1]
    
    (*ch) = (*ch)[:l-1]
    
    return item
}



type pfheap []project

func (ch pfheap) Len() int {
    return len(ch)
}

func (ch pfheap) Less(i, j int) bool {
    return ch[j].Pro < ch[i].Pro
}

func (ch pfheap) Swap(i, j int) {
    ch[i], ch[j] = ch[j], ch[i]
}

func (ch *pfheap) Push(x interface{}) {
    *ch = append(*ch, x.(project))
}

func (ch *pfheap) Pop() interface{} {
    l := len(*ch)
    
    item := (*ch)[l-1]
    
    (*ch) = (*ch)[:l-1]
    
    return item
}


func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    
    capHeap := &capheap{}
    profHeap := &pfheap{}
    
    heap.Init(capHeap)
    heap.Init(profHeap)
    
    for k, v := range capital {
        heap.Push(capHeap, project{
            Cap : v,
            Pro : profits[k],
        })
    }
    
    for k > 0 {
        
        k--
        
        for len(*capHeap) > 0 && (*capHeap)[0].Cap <= w {
            heap.Push(profHeap, heap.Pop(capHeap))
        }
        
        if len(*profHeap) == 0 {
            break
        }
        
        item := heap.Pop(profHeap).(project)
        
        w = w + item.Pro
        
    }
    
    return w
    
}