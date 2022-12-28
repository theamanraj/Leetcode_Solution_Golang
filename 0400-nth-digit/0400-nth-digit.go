func findNthDigit(n int) int {
    table := make([]int, 10) 
    // table[i] = total number of digits from 1 to 10^(i+1) (not including 10^(i+1))
    s := 9
    table[0] = 9
    for i := 1; i < 10; i ++{
        s = s * 10
        table[i] = table[i - 1] + s * (i + 1)
    }    
    nLen := 0 // the lenhth of the number that contains nth digit 
    s = 1
    for ; nLen < 10; nLen ++{
        if n <= table[nLen] {
            break
        }
        s = s * 10 //s will be the starting number, i.e., 10, 100, etc.
    }
    if nLen == 0 {
        return n 
    }
    p := table[nLen - 1]
    num := strconv.Itoa((n - p - 1) / (nLen + 1) + s)
    fmt.Println(num)
    return int(num[(n - p - 1) % (nLen + 1)] - '0')    
}