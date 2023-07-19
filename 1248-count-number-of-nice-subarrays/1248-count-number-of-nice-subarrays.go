func numberOfSubarrays(nums []int, k int) (ans int) {
    current, count, last := 0, 0, -1 
    for _, v := range nums {
        if v & 1 == 0 {
           ans += current
        } else {
            count++ 
            if count == k {
                count--
                
                next := last + 1
                for nums[next] & 1 == 0 {
                    next++
                }
                current, last = next - last, next
                
                ans += current
            }
        }
    }
    
    return ans
}