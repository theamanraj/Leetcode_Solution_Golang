import "math"
func rangeBitwiseAnd(left int, right int) int {
    msbx := getMostSignificantBit(left)
    msby := getMostSignificantBit(right)
 
    var result int
    
    for msbx == msby && left >0 {
        result = result + int(math.Pow(2,float64(msbx)))
        
        left = left - int(math.Pow(2,float64(msbx)))
        right = right - int(math.Pow(2,float64(msbx)))
        
         msbx = getMostSignificantBit(left)
         msby= getMostSignificantBit(right)
    }
    
    return result
}

func getMostSignificantBit(num int) int {
    
    var i int
    arr := make([]int,32)
    
    for num > 0 {
        rem := num%2
        arr[i] = rem
        i++
        num = num/2
    }
    
    return i-1
}