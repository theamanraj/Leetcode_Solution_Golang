const limit = 1000000007

type mins struct {
	val, count int
}

func sumSubarrayMins(A []int) int {
    left := make([]int, len(A))
    right := make([]int, len(A))
    lstack := []mins{}
    rstack := []mins{}

    for i := 0; i < len(A); i++ {
        count := 1
        for len(lstack) != 0 && lstack[len(lstack)-1].val > A[i] {
            count += lstack[len(lstack)-1].count
            lstack = lstack[:len(lstack)-1]
        }
        lstack = append(lstack, mins{A[i], count})
        left[i] = count
    }

    for i := len(A) - 1; i >= 0; i-- {
        count := 1
        for len(rstack) != 0 && rstack[len(rstack)-1].val >= A[i] {
            count += rstack[len(rstack)-1].count
            rstack = rstack[:len(rstack)-1]
        }
        rstack = append(rstack, mins{A[i], count})
        right[i] = count
    }

    sum := 0

    for i := 0; i < len(A); i++ {
        sum += A[i] * left[i] * right[i]
        sum %= limit
    }

    return sum
}