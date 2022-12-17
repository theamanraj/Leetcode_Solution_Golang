func findDuplicate(nums []int) int {
    l := len(nums)
    for i := 0; i < l; i++ {
        if nums[i] > 0 {
            if nums[nums[i]] < 0 {
                return nums[i]
            }
            nums[nums[i]] *= -1
        } else {
            if nums[-nums[i]] < 0 {
                return -nums[i]
            }
            nums[-nums[i]] *= -1
        }
    }
    return -1
}