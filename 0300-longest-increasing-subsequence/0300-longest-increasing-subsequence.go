func lengthOfLIS(nums []int) int {
    var r = make([]int, 0, len(nums))
    
    for _,v := range nums{
        var tmep_i = bisec(r, v)
        if tmep_i == len(r){
            r = append(r, v)
        }else if tmep_i < len(r){
            r[tmep_i] = v
        }
    }
    
    return len(r)
    
}

func bisec(nums []int, t int) int {
    if len(nums) == 0{
        return 0
    }
    
    l, r := 0, len(nums)
    for l < r{
        m := (l + r)/2
        if t <= nums[m]{
            r = m
        }else{
            l = m + 1
        }
    }
    return l
}