func countOdds(low int, high int) int {
    odd :=0
    if low % 2 == 1 || high %2 ==1{
        odd++
    }
    
    odd += (high - low )/2
    return odd
}