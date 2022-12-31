var cap = 1000000007
func maxProfit(inventory []int, orders int) (sum int) {
    highest := 0
    for _, i := range inventory { if i > highest { highest = i  }  }
    
    // get the target price: after no mre than orders reduction,     
    targetPrice := sort.Search(highest , func(target int) bool {  // use binary search to get the bound 
        trans := 0 // trans is the number of transactions. 6->3: need 3 tiems trade.
        for _, val := range inventory {
            if val - target < 0  { continue } // ignore as note 1.
            trans += val - target   
        }
        return trans <= orders // cannot have more trade than orders.
    })   // find the target price. 
    
    return totalValue( inventory, orders, targetPrice)    
}

func totalValue(inventory []int, orders int, target int )(sum int){
    for _, value:= range inventory {        
        if value - target <= 0 {continue }                
        batch := (value + target + 1) * (value - target) / 2    // sum of (value, value-1, value-2.... bound) ,as note 2
        if sum+=batch; sum>cap { sum%=cap}     // if too large   
        orders -= value - target        // for this number, how many orders are done.
    }
    if orders>0 { sum+=target*orders} // All left orders with the bound value, as note 3
    sum = sum % cap
    return 
}