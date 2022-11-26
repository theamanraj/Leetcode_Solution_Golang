func getPermutation(n int, k int) string {
    ns := []string{}
    for i := 1 ; i<= n; i++ {
        ns = append(ns, strconv.Itoa(i))
    }
    return helper("", ns, k)
}

func helper(pre string, ns []string, k int) string {
    n := len(ns)
    if n == 1 {
        return pre + ns[0]
    }
    p := 1
    for i:=1; i<n; i++ {
       p *= i 
    }
    si := (k-1) / p
    s := ns[si]
    newNs := append(ns[:si], ns[si+1:]...)
    return helper(pre+s, newNs, (k-1)%p + 1)
}