var memo = make(map[string]int)
func superEggDrop(k int, n int) int {
    if k == 1 {return n}
    if n == 0 {return 0}
    l,r :=1,n
    key := string(k) + "-" + string(n)
    if v,has := memo[key];has{
        return v   
    }
    res :=n+1
    for l<=r{
        m := (l+r)/2
        broken := superEggDrop(k-1,m-1)
        no_broken := superEggDrop(k,n-m)
        if broken < no_broken{
            l = m + 1
            res = min(res,no_broken+1)
        }else{
            r = m - 1
            res = min(res,broken+1)
        }
    }
    memo[key] = res
    return res
}

func min(a int , b int) int{
    if a< b {return a}
    return b
}