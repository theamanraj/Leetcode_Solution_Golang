func fourSumCount(a []int, b []int, c []int, d []int) int {
    count := 0
    
    ln := len(a)
    m1 := make(map[int]int)
    for i := 0; i < ln; i++ {
        for j := 0; j < ln; j++ {
            m1[a[i]+b[j]]++
        }
    }
    
    for i := 0; i < ln; i++ {
        for j := 0; j < ln; j++ {
            count += m1[-1*(c[i]+d[j])] 
        }
    }

    return count
}