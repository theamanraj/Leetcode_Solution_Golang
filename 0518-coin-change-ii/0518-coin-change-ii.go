func change(amount int, coins []int) int {
    n := len(coins)
    prev := make([]int, amount+1)
    curr := make([]int, amount+1) 
    
    for t:=0; t<=amount; t++ {
        if(t % coins[0] == 0) {
            prev[t] = 1
        } else {
            prev[t] = 0
        }
    }
    
    for ind:= 1; ind<n; ind++ {
        for t:=0; t<=amount; t++ {
            notTake := prev[t]
            take := 0
            if(coins[ind] <= t) {
                take = curr[t- coins[ind]]
            }
            curr[t] = take + notTake
        }
        prev = curr
    }  
    return prev[amount]  
}