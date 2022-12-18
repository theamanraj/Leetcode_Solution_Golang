func isPowerOfFour(num int) bool {
    if num == 0 {
        return false
    }
    return int(math.Pow(4, float64(int(math.Log2(float64(num)) / 2)))) == num 
}