func subarraysDivByK(nums []int, k int) int {
    remCount := make([]int, k)
    remCount[0] = 1
    numSubarry, prefixSum := 0, 0
    for i := range nums {
        prefixSum += nums[i]
        rem := prefixSum % k
        if rem < 0 {
            rem += k
        }
        numSubarry += remCount[rem]
        remCount[rem]++
    }
    return numSubarry
}