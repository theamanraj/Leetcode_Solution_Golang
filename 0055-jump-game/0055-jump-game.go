func canJump(nums []int) bool {
    l := len(nums)
    if l == 1 { return true }
    if nums[0] == 0 { return false }
    
    goal := l-1

    for i:=l-1; i>=0;i--{
        if i+nums[i] >= goal {
            goal = i
        }
    }
    return goal == 0
}