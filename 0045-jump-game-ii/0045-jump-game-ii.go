func jump(nums []int) int {

    if len(nums) <= 1 {
        return 0
    }
    step := 0 + nums[0]
    cur := 1
    ten := nums[0]
    for i := 0; i <= ten; i++ {
        if ten >= len(nums) - 1 {
            return cur
        }
        k := i + nums[i] 
        if ten < k {
            ten = k
            if ten >= len(nums)-1 {
                return cur + 1
            }
        }
        if i >= step {
            cur += 1
            step = ten
        }
    }
    return 0
}