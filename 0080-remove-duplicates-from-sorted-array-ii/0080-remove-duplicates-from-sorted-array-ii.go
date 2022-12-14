func removeDuplicates(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    }
    
    cnt := 1
    cur := 0
    for i := 1; i < length; i++ {
        if nums[cur] != nums[i] {
            cnt = 1
            cur++
            nums[cur] = nums[i]
        } else if cnt < 2 {
            cnt++
            cur++
            nums[cur] = nums[i]
        }        
    }
    return cur + 1
}