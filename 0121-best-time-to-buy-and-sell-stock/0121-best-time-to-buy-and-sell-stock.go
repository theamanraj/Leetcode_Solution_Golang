func maxProfit(prices []int) int {
    l,r,maxp := 0,1,0
    for r < len(prices) {
        if prices[l] < prices[r]{
            profit := prices[r]-prices[l]
            maxp = checkMax(maxp,profit)
        } else {
            l = r
        }
        r++
    }
    return maxp
}

func checkMax(a int, b int) int {
    if a > b {
        return a
    }
    return b
}