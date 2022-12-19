func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
func findGap(a, b, n int) int {
    gap := 0 
    for (a <= n) {
        // n+1 => count node x0 too, which the gap includes it when n+1 is min
		// otherwise b is the min
        gap = gap + min(n+1, b) - a
        a = a * 10
        b = b * 10
    }
    
    return gap
}
func findKthNumber(n int, k int) int {
    current := 1
    k = k - 1 // target index is k-1
    
    for k > 0 { // break when reaching target index, where k == 0
        gap := findGap(current, current+1, n)
        
        // k completely cover the gap, move forward. 
        // if gap == k, loop will break next round since k == 0
        if gap <= k { 
            k = k - gap 
            current = current + 1 
        } else {
            // move one step down the tree
            k--
            current = current * 10
        }
    }
    
    return current
}