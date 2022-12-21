func nthSuperUglyNumber(n int, primes []int) int {
    if n==1{
        return 1
    }
    var results []int
    start:=make([]int,len(primes))
    results=append(results,1) 
    for k:=1;k<n;k++{
        var temp,pos int
        for i:=0;i<len(primes);i++{
            zz:=primes[i]*results[start[i]]
                if zz<=results[len(results)-1]{
                    start[i]++
                    zz=primes[i]*results[start[i]]
                }
                if zz>results[len(results)-1]&&temp==0{
                    temp=zz
                    pos=i
                }else if zz>results[len(results)-1]&&zz<temp{
                    temp=zz
                    pos=i
                }
        }
        start[pos]++
        results=append(results,temp)
    }
    return results[len(results)-1]
}