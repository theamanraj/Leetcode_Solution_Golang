func removeDuplicates(nums []int) int {
    l := 0
    visited := make(map[int] int)
    
    for i := 0;i< len(nums);i++ {
        visitedCount := visited[nums[i]]
        if visitedCount < 2 {
            visited[nums[i]] = visitedCount+1
            nums[l] = nums[i]
            l++
        }
    }
    return l
}