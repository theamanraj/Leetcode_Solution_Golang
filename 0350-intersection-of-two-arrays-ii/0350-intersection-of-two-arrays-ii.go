func intersect(nums1 []int, nums2 []int) []int {
    map1 := make(map[int]int, len(nums1))
	map2 := make(map[int]int, len(nums2))
    
    for i :=0;i<len(nums1);i++{
        val,ok := map1[nums1[i]]
            fmt.Println("val for",nums1[i],"val",val,"ok",ok)
        if ok {
            map1[nums1[i]] = val+1
        } else {
            map1[nums1[i]] = 1
        }
    }
    fmt.Println("map1",map1)
        for i :=0;i<len(nums2);i++{
        val,ok := map2[nums2[i]]
        if ok {
            map2[nums2[i]] = val+1
        } else {
            map2[nums2[i]] = 1
        }
    }
    fmt.Println("map2",map2)
     res := []int{}
    for key,value := range map1{
          val,ok := map2[key]
        fmt.Println(key,val,ok,value)
        if ok {
             freq := math.Min(float64(val), float64(value))
           for freq > 0 {
                res = append(res, key)
                    freq--
                }
        }
    }
    return res
}