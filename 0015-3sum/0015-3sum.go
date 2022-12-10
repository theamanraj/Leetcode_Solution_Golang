func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var threeSum [][]int
    threeSumMap := make(map[string]bool)
    for i := 0; i < len(nums) - 2; i++ {
        left, right := i + 1, len(nums) - 1
        for left < right {
            if nums[left] + nums[right] == nums[i] * -1 {
                key := strconv.Itoa(nums[i]) + ":" + strconv.Itoa(nums[left]) + ":" + strconv.Itoa(nums[right])
                if _, ok := threeSumMap[key]; !ok {
                    threeSum = append(threeSum, []int{nums[i], nums[left], nums[right]})
                    threeSumMap[key] = true
                }
                left++
                right--
            } else if nums[left] + nums[right] > nums[i] * -1 {
                right--
            } else {
                left++
            }
        }
    }
    return threeSum
}