func maximumPopulation(logs [][]int) int {
    tmp, ans := [101]int{}, 0
    for _, l := range logs {
        tmp[l[0] - 1950]++
        tmp[l[1] - 1950]--
    }
    for i := 1; i < 101; i++ {
        tmp[i] += tmp[i - 1]
        if tmp[i] > tmp[ans] { ans = i }
    }
    return ans + 1950
}