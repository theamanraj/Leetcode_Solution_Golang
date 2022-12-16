func maxProduct(nums []int) int {
    minRes, maxRes, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        tempMin := min(n, minRes * n, maxRes * n)
        maxRes = max(n, minRes * n, maxRes * n)
        minRes = tempMin
        
        res = max(maxRes, res)
    }
    return res
}

func min(nums ...int) int {
    min := nums[0]
    for _, v := range nums {
        if v < min {
            min = v
        }
    }
    return min
}

func max(nums ...int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}