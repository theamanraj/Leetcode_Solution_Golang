func findDuplicates(nums []int) []int {
    res := []int{}
    for i := 0; i < len(nums); i++ {
        index := abs(nums[i]) - 1
        if nums[index] < 0 {
            res = append(res, index + 1)
        } else {
            nums[index] = -1 * nums[index]
        }
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}