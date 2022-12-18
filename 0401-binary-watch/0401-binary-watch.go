var ans []string
var arr = []int{8,4,2,1,32,16,8,4,2,1}

func readBinaryWatch(turnedOn int) []string {
    ans = []string{}
    backtrack(turnedOn, 0, 0, 0)
    return ans
}

func backtrack(turnedOn int, H int, M int, k int) {
        
    if turnedOn == 0 {
        Hs, Ms := strconv.Itoa(H), strconv.Itoa(M)
        if M < 10 {
            Ms = "0" + Ms
        }      
        ans = append(ans, Hs + ":" + Ms)
        return
    }
    
    for i := k; i < 10; i++ {
        h, m := H, M
        
        if i < 4 {
            h+=arr[i]
        } else {
            m+=arr[i]
        }
        
        if h < 12 && m < 60 {
            backtrack(turnedOn - 1, h, m, i+1)
        }
           
    }
}