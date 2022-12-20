const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func maxProfit(prices []int, fee int) int {
    k0 := 0
    k1 := INT_MIN
    
    for _, v := range prices {
        if k0 < k1 + v {
            k0 = k1 + v
        }
        if k1 < k0 - v - fee {
            k1 = k0 - v- fee
        }
    }
    return k0
}