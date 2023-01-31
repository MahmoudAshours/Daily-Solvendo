package main
func main(){
	print(equalizeArray([]int32{1,2,2,3}))
}


func equalizeArray(arr []int32) int32 {
	var max int32
	max = 0
	freqMap := make(map[int32]int32, 0)
	
	for _, v := range arr {
		freqMap[v]++
	}
	for _, v := range freqMap {
		if v > max{
			max = v
		}
	}
	return int32(len(arr)) - max
}
