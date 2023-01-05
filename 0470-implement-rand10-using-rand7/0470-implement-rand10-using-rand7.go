func rand10() int {
    ans := 49
    for ans > 40 {
        ans = (rand7()-1)*7 + rand7()
    }
    return ans % 10 + 1
}