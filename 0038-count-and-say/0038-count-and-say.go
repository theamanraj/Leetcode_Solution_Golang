func countAndSay(n int) string {
    if n==1 { return "1"}     
    return process(countAndSay(n-1))
}

func process( s string) string{
    var ans string        
    for i:=0;i<len(s); i++{
        count := 1    
        for j:=i+1;j<len(s);j++ {
            if s[i]!=s[j]{ break}
            count++
            i++
        }
        ans+=strconv.Itoa(count)+string(s[i]) 
    } 
    return ans
}