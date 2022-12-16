func canCompleteCircuit(gas []int, cost []int) int {
    return  f(h1(gas, cost, len(gas)))
}

func h1(g []int, c[]int, N int) []int {
    rt := make([]int, N)
    for i:=0; i<N; i++ {
        rt[i] = g[i] - c[i]
    }
    return rt
}

func h2(d []int) int {
    rt := 0
    for _, v := range d {
        rt += v
    }
    return rt
}

func f(d []int) int {
    if h2(d) < 0 {
        return -1
    }    

    s := 0
    idx := 0
    
    for i, v := range d {
        s += v
        if s < 0 {
            s = 0
            idx = i+1
        }
    }
    
    return idx 
}
