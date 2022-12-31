func checkSubarraySum(nums []int, k int) bool {
    sum := 0
    prev_idx := map[int]int{0:-1}
    
    for idx, num := range nums{
        sum += num
        sum = sum%k
    
        if val, ok := prev_idx[sum]; ok {
            if (idx-val) >= 2 {
                return true
            }
        }else{
            prev_idx[sum] = idx
        }
    }
    
    return false
}