func answerQueries(nums []int, queries []int) []int {
    sort.Ints(nums)
    for i := range nums {
        if i == 0 {
            continue
        }
        nums[i] += nums[i-1]
    }
    for i := range queries {
        queries[i] = sort.Search(len(nums), func(j int) bool {
            return nums[j] > queries[i]
        })
    }
    return queries
}