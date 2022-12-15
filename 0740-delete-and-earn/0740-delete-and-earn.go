import (
    "sort"
)
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
func deleteAndEarn(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    first := 0
    second := 0
    tmp := 0
    lastKey := 0
    m1 := make(map[int]int)
    sorted_keys := make([]int, 0)
    for _, v := range nums {
        m1[v]++
    }
    for k, _ := range m1 {
        sorted_keys = append(sorted_keys, k)
    }
    sort.Ints(sorted_keys)
    for _, k := range sorted_keys {
        if lastKey + 1 == k {
            tmp = max(first + k * m1[k], second)
            first = second
            second = tmp
        } else {
            first = second
            second = second + k * m1[k]
        }
        lastKey = k
    }	
    return second
}