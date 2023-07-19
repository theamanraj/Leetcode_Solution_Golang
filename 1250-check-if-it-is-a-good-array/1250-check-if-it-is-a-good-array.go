func isGoodArray(nums []int) bool {
    
    gcd := nums[0]
    
    for  i :=1; i< len(nums);i++ {
        gcd = GCD(gcd,nums[i])
        if gcd == 1 {
            break
        }
    }
    
    return gcd ==1
}



func GCD(j,k int) int {
    a:=0
    b:=0
    
    if(j>k){
        a=j
        b=k
    }else{
        a=k
        b=j
    }
	
	
	for (true){
		if( a%b == 0.0) {
			return b
		}
		
		temp := b
		b= a%b
		a=temp
		
	}
	
	return -100

}