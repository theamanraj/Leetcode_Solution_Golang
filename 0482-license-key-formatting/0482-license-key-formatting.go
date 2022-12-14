func licenseKeyFormatting(s string, k int) string {
    res := []byte{}
    counter := 0
    for i := len(s) - 1; i >= 0; i--{
        if s[i] != '-'{
            if counter % k == 0 && counter > 0{
                res = append(res,'-')
            }            
            if s[i] >= 'a' && s[i] <= 'z'{
                res = append(res, byte(s[i] + 'A' - 'a'))
            }else{
                res = append(res, s[i])
            }
            counter++
        }
    }
    for i,j := 0,len(res)-1; i < j; i,j = i+1, j-1{
        res[i], res[j] = res[j], res[i]
    }
    return string(res)
}