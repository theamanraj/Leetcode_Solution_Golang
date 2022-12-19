func lastRemaining(n int) int {
    s := 1 // power of 2
    d := 0 // count the number of steps
    res := 0 // store the result
    // head is the first (smallest) number before elimination
    // tail is the last (largest) number before elimination    
    head := 1; tail := n 
    for tail > head {
        if d % 2 == 0 { // eliminate from top to bottom
            if head % 2 == 0 {
                res += s
            }
            if (tail - head) % 2 == 0 {
                tail --
            }
            head ++
        } else {  // eliminate from bottom to top
            if tail % 2 == 0 {
                res += s
            } 
            if (tail - head) % 2 == 0 {
                head ++
            }
            tail --
        }
        tail = tail / 2
        head = head / 2
        s = s * 2
        d ++
    }
    return res + tail * s 
}