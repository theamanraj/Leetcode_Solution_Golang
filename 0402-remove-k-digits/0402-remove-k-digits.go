func removeKdigits(num string, k int) string {
    // greedy monotonically inc stack
    stack := make([]byte, 0)
    
    for i := 0; i<len(num); i++ {
        for len(stack) > 0 && num[i] < stack[len(stack)-1] && k > 0  {
            stack = stack[:len(stack)-1]
            k--
        }
        
        // stack integrity should be held, no leading zeros
        if len(stack) > 0 || num[i] > '0' {
            stack = append(stack, num[i])
        }
    }
    
    for k > 0 && len(stack) > 0 {
        stack = stack[:len(stack)-1]
        k--
    }
    
    if len(stack) == 0 {
        return "0"
    }
    
    return string(stack)
}