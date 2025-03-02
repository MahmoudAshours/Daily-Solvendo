func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    mp := make(map[int]int)

    for _,v := range nums1{
        mp[v[0]] = v[1]
    }

    for _,v := range nums2{
        if _,ok:=mp[v[0]];ok{
            mp[v[0]] = mp[v[0]] + v[1]
        }else{
            mp[v[0]] = v[1]
        }
    }
    res := make([][]int,0)
    for i:=1; i <= 1000; i++{
        if _,ok:=mp[i];ok{
            res = append(res,[]int{i,mp[i]})
        }
    }
    return res
}