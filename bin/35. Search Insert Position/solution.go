package main
func main(){
	print(searchInsert([]int{1,3,5,6},0))
}

func searchInsert(nums []int, target int) int {
    l,h,mid:=0,len(nums),0

     
    for l < h{
    	mid = (l+h)/2
    	 
    	if (target >nums[mid]){
    		l = mid + 1
    	}else{
    		h = mid 
    	}
    }
return l
}