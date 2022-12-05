func maximumWealth(accounts [][]int) int {
    ans := 0
    for i := 0; i < len(accounts); i++ {
        tmp := 0
        for _, v := range accounts[i] { tmp += v }
        if tmp > ans { ans = tmp }
    }
    return ans
}