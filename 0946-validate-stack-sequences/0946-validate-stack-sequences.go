func validateStackSequences(pushed []int, popped []int) bool {
    stk := stack{}
    pu, po := 0, 0
    
    for pu <= len(pushed) && po < len(popped) {
        if stk.Len() > 0 && stk.Top() == popped[po] {
            stk.Pop()
            po++
        } else if pu < len(pushed) {
            stk.Push(pushed[pu])
            pu++
        } else {
            return false
        }
    }
    
    
    return true
}


type stack []int

func (s *stack) Push(v int) {
    *s = append(*s, v)
}

func (s *stack) Pop() (int, error) {
    if s.Len() == 0 {
        return 0, errors.New("Empty Stack")
    }
    
    res:=(*s)[len(*s)-1]
    *s=(*s)[:len(*s)-1]
    return res, nil
}

func (s *stack) Top() int {
    return (*s)[len(*s)-1]
}

func (s *stack) Len() int {
    return len(*s)
}