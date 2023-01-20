func selfDividingNumbers(left int, right int) []int {
    list := make([]int, 0)
    for num := left; num <= right; num++ {
        if (isSelfDivisible(num)) {
            list = append(list, num)
        }
    }
    return list
}

func isSelfDivisible(val int) bool {
    num := val
    for ;num > 0; {
        digit := num % 10
        if digit == 0 || val % digit != 0 {
            return false
        }
        num = num / 10
    }
    return true
}