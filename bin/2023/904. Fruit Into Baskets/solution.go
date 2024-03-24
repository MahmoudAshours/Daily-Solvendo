package main

import "math"

func main(){
	print(totalFruit([]int{1,2,2,3,1}))
}

func totalFruit(fruits []int) int {
    hashM := make(map[int]int, 0)
    res,l,total :=0,0,0
    for _, v := range fruits {
    	hashM[v]++
   		total++
   		for len(hashM) > 2{
   			f := fruits[l]
   			hashM[f]--
   			total--
   			l++
   			if hashM[f] == 0{
   				delete(hashM,f)
   			}
   		}
   		res = int(math.Max(float64(res),float64(total)))
    } 
    return res
}