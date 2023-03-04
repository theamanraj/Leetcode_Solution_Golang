func pivotInteger(n int) int {
    
    if n == 1 {
        return 1
    }

    pivot := n/2;
    left := arithSum(1,pivot)
    right := arithSum(pivot, n-pivot+1)

    for left != right {

        if pivot > n {
            return -1
        }
        
        right -= pivot
        pivot++
        left += pivot
    }

    return pivot
}

func arithSum(a, n int) int {
    ret := (2*a + (n-1)) * n
    ret /= 2;
    return ret
}