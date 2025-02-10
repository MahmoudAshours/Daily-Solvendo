func minChanges(s string) int {
    changes:=0
   if len(s) == 2{
    ones,zeros := onesZerosCount(s)
    return min(ones,zeros)
   }
   for i:=0; i < len(s)-1; i+=2{
    if s[i] != s[i+1]{
        changes++
    }
   }
   return changes
}

func onesZerosCount(s string)(int,int){
    ones:=0
    zeros:=0
    for _,v := range s{
        if v == '1'{
            ones++
        }else{
            zeros++
        }
    }
    return ones,zeros
}
//11001100