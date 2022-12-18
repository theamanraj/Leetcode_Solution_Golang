func countBits(num int) []int {
    ret := []int{0}
	for i:=1;i<=num;i++{
		info :=0
        if i&1 == 1 {
            info = ret[len(ret)-1]+1 
        }else{
            info = ret[len(ret)/2]
        }
        ret= append(ret,info)
	}
	return ret
}