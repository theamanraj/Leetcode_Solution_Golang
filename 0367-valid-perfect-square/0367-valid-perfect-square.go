func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    
    for i:=2;i*i<=num;i++ {
        if i*i == num {
            return true
        }
    }
    return false
}