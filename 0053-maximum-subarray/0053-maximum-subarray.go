func maxSubArray(nums []int) int {
    maxSum,tempSum := nums[0],0
    for i := 0;i<len(nums);i++{
        tempSum = tempSum+nums[i]
        if tempSum > maxSum{
            maxSum = tempSum
        }
        if tempSum < 0{tempSum = 0}
    }
    return maxSum
}