func maxProfit(prices []int) int {
    min,maxp := prices[0],0
    for _,val := range prices{
        if val < min{
            min = val
        }
        profit := val-min
        if profit > maxp{
            maxp = profit
        }
    }
    return maxp
}