func findNumberOfLIS(a []int) int {
    n := len(a)
    res := 0
    maxLen := 0

    length := make([]int, n)
    count := make([]int, n)

    for j := 0; j < n; j++ {
        length[j] = 1
        count[j] = 1
        for i := 0; i < j; i++ {
            if a[i] < a[j] {
                if length[j] == length[i]+1 {
                    count[j] += count[i]
                } else if length[j] < length[i]+1 {
                    length[j] = length[i] + 1
                    count[j] = count[i]
                }
            }
        }

        if maxLen == length[j] {
            res += count[j]
        } else if maxLen < length[j] {
            res = count[j]
            maxLen = length[j]
        }
    }

    return res
}