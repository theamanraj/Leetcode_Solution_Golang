func triangleNumber(nums []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    
    n := len(nums)
    ans := 0
    
    for c:=0; c<n-2; c++ {
        b := c+1
        a := n-1
        for b<a {
            if nums[a]+nums[b] > nums[c] {
                ans += a-b
                b++
            } else {
                a--
            }
        }
    }
    
    return ans
}