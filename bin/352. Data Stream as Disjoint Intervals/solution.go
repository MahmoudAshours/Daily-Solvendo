package main

import "sort"

func printIntervals(obj *SummaryRanges){
	  for _, v := range obj.GetIntervals() {
	  	for _, d := range v {
	  		print(d)
	  	}
	  	print("\n")
	  } 
}

func main(){
	obj := Constructor()
	obj.AddNum(1) 
	obj.AddNum(3) 
	obj.AddNum(7) 
	obj.AddNum(2) 
	
	printIntervals(&obj)
	obj.AddNum(6)
}

type SummaryRanges struct {
    stream []int
}


func Constructor() SummaryRanges {
    return SummaryRanges{
    	stream : []int{},
    }
}


func (this *SummaryRanges) AddNum(value int)  {
    this.stream = append(this.stream,value)
    sort.Ints(this.stream)
}


func (this *SummaryRanges) GetIntervals() [][]int {
	disjoint:=make([][]int, 0) 
	start:= 0
	end := 0
	for i := 0; i < len(this.stream)-1; i++ {
		if this.stream[i]+1 == this.stream[i+1] || this.stream[i] == this.stream[i+1]{
			end = i+1	
		} else {
			disjoint = append(disjoint,[]int{this.stream[start],this.stream[i]})
			end = len(this.stream)-1
			start = i + 1 
		}
	} 
 if end == len(this.stream)-1{
 
	disjoint = append(disjoint,[]int{this.stream[start],this.stream[end]})
 }
	return disjoint
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */