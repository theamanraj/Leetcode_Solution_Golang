func sumBase(n int, k int) int {
    var res int = 0
    for  ; n>0; n/=k { // Loop until the number is greater than 0
        res += (n%k) // Each digit will be the remainder when divided by given base value
    }
    return res
}