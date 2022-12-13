func convert(s string, numRows int) string {
    if 1 == numRows{
        return s
    }
    
    arr := make([][]byte, numRows)
    
    for i, flipper, j := 0, -1, 0; i < len(s); i++ {    
        arr[j] = append(arr[j], s[i])
        
        if (j%(numRows-1) == 0){
            flipper *= -1
        }  
        j = j + flipper
    }
    
    
    for i := 1; i < numRows; i++{
        arr[0] = append(arr[0], arr[i]...)
    }
    
    return string(arr[0])
}