func largestSumAfterKNegations(A []int, K int) int {
    sort.Slice(A, func(i,j int) bool {
        return A[i] < A[j]
    })
    
    sum := 0
    
    for i, n := range A {
        if n < 0 && K > 0 {
            A[i] = -A[i]
            
            K--
        }
        
        sum += A[i]
    }    
    
    if K % 2 != 0 {
        min := A[0]
        
        for _, n := range A {
            if min > n {
                min = n
            }
        }
        
        
        sum -= min * 2
    }
    
    return sum
}