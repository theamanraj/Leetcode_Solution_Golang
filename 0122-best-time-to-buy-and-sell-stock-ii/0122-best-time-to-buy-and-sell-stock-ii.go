func maxProfit(prices []int) int {
    maxp := 0
    
    for i:=1;i<len(prices);i++{
        if prices[i] > prices[i-1]{
            profit := prices[i]-prices[i-1]
            maxp += profit
        }
    }
    return maxp
}