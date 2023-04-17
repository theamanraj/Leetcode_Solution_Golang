func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := -1
    posibilities := make([]bool, 0)
    for i := range candies {
        if candies[i] > max {
            max = candies[i]
        }
    }
    
    for i := range candies {
        posibilities = append(posibilities, candies[i] + extraCandies >= max)
    }
    
    return posibilities
}