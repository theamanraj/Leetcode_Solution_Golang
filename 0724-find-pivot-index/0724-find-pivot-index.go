func pivotIndex(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    
    leftSum := 0
    for i := range nums {
        if leftSum == totalSum - leftSum - nums[i] {
            return i
        }
        leftSum += nums[i]
    }
    
    return -1
}