func max( a int, b int ) int{
    if a>b{
        return a
    }else{
        return b
    }
}
func solve(nums[] int) int{
    n := len(nums)
    prev, curr := nums[0], nums[0] 
    if n == 1 {
      return curr  
    }
    curr = max(nums[0], nums[1])
    if n == 2{
        return curr
    }
    
    var temp int
    for i := 2; i < n; i++{
        temp = prev
        prev = curr
        curr = max(temp + nums[i], prev)
    }
    return curr
}
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
      return nums[0]  
    }
   
    return max(solve(nums[1:]), solve(nums[:n-1]))
}