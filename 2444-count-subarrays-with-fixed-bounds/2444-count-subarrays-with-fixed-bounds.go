func countSubarrays(nums []int, minK int, maxK int) int64 {
    var res int64
    jbad, jmin, jmax, n := -1, -1, -1, len(nums)
    for i := 0; i < n; i++ {
        if nums[i] < minK || nums[i] > maxK {
            jbad = i
        }
        if nums[i] == minK {
            jmin = i
        }
        if nums[i] == maxK {
            jmax = i
        }
        res += int64(max(0, min(jmin, jmax) - jbad))
    }
    
    return res
}


func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}