func findMax(value1 int, value2 int) int {
    if (value1 >= value2){
        return value1
    }else{
        return value2
    }
}

func maxProfit(prices []int) int {

	//  make three dp arrays 
	// buy - keep track the maximum profit history so far if you were buy today
	// sell - keep track the profit if you were to sell today
	// total - keep track the maximum profit so far between if were to sold it yesterday and the history
    
    var num int = len(prices) 
    buy := make([]int, num)
    sell := make([]int, num)
    total := make([]int, num)
    
    if (len(prices) < 2){
        return 0
    }
    
    buy[0] = -prices[0]
    for i := 1; i < len(prices); i++ {
        buy[i] = findMax(buy[i-1], total[i-1]-prices[i])
        sell[i] = buy[i-1] + prices[i]
        total[i] = findMax(sell[i-1], total[i-1])
    }
    
    return findMax(findMax(total[num-1], sell[num-1]), buy[num-1])
}