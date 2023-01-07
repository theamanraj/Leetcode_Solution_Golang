type direction int

const (
    UP direction = iota
    DOWN
    FLAT
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestMountain(A []int) int {
    N := len(A)
    if N < 3 {
        return 0
    }
    longest := 0
    i := 1
    for ; i < N && A[i] <= A[i - 1]; i++ {;}
    direc, upStep, downStep := UP, 1, 0
    for i++; i < N; i++ {
        if direc == UP {
            if A[i] > A[i - 1] {
                upStep++
            } else if A[i] < A[i - 1] {
                direc, upStep, downStep = DOWN, upStep + 1, 1
            } else {
                direc, upStep, downStep = FLAT, 0, 0
            }
        } else if direc == FLAT {
            for ; i < N && A[i] <= A[i - 1]; i++ {;}
            direc, upStep, downStep = UP, 1, 0
        } else {
            if A[i] > A[i - 1] {
                longest = max(longest, upStep + downStep)
                direc, upStep, downStep = UP, 1, 0
            } else if A[i] < A[i - 1] {
                downStep++
            } else {
                longest = max(longest, upStep + downStep)
                direc, upStep, downStep = FLAT, 0, 0
            }
        }
    }
    if direc == DOWN {
        longest = max(longest, upStep + downStep)
    }
    return longest
}