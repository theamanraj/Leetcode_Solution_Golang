func maxAbsValExpr(arr1 []int, arr2 []int) int {
    maxSum := 0
    for i := 0; i < len(arr1); i++ {
        for j := i+1; j < len(arr2); j++ {
            sum := Abs(arr1[i] - arr1[j]) + Abs(arr2[i] - arr2[j]) + Abs(i - j)   
            if sum > maxSum {
                maxSum = sum
            }
        }
    }
    return maxSum
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}