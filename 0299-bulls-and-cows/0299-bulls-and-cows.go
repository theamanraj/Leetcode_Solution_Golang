func getHint(secret string, guess string) string {
    l := len(secret)
    bulls := 0
    cnt := [10]int{}
    for i := 0; i < l; i++{
        if secret[i] == guess[i]{
            bulls++
        }else{
            cnt[secret[i]-48]++
            cnt[guess[i]-48]--
        }
    }
    sum := 0
    for _,v := range(cnt){
        if v <= 0{
            sum -= v
        }else{
            sum += v
        }
    }
    cows := l-bulls-sum/2
    res := strconv.Itoa(bulls)
    res += "A"
    res += strconv.Itoa(cows)
    res += "B"
    return res
}