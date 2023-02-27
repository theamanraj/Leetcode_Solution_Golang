func minSetSize(arr []int) int {
    ans := 0
    half := len(arr) / 2
    count := []int{}
    fre := 1
    sort.Ints(arr)
    for i := 1; i < len(arr); i++ {
        if arr[i - 1] == arr[i] {
            fre++
        } else {
            count = append(count, fre)
            fre = 1
        }
    }
    count = append(count, fre)
    sort.Slice(count, func(a, b int) bool { return count[a] > count[b] })
    for i := 0; i < len(count) && half > 0; i++ {
        half = half - count[i]
        ans++
    }
    return ans
}