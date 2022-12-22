func arrayPairSum(nums []int) (sum int) {
    sort.IntSlice(nums).Sort()
    for i := 0; i < len(nums); i += 2 {
        sum += nums[i]
    }
    return sum
}