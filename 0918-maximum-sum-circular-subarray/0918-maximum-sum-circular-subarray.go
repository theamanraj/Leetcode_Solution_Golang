func maxSubarraySumCircular(nums []int) int {
    var sum, minVal, maxVal int // maxValue or minValue of current sub string.
    globalMax, globalMin := math.MinInt32, math.MaxInt32 // maxValue or minValue of the whole string.
    
    for _, v := range nums{
        sum+=v
        maxVal, minVal = max(maxVal+v, v), min(minVal+v, v)
        globalMax, globalMin = max(globalMax,maxVal ), min(globalMin, minVal )                  
    }   
    if globalMax<=0 { return globalMax}
    return max(globalMax, sum-globalMin) // Circular max= sum - minValue
}

func max(i, j int)int{
    if i>j{ return i}
    return j
}

func min(i, j int)int{
    if i<j{ return i}
    return j
}