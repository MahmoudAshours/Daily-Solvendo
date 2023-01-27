package main

func main(){
	print(canConstruct("a","aba"))
}

func canConstruct(ransomNote string, magazine string) bool {
  magazineMap := make(map[rune]int, 0)
 
 for _, v := range magazine {
 	 	magazineMap[v]++
 }
for _, v := range ransomNote {
	magazineMap[v]--
}
for _, v := range magazineMap {
	if v < 0 {
		return false
	}
}
 return true   
}