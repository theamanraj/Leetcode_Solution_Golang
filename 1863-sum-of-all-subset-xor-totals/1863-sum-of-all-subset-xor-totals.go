func subsetXORSum(nums []int) int {
    res := 0
    list := make([]int, 0)
    backTrack(nums, 0, list, &res)
    return res
}

func backTrack(nums []int, pos int, list []int, res *int) {
    *res += xorSum(list)
    for i := pos; i < len(nums); i++ {
        list = append(list, nums[i])
        backTrack(nums, i+1, list, res)
        list = list[:len(list) - 1]
    }
}

func xorSum(list []int)int{
    if len(list) == 0 {
        return 0
    }
    res := list[0]
    for i := 1; i < len(list); i++ {
        res ^= list[i]
    }
    return res
}

