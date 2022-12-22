func checkPerfectNumber(num int) bool {
    var divisiorSum int = 1
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            divisiorSum += i + (num/i)
        } 
    }
    return num != 1 && divisiorSum == num
}