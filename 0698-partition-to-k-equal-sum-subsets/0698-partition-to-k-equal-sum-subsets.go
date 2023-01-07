func dfs(nums []int, k, target, current, used int) bool {
    if k == 0 {
        return used == (1 << len(nums)) - 1
    }
    
    for i := 0; i < len(nums); i++ {
        if used & (1 << i) != 0 {
            continue
        }
        
        accu := current + nums[i]
        if (accu > target) { break } // Triming 
        
        newUsed := used | (1 << i)
        
        if accu == target {
            if dfs(nums, k-1, target, 0, newUsed) {
                return true
            }
        } else {
            if dfs(nums, k, target, accu, newUsed) {
                return true
            }
        }
    }
    
    return false    
}

func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0 
    for _, val := range nums {
        sum = sum + val
    }
    
    if sum % k != 0 {
        return false
    }
    
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    
    return dfs(nums, k, sum/k, 0, 0)
}