func arraySign(nums []int) int {
    var res int = 1
    for _, num := range nums {
        if num == 0 {
            return 0
        }
        if num > 0 {
            res *= 1
        } else {
            res *= -1
        }
    }
    return res
}